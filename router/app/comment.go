package app

import (
	v1 "server-fiber/api/v1/app"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type CommentRouter struct {
}

// InitCommentRouter 初始化 Comment 路由信息
func (s *CommentRouter) InitCommentRouter(Router fiber.Router) {
	commentRouter := Router.Group("comment")
	var commentApi = new(v1.CommentApi)

	commentRouter.Post("createComment", middleware.OperationRecord, commentApi.CreateComment)             // 新建Comment
	commentRouter.Delete("DeleteComment", middleware.OperationRecord, commentApi.DeleteComment)           // 删除Comment
	commentRouter.Delete("DeleteCommentByIds", middleware.OperationRecord, commentApi.DeleteCommentByIds) // 批量删除Comment
	commentRouter.Put("updateComment", middleware.OperationRecord, commentApi.UpdateComment)              // 更新Comment
	commentRouter.Put("pariseComment", middleware.OperationRecord, commentApi.PutLikeItOrDislike)         //点赞

	commentRouter.Get("findComment", commentApi.FindComment)               // 根据ID获取Comment
	commentRouter.Get("getCommentList", commentApi.GetCommentList)         // 获取Comment列表
	commentRouter.Get("getCommentTreeList", commentApi.GetCommentTreeList) //  获取Comment Tree列表

}
