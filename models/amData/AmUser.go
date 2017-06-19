package amData

type AmUser struct {
	DomainId     string `json:"domain_id"`     // 域编码
	DomainName   string `json:"domain_name"`   // 域名称
	OrgUnitId    string `json:"org_unit_id"`   // 机构编码
	OrgUnitDesc  string `json:"org_unit_desc"` // 机构描述
	UserId       string `json:"user_id"`       // 用户编码
	UserName     string `json:"user_name"`     // 用户名称
	RoleIds      string `json:"role_ids"`      // 角色编码
	RoleNames    string `json:"role_names"`    // 角色名称
	LogginStatus string `json:"loggin_status"` // 用户登录状态 0:登录  1:未登录
	UserSid      string `json:"user_sid"`      // 用户SESSION_ID
}
