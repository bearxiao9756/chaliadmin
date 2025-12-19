package system

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TeamgramDraftsApi struct {}



// CreateTeamgramDrafts 创建drafts表
// @Tags TeamgramDrafts
// @Summary 创建drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamgramDrafts true "创建drafts表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /teamgramdrafts/createTeamgramDrafts [post]
func (teamgramdraftsApi *TeamgramDraftsApi) CreateTeamgramDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var teamgramdrafts system.TeamgramDrafts
	err := c.ShouldBindJSON(&teamgramdrafts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = teamgramdraftsService.CreateTeamgramDrafts(ctx,&teamgramdrafts)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTeamgramDrafts 删除drafts表
// @Tags TeamgramDrafts
// @Summary 删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamgramDrafts true "删除drafts表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /teamgramdrafts/deleteTeamgramDrafts [delete]
func (teamgramdraftsApi *TeamgramDraftsApi) DeleteTeamgramDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := teamgramdraftsService.DeleteTeamgramDrafts(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTeamgramDraftsByIds 批量删除drafts表
// @Tags TeamgramDrafts
// @Summary 批量删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /teamgramdrafts/deleteTeamgramDraftsByIds [delete]
func (teamgramdraftsApi *TeamgramDraftsApi) DeleteTeamgramDraftsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := teamgramdraftsService.DeleteTeamgramDraftsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTeamgramDrafts 更新drafts表
// @Tags TeamgramDrafts
// @Summary 更新drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.TeamgramDrafts true "更新drafts表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /teamgramdrafts/updateTeamgramDrafts [put]
func (teamgramdraftsApi *TeamgramDraftsApi) UpdateTeamgramDrafts(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var teamgramdrafts system.TeamgramDrafts
	err := c.ShouldBindJSON(&teamgramdrafts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = teamgramdraftsService.UpdateTeamgramDrafts(ctx,teamgramdrafts)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTeamgramDrafts 用id查询drafts表
// @Tags TeamgramDrafts
// @Summary 用id查询drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询drafts表"
// @Success 200 {object} response.Response{data=system.TeamgramDrafts,msg=string} "查询成功"
// @Router /teamgramdrafts/findTeamgramDrafts [get]
func (teamgramdraftsApi *TeamgramDraftsApi) FindTeamgramDrafts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reteamgramdrafts, err := teamgramdraftsService.GetTeamgramDrafts(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reteamgramdrafts, c)
}
// GetTeamgramDraftsList 分页获取drafts表列表
// @Tags TeamgramDrafts
// @Summary 分页获取drafts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamgramDraftsSearch true "分页获取drafts表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /teamgramdrafts/getTeamgramDraftsList [get]
func (teamgramdraftsApi *TeamgramDraftsApi) GetTeamgramDraftsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo systemReq.TeamgramDraftsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := teamgramdraftsService.GetTeamgramDraftsInfoList(ctx,pageInfo)
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

// GetTeamgramDraftsPublic 不需要鉴权的drafts表接口
// @Tags TeamgramDrafts
// @Summary 不需要鉴权的drafts表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /teamgramdrafts/getTeamgramDraftsPublic [get]
func (teamgramdraftsApi *TeamgramDraftsApi) GetTeamgramDraftsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    teamgramdraftsService.GetTeamgramDraftsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的drafts表接口信息",
    }, "获取成功", c)
}
