package common

import (
	// "pccqcpa.com.cn/components/zlog"

	"github.com/lunny/tango"

	modelsUtil "pccqcpa.com.cn/app/rpm/api/models/util"
	"pccqcpa.com.cn/app/rpm/api/util"
)

type TreeAction struct {
	tango.Json
	tango.Ctx
}

var tableNames = map[string]string{
	"organ":    "SYS_SEC_ORGAN T",    //机构表
	"industry": "RPM_DIM_INDUSTRY T", //行业表
	"product":  "RPM_DIM_PRODUCT T",  //产品表
	"mortgage": "RPM_DIM_MORTGAGE T", //押品表
}

// url /api/rpm/tree/(:structName)?code=?&topCode=?
func (treeAction *TreeAction) Get() util.RstMsg {
	structName := treeAction.Param(":structName")
	topCode := treeAction.Form("topCode")
	if "" == topCode {
		topCode = treeAction.Form("code")
	}
	if "" == topCode {
		topCode = util.TREE_TOP_CODE
	}

	return modelsUtil.SelectTree(structName, topCode, tableNames[structName])
}
