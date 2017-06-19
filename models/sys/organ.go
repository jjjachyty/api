package sys

import (
	"database/sql"
	"fmt"
	"platform/dbobj"
	"strings"
	"time"

	"github.com/lunny/log"
	// "pccqcpa.com.cn/app/rpm/api/models/util"
)

type Organ struct {
	// util.BaseEntity           //基础字段
	UUID        string    //主键
	OrganCode   string    //机构码值
	OrganName   string    //机构名称
	OrganLevel  string    //机构层级
	ParentOrgan string    //父级机构
	LeafFlag    string    //是否叶子节点
	Flag        string    //生效标志
	Status      string    //状态
	CreateTime  time.Time //创建时间
	CreateUser  string    //创建人
	UpdateTime  time.Time //更新时间
	UpdateUser  string    //更新人
}
type NullOrgan struct {
	UUID        sql.NullString
	OrganCode   sql.NullString
	OrganName   sql.NullString
	OrganLevel  sql.NullString
	ParentOrgan sql.NullString
	LeafFlag    sql.NullString
	Flag        sql.NullString
	Status      sql.NullString
	CreateTime  time.Time
	CreateUser  sql.NullString
	UpdateTime  time.Time
	UpdateUser  sql.NullString
}

var selectOrganSql string = `
	select 
		 UUID
		,coalesce(organ_code,' ')
		,coalesce(organ_name,' ')
		,coalesce(organ_level,' ')
		,coalesce(parent_organ,' ')
		,coalesce(leaf_flag,' ')
		,coalesce(flag,' ')
		,coalesce(status,' ')
		,coalesce(create_time,trunc(sysdate))
		,coalesce(create_user,' ')
		,coalesce(update_time,create_time)
		,coalesce(update_user,create_user)
	from sys_sec_organ
`

func selectOrgan(rows *sql.Rows) ([]*Organ, error) {
	var rst []*Organ
	for rows.Next() {
		var one Organ
		//var parentOrgan Organ
		err := rows.Scan(
			&one.UUID,
			&one.OrganCode,
			&one.OrganName,
			&one.OrganLevel,
			&one.ParentOrgan,
			&one.LeafFlag,
			&one.Flag,
			&one.Status,
			&one.CreateTime,
			&one.CreateUser,
			&one.UpdateTime,
			&one.UpdateUser,
		)
		//one.ParentOrgan = &parentOrgan
		if nil != err {
			log.Info("通过机构编码查询机构号rows.Scan()错误", err)
		}

		rst = append(rst, &one)
	}
	return rst, nil
}

func SelectOrganByOrgCode(organCode string) (*Organ, error) {
	sql := selectOrganSql + `
		where organ_code = :1
	`
	rows, err := dbobj.Default.Query(sql, organCode)
	defer rows.Close()
	if nil != err {
		log.Info("通过机构编码查询机构错误，请检查机构号:"+organCode+"是否正确", err)
		return nil, err
	}

	rst, err := selectOrgan(rows)
	if nil != err {
		log.Info("通过机构号查询机构rows.Scan()错误", err)
	}
	if 1 != len(rst) {
		log.Info("机构编码：" + organCode + "数据库没有纪录或者有多条纪录，请检查数据")
		return nil, fmt.Errorf("机构编码：%s数据库没有纪录或者有多条纪录，请检查数据", organCode)
	}
	return rst[0], nil
}

func SelectOrganByParams(paramMap map[string]interface{}) ([]*Organ, error) {
	sql := selectOrganSql + `where 1=1`
	for k, v := range paramMap {
		sql += `and ` + k + `= '` + v.(string) + `'`
	}
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if nil != err {
		log.Info("多参数查询机构出错，请核对信息后查询", err)
		return nil, nil
	}
	rst, err := selectOrgan(rows)
	if nil != err {
		log.Info("多参数查询机构纪录错误", err)
		return nil, err
	}
	return rst, nil
}

func InsertOrgan(organ Organ) error {
	sql := `
		insert into sys_sec_organ(
				UUID,
				organ_code,
				organ_name,
				organ_level,
				parent_organ,
				leaf_flag,
				flag,
				status,
				create_time,
				create_user,
				update_time,
				update_user
			)
		values(
				sys_guid(),:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11,:12
			)
	`
	err := dbobj.Default.Exec(sql,
		organ.OrganCode,
		organ.OrganName,
		organ.OrganLevel,
		organ.ParentOrgan,
		organ.LeafFlag,
		organ.Flag,
		organ.Status,
		organ.CreateTime,
		organ.CreateUser,
		organ.UpdateTime,
		organ.UpdateUser,
	)
	if nil != err {
		log.Info("新增机构信息出错，请检查数据", err)
		return err
	}
	return nil
}

func DeleteOrganByOrganCode(organCode []string) error {
	sql := `
		delete from sys_sec_organ where organ_code = :1
	`
	err := dbobj.Default.Exec(sql, organCode)
	if nil != err {
		log.Info("删除机构数据出错", err)
		return err
	}
	return nil
}

func DeleteOrganByUUIDs(UUIDs []string) error {
	sql := `
		delete from sys_sec_organ where UUID in (`
	for _, v := range UUIDs {
		sql += `'` + v + `',`
	}
	strings.TrimRight(sql, ",")
	sql += ")"

	err := dbobj.Default.Exec(sql)
	if nil != err {
		log.Info("删除机构信息失败", err)
		return err
	}
	return nil
}

func UpdateOrgan(setMsg, whereMsg map[string]interface{}) error {
	sql := `
		update sys_sec_organ set 
	`
	for k, v := range setMsg {
		sql += k + `=` + v.(string) + `,`
	}
	strings.TrimRight(sql, ",")

	sql += `where 1=1 `
	for k, v := range whereMsg {
		sql += `and ` + k + `=` + v.(string)
	}

	err := dbobj.Default.Exec(sql)
	if nil != err {
		log.Info("更新机构信息失败", err)
		return err
	}
	return nil
}
