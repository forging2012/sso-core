package entity

type SsoRouteRegisterApi struct {
	ServiceCd   string `json:"serviceCd"`
	RegisterUrl string `json:"registerUrl"`
	RouteDesc   string `json:"routeDesc"`
	CreateTime  string `json:"createTime"`
	CreateUser  string `json:"createUser"`
	ModifyUser  string `json:"modifyUser"`
	ModifyTime  string `json:"modifyTime"`
	Uuid        string `json:"uuid"`
	RemoteUrl   string `json:"remoteUrl"`
}
