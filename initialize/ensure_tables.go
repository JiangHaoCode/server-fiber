package initialize

import (
	"context"
	"server-fiber/service/system"

	"gorm.io/gorm"
)

// 初始化数据库

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run 初始化数据库
func init() {
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (ensureTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		// sysModel.SysApi{},
		// sysModel.SysUser{},
		// sysModel.SysBaseMenu{},
		// sysModel.SysAuthority{},
		// sysModel.JwtBlacklist{},
		// sysModel.SysDictionary{},
		// sysModel.SysAutoCodeHistory{},
		// sysModel.SysOperationRecord{},
		// sysModel.SysDictionaryDetail{},
		// sysModel.SysBaseMenuParameter{},
		// sysModel.SysBaseMenuBtn{},
		// sysModel.SysAuthorityBtn{},
		// sysModel.SysAutoCode{},

		// adapter.CasbinRule{},

		// example.ExaFile{},
		// example.ExaCustomer{},
		// example.ExaFileChunk{},
		// example.ExaFileUploadAndDownload{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}
	return ctx, nil
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		// sysModel.SysApi{},
		// sysModel.SysUser{},
		// sysModel.SysBaseMenu{},
		// sysModel.SysAuthority{},
		// sysModel.JwtBlacklist{},
		// sysModel.SysDictionary{},
		// sysModel.SysAutoCodeHistory{},
		// sysModel.SysOperationRecord{},
		// sysModel.SysDictionaryDetail{},
		// sysModel.SysBaseMenuParameter{},
		// sysModel.SysBaseMenuBtn{},
		// sysModel.SysAuthorityBtn{},
		// sysModel.SysAutoCode{},

		// adapter.CasbinRule{},

		// example.ExaFile{},
		// example.ExaCustomer{},
		// example.ExaFileChunk{},
		// example.ExaFileUploadAndDownload{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
