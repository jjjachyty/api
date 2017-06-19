package sys

import (
	"fmt"
	"platform/dbobj"

	"pccqcpa.com.cn/components/zlog"
)

type User struct {
	Uuid     string
	UserName string
	Password string
	Organ    Organ
}

func SelectUserByParam(userName, password string) (*User, error) {
	sql := `
		select uuid,user_name,password from sys_sec_user where user_name = :1 and password = :2
	`
	rows, err := dbobj.Default.Query(sql, userName, password)
	if nil != err {
		zlog.Error("查询用户表错误", err)
		return nil, err
	}

	var rst User
	for rows.Next() {
		err = rows.Scan(&rst.Uuid, &rst.UserName, &rst.Password)
	}
	if nil != err {
		zlog.Error("查询用户表rows.scan()错误", err)
		return nil, err
	}

	if "" == rst.Uuid {
		zlog.Error("未查询到用户信息", nil)
		return nil, fmt.Errorf("未查询到用户信息")
	}
	return &rst, nil

}

func GetCurrentUser() *User {
	var organ = Organ{
		OrganCode: "panchina001",
	}
	var user = User{
		Uuid:     "ABCDE",
		UserName: "Jason",
		Organ:    organ,
	}
	return &user
}
