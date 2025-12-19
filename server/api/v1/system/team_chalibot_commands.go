package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TeamChaliBotCommandsApi struct {}



// CreateTeamChaliBotCommands 创建业务机器人命令
// @Tags TeamChaliBotCommands
// @Summary 创建业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBotCommands true "创建业务机器人命令"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /TeamChalibotCommands/createTeamChaliBotCommands [post]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) CreateTeamChaliBotCommands(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var TeamChalibotCommands system.TeamChaliBotCommands
	err := c.ShouldBindJSON(&TeamChalibotCommands)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TeamChalibotCommandsService.CreateTeamChaliBotCommands(ctx,&TeamChalibotCommands)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTeamChaliBotCommands 删除业务机器人命令
// @Tags TeamChaliBotCommands
// @Summary 删除业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBotCommands true "删除业务机器人命令"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /TeamChalibotCommands/deleteTeamChaliBotCommands [delete]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) DeleteTeamChaliBotCommands(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := TeamChalibotCommandsService.DeleteTeamChaliBotCommands(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTeamChaliBotCommandsByIds 批量删除业务机器人命令
// @Tags TeamChaliBotCommands
// @Summary 批量删除业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /TeamChalibotCommands/deleteTeamChaliBotCommandsByIds [delete]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) DeleteTeamChaliBotCommandsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := TeamChalibotCommandsService.DeleteTeamChaliBotCommandsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTeamChaliBotCommands 更新业务机器人命令
// @Tags TeamChaliBotCommands
// @Summary 更新业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBotCommands true "更新业务机器人命令"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /TeamChalibotCommands/updateTeamChaliBotCommands [put]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) UpdateTeamChaliBotCommands(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var TeamChalibotCommands system.TeamChaliBotCommands
	err := c.ShouldBindJSON(&TeamChalibotCommands)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TeamChalibotCommandsService.UpdateTeamChaliBotCommands(ctx,TeamChalibotCommands)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTeamChaliBotCommands 用id查询业务机器人命令
// @Tags TeamChaliBotCommands
// @Summary 用id查询业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询业务机器人命令"
// @Success 200 {object} response.Response{data=system.TeamChaliBotCommands,msg=string} "查询成功"
// @Router /TeamChalibotCommands/findTeamChaliBotCommands [get]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) FindTeamChaliBotCommands(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reTeamChalibotCommands, err := TeamChalibotCommandsService.GetTeamChaliBotCommands(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reTeamChalibotCommands, c)
}
// GetTeamChaliBotCommandsList 分页获取业务机器人命令列表
// @Tags TeamChaliBotCommands
// @Summary 分页获取业务机器人命令列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamChaliBotCommandsSearch true "分页获取业务机器人命令列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /TeamChalibotCommands/getTeamChaliBotCommandsList [get]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) GetTeamChaliBotCommandsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.TeamChaliBotCommandsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := TeamChalibotCommandsService.GetTeamChaliBotCommandsInfoList(ctx,pageInfo)
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

// GetTeamChaliBotCommandsPublic 不需要鉴权的业务机器人命令接口
// @Tags TeamChaliBotCommands
// @Summary 不需要鉴权的业务机器人命令接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TeamChalibotCommands/getTeamChaliBotCommandsPublic [get]
func (TeamChalibotCommandsApi *TeamChaliBotCommandsApi) GetTeamChaliBotCommandsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    TeamChalibotCommandsService.GetTeamChaliBotCommandsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的业务机器人命令接口信息",
    }, "获取成功", c)
}
