package functionIndicators

import (
	"fmt"
	// "time"
	"errors"

	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
	"pccqcpa.com.cn/components/zlog"
)

var lgdBasel par.LgdBasel

type LgdRate struct {
}

var lnMort ln.LnMort
var lnGuarante ln.LnGuarante

// 计算违约损失率
// 参数1:业务单号
// 参数2:贷款金额
func (this *LgdRate) Calulate(paramMap map[string]interface{}) (float64, error) {
	v, ok := paramMap["principal"]
	if !ok {
		err := fmt.Errorf("计算违约损失率未传参数贷款金额")
		zlog.Errorf(err.Error(), err)
		return -1, err
	}
	var principal float64 = v.(float64)

	// 查询押品数组按照押品类型排序
	v, ok = paramMap["business_code"]
	if !ok {
		err := fmt.Errorf("计算违约损失率未传参数业务单号")
		zlog.Errorf(err.Error(), err)
		return -1, err
	}
	var params = map[string]interface{}{
		"business_code": v.(string),
		"sort":          "mortgage.mortgage_basel",
	}
	lnMorts, err := lnMort.Find(params)
	if nil != err {
		zlog.Error("计算违约损失率查询押品数据报错", err)
		return -1, err
	}

	// 查询保证人
	params = map[string]interface{}{
		"business_code": v.(string),
	}
	lnGuarantes, err := lnGuarante.Find(params)
	if nil != err {
		er := errors.New("计算违约损失率时查询保证人出错")
		zlog.Errorf(er.Error(), err)
		return -1, er
	}
	var havaLnGuarante bool = false
	if 0 < len(lnGuarantes) {
		havaLnGuarante = true
	}
	// 计算lgd值
	lgd, err := this.calulateLgd(lnMorts, havaLnGuarante, principal)
	if nil != err {
		zlog.Error("计算LGD出错", err)
		return -1, err
	}
	return lgd, nil
}

// 参数1:押品数组（已排序）
// 参数2:贷款金额
// 计算lgd值
// 逻辑：
// 风险暴露金额 ead ＝ 贷款金额
// 违约损失率 lgd ＝ 0
// mortList := 根据业务单号查询押品信息根据押品类型排序（金融质押品、应收账款、商用／民用住房、其他抵质押品、保证）
// for mort := range mortList{
//		抵质押水平 ＝ 押品金额／ead
// 		if 抵质押水平 > 最高抵质押水平  {
//			lgd += 最高抵质押水平
// 			ead -= 押品金额／最高抵质押水平
//		} else if 最低抵质押水平 < 抵质押水平 ｛
//			lgd += (押品金额／最高抵质押水平 ＊ 最低违约损失率)／ 贷款金额
//			ead ＝ ead - 押品金额／最高抵质押水平
// 		}
// ｝
// 判断是否有保证人
// 如果有保证人,则算保证类型，默认全额覆盖 违约损失率40%
func (l *LgdRate) calulateLgd(lnMorts []*ln.LnMort, havaGuarante bool, principal float64) (float64, error) {
	var ead float64 = principal
	var lgd float64 = 0
	for _, mort := range lnMorts {
		pledgeRadio := mort.MortgageValue / ead   //抵质押水平
		pledgeType := mort.Mortgage.MortgageBasel //抵质押类型
		value := util.GetCacheByCacheName(util.RPM_LGD_CACHE, pledgeType)
		if nil == value {
			er := fmt.Errorf("押品类型映射不对")
			zlog.Error(er.Error(), er)
			return -1, er
		}
		var lgdBasel *par.LgdBasel = value.(*par.LgdBasel)
		var ex = ead - mort.MortgageValue/lgdBasel.HighPledgeRadio
		if ex >= 0 {
			// 计算违约损失率
			var err error
			var formual string = ""
			if pledgeRadio > lgdBasel.HighPledgeRadio {
				formual = fmt.Sprintf("%f", lgd) + `+` + fmt.Sprintf("%f", lgdBasel.HighPledgeRadio)
				lgd, err = util.Calculate(formual)
				fmt.Println("－－－－－－－违约损失率：－－－－－－－－", formual)
				if nil != err {
					zlog.Error("lgd计算错误", err)
				}
				fmt.Println("超额抵质押违约损失率计算", formual)
				// 计算违约损失率风险暴露计算
				ead, err = l.calulateEad(ead, *mort, *lgdBasel)
				if nil != err {
					return -1, err
				}
			} else if pledgeRadio >= lgdBasel.LowPledgeRadio {
				// lgd += mort.MortgageValue / lgdBasel.HighPledgeRadio * lgdBasel.LowLgd / principal
				// ead -= mort.MortgageValue / lgdBasel.HighPledgeRadio
				formual = fmt.Sprintf("%f", lgd) + `+` +
					fmt.Sprintf("%f", mort.MortgageValue) + `/` +
					fmt.Sprintf("%f", lgdBasel.HighPledgeRadio) + `*` +
					fmt.Sprintf("%f", lgdBasel.LowLgd) + `/` +
					fmt.Sprintf("%f", principal)
				lgd, err = util.Calculate(formual)
				fmt.Println("－－－－－－－违约损失率：－－－－－－－－", formual)
				if nil != err {
					zlog.Error("lgd计算错误", err)
				}
				fmt.Println("违约损失率计算", formual)
				// 计算违约损失率风险暴露计算
				ead, err = l.calulateEad(ead, *mort, *lgdBasel)
				if nil != err {
					return -1, err
				}
			}

		} else {
			// lgd += ead * lgdBasel.LowLgd / principal
			var formual = fmt.Sprintf("%f", lgd) + `+` +
				fmt.Sprintf("%f", ead) + `*` +
				fmt.Sprintf("%f", lgdBasel.LowLgd) + `/` +
				fmt.Sprintf("%f", principal)
			fmt.Println("－－－－－－－违约损失率：－－－－－－－－", formual)
			var err error
			lgd, err = util.Calculate(formual)
			if nil != err {
				zlog.Error("lgd计算错误", err)
			}
		}
		if ead <= 0 || ex <= 0 {
			return lgd, nil
		}
	}
	if havaGuarante {
		value := util.GetCacheByCacheName(util.RPM_LGD_CACHE, "5")
		if nil == value {
			er := fmt.Errorf("押品类型映射不对")
			zlog.Error(er.Error(), er)
			return -1, er
		}
		var lgdBasel *par.LgdBasel = value.(*par.LgdBasel)
		var formual = fmt.Sprintf("%f", lgd) + `+` +
			fmt.Sprintf("%f", ead) + `*` +
			fmt.Sprintf("%f", lgdBasel.LowLgd) + `/` +
			fmt.Sprintf("%f", principal)
		lgd, err := util.Calculate(formual)
		if nil != err {
			zlog.Error("lgd计算保证人错误", err)
		}
		return lgd, nil
	}
	lgd += ead * util.CREDIT_LOW_LGD / principal
	return lgd, nil
}

func (l LgdRate) calulateEad(ead float64, mort ln.LnMort, lgdBasel par.LgdBasel) (float64, error) {
	// 计算违约损失率风险暴露计算
	var formual = fmt.Sprintf("%f", ead) + `-` +
		fmt.Sprintf("%f", mort.MortgageValue) + `/` +
		fmt.Sprintf("%f", lgdBasel.HighPledgeRadio)
	ead, err := util.Calculate(formual)
	fmt.Println("－－－－－－－违约损失率风险暴露计算：－－－－－－－－", formual)
	if nil != err {
		zlog.Error("风险暴露计算错误", err)
		return -1, err
	}
	fmt.Println("风险暴露计算", formual)
	return ead, nil
}
