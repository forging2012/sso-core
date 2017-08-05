package entity

type SsoSubsystemEntity struct {
	ServiceCd    string `json:"serviceCd"`
	RemoteScheme string `json:"remoteScheme"`
	RemoteHost   string `json:"remoteHost"`
	RemotePort   string `json:"remotePort"`
	CreateTime   string `json:"createTime"`
	ModifyTime   string `json:"modifyTime"`
	CreateUser   string `json:"createUser"`
	ModifyUser   string `json:"modifyUser"`
	ServiceName  string `json:"serviceName"`
	PrefixUrl    string `json:"prefixUrl"`
}
