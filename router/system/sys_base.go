package system

import (
	v1 "server-fiber/api/v1/system"
	"server-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router fiber.Router) {
	baseRouter := Router.Group("base")
	baseApi := new(v1.BaseApi)

	baseRouter.Post("login", baseApi.Login)
	baseRouter.Get("captcha", middleware.NeedInit, baseApi.Captcha)

}
