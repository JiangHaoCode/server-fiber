package frontend

import (
	"server/global"
	"server/model/frontend"
)

type FrontendImages struct{}

func (s *FrontendImages) GetImagesList() (list []frontend.ExaFile, err error) {
	err = global.DB.Model(&frontend.ExaFile{}).Where("proportion > 1.6").Order("id desc").Find(&list).Error
	return
}
