package dao

import (
	"Solar_motion/repository/es"
	"Solar_motion/repository/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

const pageSize = 10

type BlogDao1 struct {
	es  *elastic.Client
	ctx context.Context
}

func NewBlogDao1(ctx context.Context) *BlogDao1 {
	return &BlogDao1{es: es.NewEs(), ctx: ctx}
}

// IndexBlog 索引blog
func (dao *BlogDao1) IndexBlog(blog *model.Blog1) (*elastic.IndexResponse, error) {
	return dao.es.Index().Index("blog").BodyJson(blog).Do(dao.ctx)
}

func (dao *BlogDao1) UploadBlogPictures(id string, urls []string) error {
	_, err := dao.es.Update().Index("blog").Id(id).Doc(map[string]interface{}{"pictures": urls}).Do(dao.ctx)
	return err
}

func (dao *BlogDao1) SearchByKeyWord(keyword string, page int) (err error, blogs []*model.Blog1, highlights []string) {
	blogTitleHighlight := elastic.NewHighlighterField("blogTitle")
	contentHighlight := elastic.NewHighlighterField("content")
	contentHighlight.FragmentSize(5)
	blogTitleHighlight.PreTags("<b>")
	blogTitleHighlight.PostTags("</b>")
	contentHighlight.PreTags("<b>")
	contentHighlight.PostTags("</b>")
	highlight := elastic.NewHighlight().Fields(blogTitleHighlight, contentHighlight)
	boolQuery := elastic.NewBoolQuery().
		Should(elastic.NewQueryStringQuery("blogTitle:"+keyword), elastic.NewQueryStringQuery("content:"+keyword))
	var result *elastic.SearchResult
	result, err = dao.es.
		Search("blog").
		Highlight(highlight).
		Size(pageSize).
		From((page - 1) * pageSize).
		Size(pageSize).
		Query(boolQuery).
		Do(context.TODO())
	if err != nil {
		return
	}
	for _, hit := range result.Hits.Hits {
		blog := &model.Blog1{}
		if err = json.Unmarshal([]byte(hit.Source), blog); err != nil {
			return
		}
		blogs = append(blogs, blog)
		if s := hit.Highlight["blogTitle"]; s != nil {
			highlights = append(highlights, s[0])
		} else {
			highlights = append(highlights, hit.Highlight["content"][0])
		}
		highlights = append(highlights)
	}
	return
}
func (dao *BlogDao1) GetBlogList(way string, page int) ([]*model.Blog1, error) {
	var blogs []*model.Blog1
	searchResult, err := dao.es.
		Search("blog").
		From((page-1)*pageSize).
		Size(pageSize).
		Sort(way, false). //根据way降序排序
		Do(context.TODO())

	if err != nil {
		return blogs, err
	}
	for _, hit := range searchResult.Hits.Hits {
		blog := &model.Blog1{}
		if err := json.Unmarshal([]byte(hit.Source), blog); err != nil {
			return blogs, err
		}
		blog.BlogId = hit.Id
		blogs = append(blogs, blog)
	}
	return blogs, nil
}
