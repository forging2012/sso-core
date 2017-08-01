package entity

type SysSecUser struct {
	UserId           string `json:"userId"`
	UserPasswd       string `json:"userPassword"`
	StatusId         int    `json:"statusId"`
	ContinueErrorCnt int    `json:"continueErrorCnt"`
}
