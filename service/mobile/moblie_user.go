package mobile

import (
	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/mobile"
	mobileReq "server-fiber/model/mobile/request"
)

type MobileUserService struct{}

// CreateMobileUser 创建MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) CreateMobileUser(mobileUser mobile.MobileUser) (err error) {
	err = global.DB.Create(&mobileUser).Error
	return err
}

// DeleteMobileUser 删除MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) DeleteMobileUser(id uint) (err error) {
	err = global.DB.Delete(&mobile.MobileUser{}, id).Error
	return err
}

// DeleteMobileUserByIds 批量删除MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) DeleteMobileUserByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]mobile.MobileUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMobileUser 更新MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) UpdateMobileUser(mobileUser mobile.MobileUser) (err error) {
	err = global.DB.Save(&mobileUser).Error
	return err
}

// GetMobileUser 根据id获取MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) GetMobileUser(id uint) (mobileUser mobile.MobileUser, err error) {
	err = global.DB.Where("id = ?", id).First(&mobileUser).Error
	return
}

// GetMobileUserInfoList 分页获取MobileUser记录
// Author [jianghao](https://github.com/JiangHaoCode)
func (mobileUserService *MobileUserService) GetMobileUserInfoList(info mobileReq.MobileUserSearch) (list []mobile.MobileUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&mobile.MobileUser{})
	if info.Username != "" {
		db = db.Where("username like ?", "%"+info.Username+"%")
	}
	var mobileUsers []mobile.MobileUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&mobileUsers).Error
	return mobileUsers, total, err
}
