package types

type ManagerRegisterReq struct {
	UserName    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}
