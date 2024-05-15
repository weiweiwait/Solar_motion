package blog

import (
	"Solar_motion/pkg/utils/ctl"
	"Solar_motion/repository/dao"
	"Solar_motion/repository/model"
	"Solar_motion/types"
	"context"
	"sync"
)

var BlogSrvIns *BlogSrv
var BlogSrvOnce sync.Once

type BlogSrv struct {
}

func GetBlogSrv() *BlogSrv {
	BlogSrvOnce.Do(func() {
		BlogSrvIns = &BlogSrv{}
	})
	return BlogSrvIns
}
func (s *BlogSrv) UserPushBlog(ctx context.Context, req *types.Blog) (resp interface{}, err error) {
	userDao := dao.NewBlogDao(ctx)
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	blog := &model.Blog{
		UserId:   u.Id,
		Title:    req.Title,
		Contexts: req.Contexts,
	}
	err = userDao.CreateBlog(blog)
	if err != nil {
		return nil, err
	}
	return
}
