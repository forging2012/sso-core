package entity

type SsoProxyStaticRoute struct {
	RegisterUrl string `json:"registerUrl"`
	RouteDesc   string `json:"routeDesc"`
	CreateTime  string `json:"createTime"`
	RemoteUrl   string `json:"remoteUrl"`
	CreateUser  string `json:"createUser"`
	ModifyUser  string `json:"modifyUser"`
	ModifyTime  string `json:"modifyTime"`
	Uuid        string `json:"uuid"`
	ServiceCd   string `json:"serviceCd"`
}
