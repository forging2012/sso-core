package entity

type SsoRouteRegister struct {
	RegisterUrl string `json:"registerUrl"`
	ServiceCd   string `json:"serviceCd"`
	RemoteUrl   string `json:"remoteUrl"`
}
