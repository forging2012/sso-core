package entity

type SysUserInfo struct {
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
	UserPhone string `json:"userPhone"`
	OrgUnitId string `json:"orgUnitId"`
}
