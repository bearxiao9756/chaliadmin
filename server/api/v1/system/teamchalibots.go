package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TeamChaliBotsApi struct {}



// CreateTeamChaliBots 创建bots表
// @Tags TeamChaliBots
// @Summary 创建bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBots true "创建bots表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /teamchalibots/createTeamChaliBots [post]
func (teamchalibotsApi *TeamChaliBotsApi) CreateTeamChaliBots(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var teamchalibots system.TeamChaliBots
	err := c.ShouldBindJSON(&teamchalibots)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = teamchalibotsService.CreateTeamChaliBots(ctx,&teamchalibots)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTeamChaliBots 删除bots表
// @Tags TeamChaliBots
// @Summary 删除bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBots true "删除bots表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /teamchalibots/deleteTeamChaliBots [delete]
func (teamchalibotsApi *TeamChaliBotsApi) DeleteTeamChaliBots(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := teamchalibotsService.DeleteTeamChaliBots(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTeamChaliBotsByIds 批量删除bots表
// @Tags TeamChaliBots
// @Summary 批量删除bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /teamchalibots/deleteTeamChaliBotsByIds [delete]
func (teamchalibotsApi *TeamChaliBotsApi) DeleteTeamChaliBotsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := teamchalibotsService.DeleteTeamChaliBotsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTeamChaliBots 更新bots表
// @Tags TeamChaliBots
// @Summary 更新bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamChaliBots true "更新bots表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /teamchalibots/updateTeamChaliBots [put]
func (teamchalibotsApi *TeamChaliBotsApi) UpdateTeamChaliBots(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var teamchalibots system.TeamChaliBots
	err := c.ShouldBindJSON(&teamchalibots)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = teamchalibotsService.UpdateTeamChaliBots(ctx,teamchalibots)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTeamChaliBots 用id查询bots表
// @Tags TeamChaliBots
// @Summary 用id查询bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询bots表"
// @Success 200 {object} response.Response{data=system.TeamChaliBots,msg=string} "查询成功"
// @Router /teamchalibots/findTeamChaliBots [get]
func (teamchalibotsApi *TeamChaliBotsApi) FindTeamChaliBots(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reteamchalibots, err := teamchalibotsService.GetTeamChaliBots(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reteamchalibots, c)
}
// GetTeamChaliBotsList 分页获取bots表列表
// @Tags TeamChaliBots
// @Summary 分页获取bots表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamChaliBotsSearch true "分页获取bots表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /teamchalibots/getTeamChaliBotsList [get]
func (teamchalibotsApi *TeamChaliBotsApi) GetTeamChaliBotsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.TeamChaliBotsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := teamchalibotsService.GetTeamChaliBotsInfoList(ctx,pageInfo)
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

// GetTeamChaliBotsPublic 不需要鉴权的bots表接口
// @Tags TeamChaliBots
// @Summary 不需要鉴权的bots表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /teamchalibots/getTeamChaliBotsPublic [get]
func (teamchalibotsApi *TeamChaliBotsApi) GetTeamChaliBotsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    teamchalibotsService.GetTeamChaliBotsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的bots表接口信息",
    }, "获取成功", c)
}
