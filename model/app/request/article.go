package request

import (
	"server/model/app"
	"server/model/common/request"
)

type ArticleSearch struct {
	app.Article
	TagId string `json:"tag_id" form:"tag_id"`
	request.PageInfo
}
