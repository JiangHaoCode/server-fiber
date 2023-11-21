package service

import (
	"server/service/app"
	"server/service/example"
	"server/service/frontend"
	"server/service/mobile"
	"server/service/order"
	"server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	OrderServiceGroup    order.ServiceGroup
	AppServiceGroup      app.ServiceGroup
	FrontendServiceGroup frontend.ServiceGroup
	MobileServiceGroup   mobile.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
