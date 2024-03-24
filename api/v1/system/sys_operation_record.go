package system

import (
	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	systemReq "server-fiber/model/system/request"
	"server-fiber/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {object} return response.Response{msg=string} "创建SysOperationRecord"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *fiber.Ctx) error {
	var sysOperationRecord system.SysOperationRecord
	_ = c.BodyParser(&sysOperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		return response.FailWithMessage("创建失败", c)
	} else {
		return response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {object} return response.Response{msg=string} "删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *fiber.Ctx) error {
	var operation system.SysOperationRecord
	_ = c.BodyParser(&operation)
	if err := operationRecordService.DeleteSysOperationRecord(operation.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		return response.FailWithMessage("删除失败", c)
	} else {
		return response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {object} return response.Response{msg=string} "批量删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *fiber.Ctx) error {
	var IDS request.IdsReq
	_ = c.QueryParser(&IDS)
	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Error(err))
		return response.FailWithMessage("批量删除失败", c)
	} else {
		return response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysOperationRecord true "Id"
// @Success 200 {object} return response.Response{data=map[string]interface{},msg=string} "用id查询SysOperationRecord"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *fiber.Ctx) error {
	var sysOperationRecord system.SysOperationRecord
	_ = c.QueryParser(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	if resysOperationRecord, err := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		return response.FailWithMessage("查询失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} return response.Response{data=return response.PageResult,msg=string} "分页获取SysOperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *fiber.Ctx) error {
	var pageInfo systemReq.SysOperationRecordSearch
	_ = c.QueryParser(&pageInfo)
	if pageInfo.TypePort == system.Backend {
		if list, total, err := operationRecordService.GetSysOperationRecordInfoList(&pageInfo); err != nil {
			global.LOG.Error("获取失败!", zap.Error(err))
			return response.FailWithMessage("获取失败", c)
		} else {
			return response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
		}
	} else {
		if list, total, err := operationRecordService.GetSysOperationRecordInfoFrontendList(pageInfo); err != nil {
			global.LOG.Error("获取失败!", zap.Error(err))
			return response.FailWithMessage("获取失败", c)
		} else {
			return response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
		}
	}

}
