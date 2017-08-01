package entity

type LocalStaticEntity struct {
	Url        string `json:"url"`
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
	CreateUser string `json:"createUser"`
	ModifyTime string `json:"modifyTime"`
	ModifyUser string `json:"modifyUser"`
	Uuid       string `json:"uuid"`
}
