package functionIndicators

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/services/parService"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

type DpOperationRiskRate struct {
}

// 计算存款操作风险率
func (dpo DpOperationRiskRate) Calulate(paramMap map[string]interface{}) (float64, error) {

	var organCode string

	// 加入生效期限
	err := util.GetStartTimeParam(paramMap)
	if nil != err {
		return 0, err
	}
	if _, ok := paramMap["flag"]; !ok {
		paramMap["flag"] = util.FLAG_TRUE
	}

	if organ, ok := paramMap["organ"]; ok {
		organCodeSql := `
		select t.organ
		  from (select distinct dpop_.organ, orgn_.l
		          from rpm_par_dp_op dpop_
		         inner join (select t.ORGAN_CODE, LEVEL as l
		                      from sys_sec_organ t
		                     start with t.ORGAN_CODE = :1
		                    connect by prior t.PARENT_ORGAN = t.ORGAN_CODE) orgn_
		            on (orgn_.ORGAN_CODE = dpop_.ORGAN and dpop_.flag = :2)
		         order by orgn_.l) t
		 where rownum = 1
	`
		rows, err := util.OracleQuery(organCodeSql, organ, util.FLAG_TRUE)
		if nil != err {
			er := fmt.Errorf("递归查询存款操作风险率机构编码出错")
			zlog.Error(er.Error(), er)
			return -1, er
		}
		for rows.Next() {
			rows.Scan(&organCode)
			paramMap["organ"] = organCode
		}
		rows.Close()
	} else {
		er := fmt.Errorf("查询操作风险率未传机构编码信息")
		zlog.Error(er.Error(), er)
		return 0, er
	}

	for "#" != organCode {
		dpOp, err := parService.DpOpService{}.Find(paramMap)
		if nil != err {
			return 0, err
		}
		if 0 < len(dpOp) {
			return dpOp[0].OpRate, nil
		}
		organ, nil := sys.SelectOrganByOrgCode(organCode)
		if nil != err {
			return 0, err
		}
		organCode = organ.ParentOrgan
		paramMap["organ"] = organCode
	}
	return 0, fmt.Errorf("总行产品【%s】存款操作风险率为空", paramMap["product"])
}
