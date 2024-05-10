package types

type UserRegisterReq struct {
	UserName    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	QQ          string `form:"qq" json:"qq"`
}
type UserLoginReq struct {
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Password    string `form:"password" json:"password"`
}
type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
type UserLoginReply struct {
	PhoneNumber string `form:"password" json:"password"`
	Integral    int    `form:"integral" json:"integral"`
}
type UserAvatar struct {
	Avatar string `form:"avatar" json:"avatar"`
}
type UserSendEmail struct {
	QQ string `form:"qq" json:"qq"`
}
type UserSendCode struct {
	Password string `form:"password" json:"password"`
	QQ       string `form:"qq" json:"qq"`
	Code     string `form:"code" json:"code"`
}
type UserNameUpdate struct {
	UserName string `form:"username" json:"username"`
}
