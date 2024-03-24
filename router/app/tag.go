package app

import (
	v1 "server-fiber/api/v1"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type TagRouter struct{}

// InitTagRouter 初始化 Tag 路由信息
func (s *TagRouter) InitTagRouter(Router fiber.Router) {
	var tagApi = v1.ApiGroupApp.AppApiGroup.TagApi

	tagRouter := Router.Group("tag")
	tagRouter.Post("createTag", middleware.OperationRecord, tagApi.CreateTag)             // 新建Tag
	tagRouter.Delete("deleteTag/:id", middleware.OperationRecord, tagApi.DeleteTag)       // 删除Tag
	tagRouter.Delete("deleteTagByIds", middleware.OperationRecord, tagApi.DeleteTagByIds) // 批量删除Tag
	tagRouter.Put("updateTag", middleware.OperationRecord, tagApi.UpdateTag)              // 更新Tag

	tagRouter.Get("findTag/:id", tagApi.FindTag)   // 根据ID获取Tag
	tagRouter.Get("getTagList", tagApi.GetTagList) // 获取Tag列表

}
