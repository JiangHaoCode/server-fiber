package app

import (
	"server/global"
	"server/model/app"
	appReq "server/model/app/request"
	"server/model/common/request"

	"strings"
	"time"
)

type ArticleService struct{}

// CreateArticle
func (*ArticleService) CreateArticle(article app.Article) (err error) {
	err = global.DB.Create(&article).Error
	return
}

// DeleteArticle delete
func (*ArticleService) DeleteArticle(id uint) (err error) {
	err = global.DB.Delete(&app.Article{}, id).Error
	return
}

// delete by ids
func (*ArticleService) DeleteArticleByIds(ids request.IdsReq) (err error) {
	var articleIds *[]app.Article = &[]app.Article{}
	err = global.DB.Delete(articleIds, "id in ?", ids.Ids).Error
	return
}

// update
func (*ArticleService) UpdateArticle(article app.Article) (err error) {
	var articleDetail app.Article

	err = global.DB.Where("id = ?", article.ID).First(&articleDetail).Error
	if err != nil {
		return
	}
	tx := global.DB.Begin()
	tx = tx.Model(&app.ArticleTag{})
	err = tx.Delete("article_id = ?", article.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	err = global.DB.Save(&article).Error
	return err
}

// getDetail by id
func (*ArticleService) GetArticle(id uint) (article app.Article, err error) {
	err = global.DB.Preload("Tags").Preload("User").Where("id = ?", id).First(&article).Error
	return
}

// getList

func (*ArticleService) GetArticleInfoList(info appReq.ArticleSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&app.Article{})
	var articles []app.Article
	aritlceList := make([]int64, 0)
	if info.TagId != "" {
		dbTag := global.DB.Model(&app.ArticleTag{})
		var articleTag []app.ArticleTag
		err = dbTag.Where("tag_id = ?", info.TagId).Find(&articleTag).Error
		if err != nil {
			return
		}
		for _, item := range articleTag {
			aritlceList = append(aritlceList, item.ArticleId)
		}
		db = db.Where("id in ?", aritlceList)
	}

	if info.Title != "" {
		db = db.Where("title like ?", strings.Join([]string{"%", info.Title, "%"}, ""))
	}
	if info.IsImportant != 0 {
		db = db.Where("is_important = ?", info.IsImportant)
	}
	// db.SetupJoinTable(&articles, "article_tag")
	// if info.TagId != 0 {
	// 	db = db.Where("tag_id = ?", info.TagId)
	// }
	// db.Association("Tags")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//
	err = db.Limit(limit).Offset(offset).Order("id desc").Preload("User").Preload("Tags").Find(&articles).Error
	return articles, total, err
}

// 批量更新
func (*ArticleService) PutArticleByIds(ids request.IdsReq) (err error) {
	err = global.DB.Model(&app.Article{}).Where("id in ?", ids.Ids).Update("is_important", 2).Error
	return
}

func (*ArticleService) GetArticleReading() (count int64, err error) {
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 99, t.Location())
	err = global.DB.Model(&app.Ip{}).Where("created_at > ? and created_at < ?", startTime, endTime).Count(&count).Error
	return
}
