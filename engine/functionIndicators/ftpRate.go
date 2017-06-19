package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var ftpRateService parService.FtpRateService

type FtpRate struct {
}

// 资金成本率查询
// 1、换算期限
// 2、多参数查询资金成本率 查出两条记录
// 3、插值法算出资金成本率返回
func (this *FtpRate) Calulate(paramMap map[string]interface{}) (float64, error) {

	// delete(paramMap, "term_mult")
	// 特殊处理产品、机构
	// if organ, ok := paramMap["organ"]; ok {
	// 	paramMap["organ"] = organ.(sys.Organ).OrganCode
	// }
	// if product, ok := paramMap["product"]; ok {
	// 	paramMap["product"] = product.(dim.Product).ProductCode
	// }
	// 获取期限
	var term int
	value, ok := paramMap["term"]
	if ok {
		term = value.(int)
	}
	delete(paramMap, "term")
	// 判断是否传参数类型
	if nil == paramMap["param_type"] {
		er := fmt.Errorf("计算资金成本率未传参数param_type")
		zlog.Error(er.Error(), er)
		return -1, er
	}

	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		return 0, err
	}

	var whereInSql = `
		select max(term)  term  from RPM_PAR_FTP t  where param_type = '` + paramMap["param_type"].(string) + `' and term <=` + fmt.Sprint(term) + `   union all
		select min(term)  term  from RPM_PAR_FTP t  where param_type = '` + paramMap["param_type"].(string) + `' and term >` + fmt.Sprint(term)
	paramMap["searchLike"] =
		append(paramMap["searchLike"].([]map[string]interface{}), map[string]interface{}{
			"type":  "in",
			"key":   "term",
			"value": whereInSql,
		})
	paramMap["FLAG"] = util.FLAG_TRUE

	// 查询机构
	organCodeSql := `select t.organ
					   from (select distinct ftp_.organ, orgn_.l
					           from rpm_par_ftp ftp_
					          inner join (select t.ORGAN_CODE, LEVEL as l
					                        from sys_sec_organ t
					                       start with t.ORGAN_CODE = :1
					                     connect by prior t.PARENT_ORGAN = t.ORGAN_CODE) orgn_
					             on orgn_.ORGAN_CODE = ftp_.ORGAN
					          where ftp_.flag = :2
					          order by orgn_.l) t
					  where rownum = 1`
	rows, err := util.OracleQuery(organCodeSql, paramMap["organ"], util.FLAG_TRUE)
	if nil != err {
		er := fmt.Errorf("递归查询资金成本率机构信息出错")
		zlog.Error(er.Error(), err)
		return -1, er
	}
	defer rows.Close()
	for rows.Next() {
		var organCode string
		err := rows.Scan(&organCode)
		if nil != err {
			er := fmt.Errorf("递归查询资金成本率机构信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return -1, er
		}
		paramMap["organ"] = organCode
	}

	return this.recursionFindFtp(term, paramMap)
}

func (this FtpRate) recursionFindFtp(term int, paramMap map[string]interface{}) (float64, error) {
	ftps, err := ftpRateService.Find(paramMap)
	if nil != err {
		zlog.Error(err.Error(), err)
		// 错误处理
		return -1, err
	}

	// 先判断查询的值是否为空
	// 循环遍历查出来的数据，线性插值计算
	// if 0 == len(ftps) {
	// 	err := fmt.Errorf("查询资金成本率为空")
	// 	zlog.Error(err.Error(), err)
	// 	return -1, err
	// }

	// 判断Ftp记录是两条还是1条
	// 如果只有一条记录，则判断Term是否相等，否则报错
	// 线性插值公式
	// ftpRate = FtpRate1 + (term - term1)/(term1-term2)*(FtpRate1-FtpRate2)
	switch len(ftps) {
	case 0:
		zlog.Infof("查询资金成本率记录为空，递归机构查询[参数%v]", nil, paramMap)
		organ, err := sys.SelectOrganByOrgCode(paramMap["organ"].(string))
		if nil != err {
			return -1, err
		}
		paramMap["organ"] = organ.ParentOrgan
		return this.recursionFindFtp(term, paramMap)

	case 1:
		if term == ftps[0].Term {
			return ftps[0].FtpRate, nil
		} else {
			var msg string = "查询资金成本率与期限不匹配"
			err := fmt.Errorf(msg)
			zlog.Error(msg, err)
			zlog.Infof("查询资金成本率与期限不匹配[%v]", nil, paramMap)
			return -1, err
		}
	case 2:
		var term1 string = fmt.Sprint(ftps[0].Term)
		var term2 string = fmt.Sprint(ftps[1].Term)
		var ftpRate1 string = fmt.Sprint(ftps[0].FtpRate)
		var ftpRate2 string = fmt.Sprint(ftps[1].FtpRate)
		// 拼接公式字符串
		var formual string = ftpRate1 + "+" + "(" + fmt.Sprint(term) + "-" + term1 + ")/(" + term1 + "-" + term2 + ")*(" + ftpRate1 + "-" + ftpRate2 + ")"
		ftpRate, err := util.Calculate(formual)
		if nil != err {
			var msg string = "资金成本率线性插值错误"
			err := fmt.Errorf(msg)
			zlog.Error(msg, err)
			return -1, err
		}
		return ftpRate, nil
	default:
		var msg string = "查询资金成本率有多条"
		err := fmt.Errorf(msg)
		zlog.Error(msg, err)
		return -1, err
	}
}
