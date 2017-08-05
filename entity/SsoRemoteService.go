package entity

type SsoRemoteService struct {
	ServiceCd    string `json:"serviceCd"`
	RemoteScheme string `json:"remoteScheme"`
	RemoteHost   string `json:"remoteHost"`
	RemotePort   string `json:"remotePort"`
	PrefixUrl    string `json:"prefixUrl"`
}
