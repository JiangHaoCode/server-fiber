package mobile

import (
	"server/global"
	"server/model/mobile"
)

type MobileRegisterService struct{}

func (*MobileRegisterService) Register(data mobile.Register) error {
	return global.DB.Create(&data).Error
}
