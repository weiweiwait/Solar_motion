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
