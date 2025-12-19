package example

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AuthsApi struct {}



// CreateAuths 创建auths表
// @Tags Auths
// @Summary 创建auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.Auths true "创建auths表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /auths/createAuths [post]
func (authsApi *AuthsApi) CreateAuths(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var auths example.Auths
	err := c.ShouldBindJSON(&auths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authsService.CreateAuths(ctx,&auths)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAuths 删除auths表
// @Tags Auths
// @Summary 删除auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.Auths true "删除auths表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /auths/deleteAuths [delete]
func (authsApi *AuthsApi) DeleteAuths(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := authsService.DeleteAuths(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAuthsByIds 批量删除auths表
// @Tags Auths
// @Summary 批量删除auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /auths/deleteAuthsByIds [delete]
func (authsApi *AuthsApi) DeleteAuthsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := authsService.DeleteAuthsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAuths 更新auths表
// @Tags Auths
// @Summary 更新auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.Auths true "更新auths表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /auths/updateAuths [put]
func (authsApi *AuthsApi) UpdateAuths(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var auths example.Auths
	err := c.ShouldBindJSON(&auths)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authsService.UpdateAuths(ctx,auths)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAuths 用id查询auths表
// @Tags Auths
// @Summary 用id查询auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询auths表"
// @Success 200 {object} response.Response{data=example.Auths,msg=string} "查询成功"
// @Router /auths/findAuths [get]
func (authsApi *AuthsApi) FindAuths(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reauths, err := authsService.GetAuths(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reauths, c)
}
// GetAuthsList 分页获取auths表列表
// @Tags Auths
// @Summary 分页获取auths表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.AuthsSearch true "分页获取auths表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /auths/getAuthsList [get]
func (authsApi *AuthsApi) GetAuthsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo exampleReq.AuthsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := authsService.GetAuthsInfoList(ctx,pageInfo)
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

// GetAuthsPublic 不需要鉴权的auths表接口
// @Tags Auths
// @Summary 不需要鉴权的auths表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /auths/getAuthsPublic [get]
func (authsApi *AuthsApi) GetAuthsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    authsService.GetAuthsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的auths表接口信息",
    }, "获取成功", c)
}
