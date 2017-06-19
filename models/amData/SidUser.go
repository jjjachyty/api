package amData

import (
	"time"
)

type SidUser struct {
	Sid        string
	UserId     string
	UserName   string
	BranchCode string
	RoleIds    string
	Status     string
	LoginTime  time.Time
	LastOpTime time.Time
	ExpTime    time.Duration
}

func (sidUser *SidUser) AmUserToSidUser(amUser AmUser) {
	sidUser.Sid = amUser.UserSid
	sidUser.UserId = amUser.UserId
	sidUser.UserName = amUser.UserName
	sidUser.BranchCode = amUser.OrgUnitId
	sidUser.LastOpTime = time.Now()
	sidUser.RoleIds = amUser.RoleIds
	sidUser.ExpTime = 3
}
