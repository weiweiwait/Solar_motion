package types

type ManagerRegisterReq struct {
	UserName    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}

type ManagerLoginReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Password    string `form:"password" json:"password"`
}
type ManagerLoginReply struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}
type ManagerTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
type ManagerDeleteRep struct {
	Username string `form:"username" json:"username"`
}
