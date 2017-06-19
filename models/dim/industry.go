package dim

import (
	"database/sql"
	"fmt"
	"github.com/lunny/log"
	"platform/dbobj"
	"strings"
	"time"
)

type Industry struct {
	UUID             string      //主键
	IndustryCode     string      //行业编号
	IndustryName     string      //行业名称
	IndustryType     string      //行业类型
	IndustryTypeDesc string      //行业类型描述
	IndustryLevel    string      //行业级别
	ParentIndustry   interface{} //父级行业
	LeafFlag         string      //是否叶子节点 0:否 1:是
	Flag             string      //生效标志 0:失效 1:生效
	Status           string      //状态
	CreateDate       time.Time   //创建时间
	CreateUser       string      //创建人
	UpdateDate       time.Time   //更新时间
	UpdateUser       string      //更新人
	ResChr1          string      //预留字段
	ResChr2          string      //预留字段
	ResChr3          string      //预留字段
}

var selectIndustrySql = `
	select 
		uuid
		,industry_code
		,industry_name
		,industry_type
		,industry_type_desc
		,industry_level
		,parent_industry
		,leaf_flag
		,flag
		,status
		,create_date
		,create_user
		,update_date
		,update_user
		,res_chr1
		,res_chr2
		,res_chr3

	from rpm_dim_industry
`

func selectIndusty(rows *sql.Rows) ([]Industry, error) {
	var one Industry
	var rst []Industry
	for rows.Next() {
		err := rows.Scan(
			&one.UUID,
			&one.IndustryCode,
			&one.IndustryName,
			&one.IndustryType,
			&one.IndustryTypeDesc,
			&one.IndustryLevel,
			&one.ParentIndustry,
			&one.LeafFlag,
			&one.Flag,
			&one.Status,
			&one.CreateDate,
			&one.CreateUser,
			&one.UpdateDate,
			&one.UpdateUser,
			&one.ResChr1,
			&one.ResChr2,
			&one.ResChr3,
		)
		if nil != err {
			log.Info("查询行业数据row.Scan()失败。", err)
			return nil, err
		}
		if " " != one.ParentIndustry.(string) {
			industry, err := SelectIndustryByCode(one.ParentIndustry.(string))
			if nil != err {
				log.Info("查询父级行业失败", err)
				return nil, err
			}

			one.ParentIndustry = industry
		}
		rst = append(rst, one)
	}
	return rst, nil
}

func SelectIndustryByCode(industryCode string) (*Industry, error) {

	sql := selectIndustrySql + ` where industry_code = :1`
	rows, err := dbobj.Default.Query(sql, industryCode)
	defer rows.Close()
	if nil != err {
		log.Info("通过行业编号查询行业纪录失败", err)
		return nil, err
	}
	rst, err := selectIndusty(rows)
	if nil != err {
		log.Info("通过行业编号查询行业数据失败", err)
		return nil, err
	}
	if 0 == len(rst) {
		return nil, fmt.Errorf("数据库没有行业编号为％s的行业纪录", industryCode)
	}
	if 1 < len(rst) {
		return nil, fmt.Errorf("数据库存在多条行业编号为％s的纪录", industryCode)
	}

	return &rst[0], nil
}

func SelectIndustryByParams(paramMap map[string]interface{}) (*[]Industry, error) {
	sql := selectIndustrySql + ` where 1=1`
	for k, v := range paramMap {
		sql += ` and ` + k + ` = '` + v.(string) + `'`
	}
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if nil != err {
		log.Info("多参数查询行业纪录失败", err)
		return nil, err
	}
	rst, err := selectIndusty(rows)
	if nil != err {
		log.Info("多参数查询行业纪录失败", err)
		return nil, err
	}
	return &rst, nil
}

func SelectIndustryTreeByTop(industryCode string) ([]*Industry, error) {

	return nil, nil
}

func DeleteIndustryByCode(industryCode string) error {
	sql := `
		delete from sys_dim_industry where industry_code = :1
	`
	err := dbobj.Default.Exec(sql, industryCode)
	if nil != err {
		log.Info("通过行业编号删除行业失败", err)
		return err
	}
	return nil

}

func DeleteIndustryByUuids(uuids []string) error {
	sql := `
		delete from rpm_dim_industry where uuid in (`
	for _, v := range uuids {
		sql += `'` + v + `',`
	}
	strings.TrimRight(sql, ",")
	sql += `)`
	err := dbobj.Default.Exec(sql)
	if nil != err {
		log.Info("删除行业纪录失败", err)
		return err
	}
	return nil
}

func UpdateIndustry(setMsg, whereMsg map[string]interface{}) error {
	sql := `
		update rpm_dim_industry set 
	`
	for k, v := range setMsg {
		sql += k + ` = '` + v.(string) + `',`
	}
	strings.TrimRight(sql, ",")
	sql += ` where 1=1`
	for k, v := range whereMsg {
		sql += ` and ` + k + ` = '` + v.(string) + `'`
	}
	err := dbobj.Default.Exec(sql)
	if nil != err {
		log.Info("更新行业纪录出错", err)
		return err
	}
	return nil
}

func InsertIndustry(industry Industry) error {
	sql := `
		insert into rpm_dim_industry(
				 uuid
				,industry_code
				,industry_name
				,industry_type
				,industry_typeDesc
				,industry_level
				,parent_industry
				,leaf_flag
				,flag
				,status
				,create_date
				,create_user
				,update_date
				,update_user
				,res_chr1
				,res_chr2
				,res_chr3
			)
		values(
				sys_guid(),:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11,:12,:13,:14,:15,:16
			)
	`
	err := dbobj.Default.Exec(sql,
		industry.IndustryCode,
		industry.IndustryName,
		industry.IndustryType,
		industry.IndustryTypeDesc,
		industry.IndustryLevel,
		industry.ParentIndustry,
		industry.LeafFlag,
		industry.Flag,
		industry.Status,
		industry.CreateDate,
		industry.CreateUser,
		industry.UpdateDate,
		industry.UpdateUser,
		industry.ResChr1,
		industry.ResChr2,
		industry.ResChr3,
	)
	if nil != err {
		log.Info("新增行业纪录失败", err)
		return err
	}
	return nil
}
