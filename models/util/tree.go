package util

import (
	"pccqcpa.com.cn/components/zlog"

	"pccqcpa.com.cn/app/rpm/api/util"
)

type Tree struct {
	Label  string `json:"name"`     //名称
	Data   string `json:"id"`       //id
	NoLeaf int    `json:"isParent"` //是否有子节点（有子节点：1 没有子节点：0）
}

// 查询树通用方法
// structName ：查询结构体名称
// 参数topCode ：父code
// 参数tableName : 表名称
func SelectTree(structName, topCode, tableName string) util.RstMsg {
	zlog.AppOperateLog("", "SelectTree", zlog.SELECT, nil, nil, "查询树状结构")
	var msg string = "查询树状结构成功"
	searchLike := make([]map[string]interface{}, 0)
	searchLike = append(searchLike, map[string]interface{}{"key": "flag", "type": "eq", "value": "1"})
	searchLike = append(searchLike, map[string]interface{}{"key": "parent_" + structName, "type": "eq", "value": topCode})

	paramMap := map[string]interface{}{
		"searchLike": searchLike,
	}
	cols := map[string]string{
		structName + `_code`: "''",
		structName + `_name`: "''",
		"leaf_flag":          "'0'",
	}
	colsSort := []string{
		structName + `_code`,
		structName + `_name`,
		"leaf_flag",
	}
	rows, err := FindRows(tableName, cols, colsSort, paramMap)
	if nil != err {
		msg = "查询树状结构失败"
		zlog.Error(msg, err)
		return util.ErrorMsg(msg, err)
	}
	defer rows.Close()
	var rst []*Tree
	for rows.Next() {
		var one Tree
		err := util.OracleScan(rows, []interface{}{&one.Data, &one.Label, &one.NoLeaf})
		if nil != err {
			msg = "查询树状结构rows.Scan()失败"
			zlog.Error(msg, err)
			return util.ErrorMsg(msg, err)
		}
		rst = append(rst, &one)
	}
	return util.SuccessMsg(msg, rst)

}
