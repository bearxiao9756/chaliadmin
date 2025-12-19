package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TeamDraftsApi struct {}



// CreateTeamDrafts 创建drafts表
// @Tags TeamDrafts
// @Summary 创建drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamDrafts true "创建drafts表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /drafts/createTeamDrafts [post]
func (draftsApi *TeamDraftsApi) CreateTeamDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var drafts system.TeamDrafts
	err := c.ShouldBindJSON(&drafts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = draftsService.CreateTeamDrafts(ctx,&drafts)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTeamDrafts 删除drafts表
// @Tags TeamDrafts
// @Summary 删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamDrafts true "删除drafts表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /drafts/deleteTeamDrafts [delete]
func (draftsApi *TeamDraftsApi) DeleteTeamDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := draftsService.DeleteTeamDrafts(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTeamDraftsByIds 批量删除drafts表
// @Tags TeamDrafts
// @Summary 批量删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /drafts/deleteTeamDraftsByIds [delete]
func (draftsApi *TeamDraftsApi) DeleteTeamDraftsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := draftsService.DeleteTeamDraftsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTeamDrafts 更新drafts表
// @Tags TeamDrafts
// @Summary 更新drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamDrafts true "更新drafts表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /drafts/updateTeamDrafts [put]
func (draftsApi *TeamDraftsApi) UpdateTeamDrafts(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var drafts system.TeamDrafts
	err := c.ShouldBindJSON(&drafts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = draftsService.UpdateTeamDrafts(ctx,drafts)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTeamDrafts 用id查询drafts表
// @Tags TeamDrafts
// @Summary 用id查询drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询drafts表"
// @Success 200 {object} response.Response{data=system.TeamDrafts,msg=string} "查询成功"
// @Router /drafts/findTeamDrafts [get]
func (draftsApi *TeamDraftsApi) FindTeamDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	redrafts, err := draftsService.GetTeamDrafts(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redrafts, c)
}
// GetTeamDraftsList 分页获取drafts表列表
// @Tags TeamDrafts
// @Summary 分页获取drafts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamDraftsSearch true "分页获取drafts表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /drafts/getTeamDraftsList [get]
func (draftsApi *TeamDraftsApi) GetTeamDraftsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.TeamDraftsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := draftsService.GetTeamDraftsInfoList(ctx,pageInfo)
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

// GetTeamDraftsPublic 不需要鉴权的drafts表接口
// @Tags TeamDrafts
// @Summary 不需要鉴权的drafts表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /drafts/getTeamDraftsPublic [get]
func (draftsApi *TeamDraftsApi) GetTeamDraftsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    draftsService.GetTeamDraftsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的drafts表接口信息",
    }, "获取成功", c)
}
