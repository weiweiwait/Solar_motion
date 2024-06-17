package types

type Blog struct {
	Title    string `form:"title" json:"title"`
	Contexts string `form:"contexts" json:"contexts"`
}
type ImagesResp struct {
	BlogId uint   `form:"blog_id" json:"blog_id"`
	Image  string `form:"image" json:"image"`
}

type ImagesReq struct {
	BlogId string `form:"blog_id" json:"blog_id"`
}
type OtherImagesReq struct {
	BlogId string `form:"blog_id" json:"blog_id"`
	UserId string `form:"user_id" json:"user_id"`
}

type BlogService struct {
	BlogId string `form:"blogId"`
	//Pictures string `form:"pictures"` // 图片
	BlogTitle string `form:"blogTitle"`
	Email     string `form:"email"`
	Content   string `form:"content"`
	Location  string `form:"location"`
}
