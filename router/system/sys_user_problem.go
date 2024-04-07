package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type ProblemRouter struct{}

func (*ProblemRouter) InitProblemRouter(Router fiber.Router) {
	problemRouter := Router.Group("problem")
	var problemApi = new(v1.UserProblem)

	problemRouter.Get("getProblemList/:id", problemApi.GetProblemSetting)
	problemRouter.Put("updateProblem", middleware.OperationRecord, problemApi.UpdateProblemSetting)

	problemRouter.Get("getIsSetting/:uid", problemApi.HasSetting)
	problemRouter.Post("verifyAnswer", middleware.OperationRecord, problemApi.VerifyAnswer)
}
