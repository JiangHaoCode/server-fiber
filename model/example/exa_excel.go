package example

import "server/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName" form:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList" form:"infoList"`
}
