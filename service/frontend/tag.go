package frontend

import (
	// "encoding/json"
	"errors"
	"server/global"
	"server/model/frontend"
	"strconv"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type FrontendTag struct{}

func (s *FrontendTag) GetTagList(c *fiber.Ctx) (list []frontend.Tag, err error) {
	var tagListStr string
	tagListStr, err = global.REDIS.Get(c, "tag-list").Result()
	var cacheTime = global.CONFIG.Cache.Time
	if err == redis.Nil {
		db := global.DB.Model(&frontend.Tag{})
		err = db.Order("id desc").Find(&list).Error
		if err != nil {
			return list, err
		}
		strList, _ := json.Marshal(list)
		err := global.REDIS.Set(c, "tag-list", strList, time.Duration(cacheTime)*time.Second).Err()
		if err != nil {
			return list, errors.New("redis 存储失败， 请查看redis服务 配置")
		}
		return list, nil
	} else if err != nil {
		return list, err
	} else {
		if tagListStr != "" {
			err = json.Unmarshal([]byte(tagListStr), &list)
			return list, err
		}
	}
	return list, err
}

func (s *FrontendTag) GetTagArticle(tagId int, c *gin.Context) (tagArticle frontend.Tag, err error) {
	tagArticleStr := ""
	tagArticleStr, err = global.REDIS.Get(c, "tag-"+strconv.Itoa(tagId)).Result()
	if err == redis.Nil {
		db := global.DB.Model(&frontend.Tag{})
		err = db.Where("id = ?", tagId).Order("id desc").Preload("Articles").First(&tagArticle).Error
		if err != nil {
			return tagArticle, err
		}
		err := global.REDIS.Set(c, "tag-list", tagArticle, 1*time.Hour).Err()
		if err != nil {
			return tagArticle, errors.New("redis 存储失败， 请查看redis服务 配置")
		}
		return tagArticle, nil
		// return tagArticle, err
	} else if err != nil {
		return tagArticle, err
	} else {
		err = json.Unmarshal([]byte(tagArticleStr), &tagArticle)
	}

	return tagArticle, err
}
