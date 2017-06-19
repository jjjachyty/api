package par

import (
	"database/sql"
	"fmt"
	"time"

	"pccqcpa.com.cn/app/rpm/api/models/dim"
	"pccqcpa.com.cn/app/rpm/api/models/sys"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
)

// 结构体对应数据库表【RPM_PAR_FTP】资金成本率
// by author Yeqc
// by time 2016-12-19 10:23:37
/**
 * @apiDefine Ftp
 * @apiSuccess {string}     UUID            主键默认值sys_guid()
 * @apiSuccess {Product}   	Product    		产品
 * @apiSuccess {Organ}   	Organ      		机构
 * @apiSuccess {string}   	Currency   		币种
 * @apiSuccess {int}   		Term    		期限
 * @apiSuccess {string}   	RateType      	利率类型
 * @apiSuccess {int}   		RepriceFreq     重定价频率
 * @apiSuccess {string}   	RpymType      	还款方式
 * @apiSuccess {int}   		RpymCapitalFreq 还本频率
 * @apiSuccess {int}   		FtpRate      	还息频率
 * @apiSuccess {float64}   	ParamType      	资金成本率
 * @apiSuccess {string}   	StartTime      	参数类型
 * @apiSuccess {string}   	Flag      		生效标志
 * @apiSuccess {time.Time}	CreateTime      创建时间
 * @apiSuccess {string}   	CreateUser      创建人
 * @apiSuccess {time.Time}	UpdateTime      更新时间
 * @apiSuccess {string}   	UpdateUser      更新人
 */
type Ftp struct {
	UUID             string      //SYS_GUID()
	Product          dim.Product //产品
	Organ            sys.Organ   //机构
	Currency         string      //币种
	Term             int         //期限
	RateType         string      //利率类型
	RepriceFreq      int         //重定价频率
	RpymType         string      //还款方式
	RpymCapitalFreq  int         //还本频率
	RpymInterestFreq int         `json:"-"` //还息频率
	FtpRate          float64     //资金成本率
	ParamType        string      //参数类型
	StartTime        time.Time   //生效日期
	CreateTime       time.Time   //创建日期
	CreateUser       string      //创建人
	UpdateTime       time.Time   //更新时间
	UpdateUser       string      //更新人
	Flag             string      //生效标志
}

const ftpTables = `
	     RPM_PAR_FTP T
	LEFT JOIN RPM_DIM_PRODUCT PRODUCt
	ON (T.PRODUCT = PRODUCT.PRODUCT_code)
	LEFT JOIN SYS_SEC_ORGAN ORGAN
	ON (T.ORGAN = ORGAN.ORGAN_CODE)
`

var ftpCols = map[string]string{
	"T.UUID":               "' '",
	"T.CURRENCY":           "' '",
	"T.TERM":               "0",
	"T.RATE_TYPE":          "' '",
	"T.REPRICE_FREQ":       "0",
	"T.RPYM_TYPE":          "' '",
	"T.RPYM_CAPITAL_FREQ":  "0",
	"T.RPYM_INTEREST_FREQ": "0",
	"T.FTP_RATE":           "0",
	"T.PARAM_TYPE":         "' '",
	"T.START_TIME":         "sysdate",
	"T.CREATE_TIME":        "sysdate",
	"T.CREATE_USER":        "' '",
	"T.UPDATE_TIME":        "sysdate",
	"T.UPDATE_USER":        "' '",
	"T.FLAG":               "' '",

	"T.PRODUCT":            "' '",
	"PRODUCT.PRODUCT_NAME": "' '",

	"T.ORGAN":          "' '",
	"ORGAN.ORGAN_NAME": "' '",
}

var ftpColsSort = []string{
	"T.UUID",
	"T.CURRENCY",
	"T.TERM",
	"T.RATE_TYPE",
	"T.REPRICE_FREQ",
	"T.RPYM_TYPE",
	"T.RPYM_CAPITAL_FREQ",
	"T.RPYM_INTEREST_FREQ",
	"T.FTP_RATE",
	"T.PARAM_TYPE",
	"T.START_TIME",
	"T.CREATE_TIME",
	"T.CREATE_USER",
	"T.UPDATE_TIME",
	"T.UPDATE_USER",
	"T.FLAG",

	"T.PRODUCT",
	"PRODUCT.PRODUCT_NAME",

	"T.ORGAN",
	"ORGAN.ORGAN_NAME",
}

