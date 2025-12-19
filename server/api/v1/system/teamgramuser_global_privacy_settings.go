package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserGlobalPrivacySettingsApi struct {}



// CreateUserGlobalPrivacySettings 创建userGlobalPrivacySettings表
// @Tags UserGlobalPrivacySettings
// @Summary 创建userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserGlobalPrivacySettings true "创建userGlobalPrivacySettings表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userGlobalPrivacySettings/createUserGlobalPrivacySettings [post]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) CreateUserGlobalPrivacySettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var userGlobalPrivacySettings system.UserGlobalPrivacySettings
	err := c.ShouldBindJSON(&userGlobalPrivacySettings)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userGlobalPrivacySettingsService.CreateUserGlobalPrivacySettings(ctx,&userGlobalPrivacySettings)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUserGlobalPrivacySettings 删除userGlobalPrivacySettings表
// @Tags UserGlobalPrivacySettings
// @Summary 删除userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserGlobalPrivacySettings true "删除userGlobalPrivacySettings表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userGlobalPrivacySettings/deleteUserGlobalPrivacySettings [delete]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) DeleteUserGlobalPrivacySettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := userGlobalPrivacySettingsService.DeleteUserGlobalPrivacySettings(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserGlobalPrivacySettingsByIds 批量删除userGlobalPrivacySettings表
// @Tags UserGlobalPrivacySettings
// @Summary 批量删除userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userGlobalPrivacySettings/deleteUserGlobalPrivacySettingsByIds [delete]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) DeleteUserGlobalPrivacySettingsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := userGlobalPrivacySettingsService.DeleteUserGlobalPrivacySettingsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserGlobalPrivacySettings 更新userGlobalPrivacySettings表
// @Tags UserGlobalPrivacySettings
// @Summary 更新userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.UserGlobalPrivacySettings true "更新userGlobalPrivacySettings表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userGlobalPrivacySettings/updateUserGlobalPrivacySettings [put]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) UpdateUserGlobalPrivacySettings(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var userGlobalPrivacySettings system.UserGlobalPrivacySettings
	err := c.ShouldBindJSON(&userGlobalPrivacySettings)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userGlobalPrivacySettingsService.UpdateUserGlobalPrivacySettings(ctx,userGlobalPrivacySettings)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserGlobalPrivacySettings 用id查询userGlobalPrivacySettings表
// @Tags UserGlobalPrivacySettings
// @Summary 用id查询userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询userGlobalPrivacySettings表"
// @Success 200 {object} response.Response{data=system.UserGlobalPrivacySettings,msg=string} "查询成功"
// @Router /userGlobalPrivacySettings/findUserGlobalPrivacySettings [get]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) FindUserGlobalPrivacySettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reuserGlobalPrivacySettings, err := userGlobalPrivacySettingsService.GetUserGlobalPrivacySettings(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reuserGlobalPrivacySettings, c)
}
// GetUserGlobalPrivacySettingsList 分页获取userGlobalPrivacySettings表列表
// @Tags UserGlobalPrivacySettings
// @Summary 分页获取userGlobalPrivacySettings表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.UserGlobalPrivacySettingsSearch true "分页获取userGlobalPrivacySettings表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userGlobalPrivacySettings/getUserGlobalPrivacySettingsList [get]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) GetUserGlobalPrivacySettingsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.UserGlobalPrivacySettingsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userGlobalPrivacySettingsService.GetUserGlobalPrivacySettingsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetUserGlobalPrivacySettingsPublic 不需要鉴权的userGlobalPrivacySettings表接口
// @Tags UserGlobalPrivacySettings
// @Summary 不需要鉴权的userGlobalPrivacySettings表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGlobalPrivacySettings/getUserGlobalPrivacySettingsPublic [get]
func (userGlobalPrivacySettingsApi *UserGlobalPrivacySettingsApi) GetUserGlobalPrivacySettingsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    userGlobalPrivacySettingsService.GetUserGlobalPrivacySettingsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的userGlobalPrivacySettings表接口信息",
    }, "获取成功", c)
}
