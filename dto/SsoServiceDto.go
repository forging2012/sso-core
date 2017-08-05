package dto

type SsoServiceDto struct {
	RegisterUrl  string `json:"registerUrl"`
	ServiceCd    string `json:"serviceCd"`
	RemoteUrl    string `json:"remoteUrl"`
	RemoteScheme string `json:"remoteScheme"`
	RemoteHost   string `json:"remoteHost"`
	RemotePort   string `json:"remotePort"`
	PrefixUrl    string `json:"prefixUrl"`
}
