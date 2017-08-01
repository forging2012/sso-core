package dto

type AuthDto struct {
	Username      string `json:"username"`
	Password      string `json:"-"`
	TargetSystem  string `json:"targetSystem"`
	EffectiveTime string `json:"effectiveTime"`
	RetCode       int    `json:"retCode"`
	RetMsg        string `json:"retMsg"`
}
