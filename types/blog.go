package types

type Blog struct {
	Title    string `form:"title" json:"title"`
	Contexts string `form:"contexts" json:"contexts"`
}
type Images struct {
	BlogId uint   `form:"blog_id" json:"blog_id"`
	Image  string `form:"image" json:"image"`
}