func (f Ftp) scan(rows *sql.Rows) (*Ftp, error) {
	var ftp = new(Ftp)
	var values = []interface{}{
		&ftp.UUID,
		&ftp.Currency,
		&ftp.Term,
		&ftp.RateType,
		&ftp.RepriceFreq,
		&ftp.RpymType,
		&ftp.RpymCapitalFreq,
		&ftp.RpymInterestFreq,
		&ftp.FtpRate,
		&ftp.ParamType,
		&ftp.StartTime,
		&ftp.CreateTime,
		&ftp.CreateUser,
		&ftp.UpdateTime,
		&ftp.UpdateUser,
		&ftp.Flag,

		&ftp.Product.ProductCode,
		&ftp.Product.ProductName,

		&ftp.Organ.OrganCode,
		&ftp.Organ.OrganName,
	}
	err := util.OracleScan(rows, values)
	return ftp, err
}

// 分页操作
// by author Yeqc
// by time 2016-12-19 10:23:37
func (this Ftp) List(param ...map[string]interface{}) (*util.PageData, error) {
	pageData, rows, err := modelsUtil.List(ftpTables, ftpCols, ftpColsSort, param...)
	defer rows.Close()
	if nil != err {
		er := fmt.Errorf("分页查询资金成本率信息出错")
		zlog.Error(er.Error(), err)
		return nil, er
	}

	var rst []*Ftp
	for rows.Next() {
		one, err := this.scan(rows)
		if nil != err {
			er := fmt.Errorf("分页查询资金成本率信息row.Scan()出错")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		rst = append(rst, one)
	}
	pageData.Rows = rst
	return pageData, nil
}

func (f Ftp) Find(param ...map[string]interface{}) ([]*Ftp, error) {
	zlog.AppOperateLog("", "Find", zlog.SELECT, nil, nil, "查询资金成本率")
	rows, err := modelsUtil.FindRows(ftpTables, ftpCols, ftpColsSort, param...)
	if nil != err {
		er := fmt.Errorf("资金成本率查询错误")
		zlog.Error(er.Error(), err)
		return nil, er
	}
	defer rows.Close()
	var ftps []*Ftp
	for rows.Next() {
		ftp, err := f.scan(rows)
		if nil != err {
			er := fmt.Errorf("资金成本率查询rows.Scan()错误")
			zlog.Error(er.Error(), err)
			return nil, er
		}
		ftps = append(ftps, ftp)
	}
	return ftps, nil
}

// 更新资金成本率信息
// by author Yeqc
// by time 2016-12-19 10:23:37
func (this Ftp) Update() error {
	paramMap := map[string]interface{}{
		"product":            this.Product.ProductCode,
		"organ":              this.Organ.OrganCode,
		"currency":           this.Currency,
		"term":               this.Term,
		"rate_type":          this.RateType,
		"reprice_freq":       this.RepriceFreq,
		"rpym_type":          this.RpymType,
		"rpym_capital_freq":  this.RpymCapitalFreq,
		"rpym_interest_freq": this.RpymInterestFreq,
		"ftp_rate":           this.FtpRate,
		"param_type":         this.ParamType,
		"start_time":         this.StartTime,
		"update_time":        util.GetCurrentTime(),
		"update_user":        this.UpdateUser,
		"flag":               this.Flag,
	}
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleUpdate("RPM_PAR_FTP", paramMap, whereParam)
	if nil != err {
		er := fmt.Errorf("更新资金成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 删除资金成本率信息
// by author Yeqc
// by time 2016-12-19 10:23:37
func (this Ftp) Delete() error {
	whereParam := map[string]interface{}{
		"uuid": this.UUID,
	}
	err := util.OracleDelete("RPM_PAR_FTP", whereParam)
	if nil != err {
		er := fmt.Errorf("删除资金成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}

// 新增资金成本率信息
// by author Yeqc
// by time 2016-12-19 10:23:37
func (this Ftp) Add() error {
	paramMap := map[string]interface{}{
		"product":           this.Product.ProductCode,
		"organ":             this.Organ.OrganCode,
		"currency":          this.Currency,
		"term":              this.Term,
		"rate_type":         this.RateType,
		"reprice_freq":      this.RepriceFreq,
		"rpym_type":         this.RpymType,
		"rpym_capital_freq": this.RpymCapitalFreq,
		//		"rpym_interest_freq": this.RpymInterestFreq,
		"ftp_rate":    this.FtpRate,
		"param_type":  this.ParamType,
		"start_time":  this.StartTime,
		"create_time": util.GetCurrentTime(),
		"create_user": this.CreateUser,
		"update_time": util.GetCurrentTime(),
		"update_user": this.UpdateUser,
		"flag":        this.Flag,
	}
	err := util.OracleAdd("RPM_PAR_FTP", paramMap)
	if nil != err {
		er := fmt.Errorf("新增资金成本率信息出错")
		zlog.Error(er.Error(), err)
		return er
	}
	return nil
}
