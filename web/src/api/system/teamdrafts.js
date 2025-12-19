import service from '@/utils/request'
// @Tags TeamDrafts
// @Summary 创建drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamDrafts true "创建drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /drafts/createTeamDrafts [post]
export const createTeamDrafts = (data) => {
  return service({
    url: '/drafts/createTeamDrafts',
    method: 'post',
    data
  })
}

// @Tags TeamDrafts
// @Summary 删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamDrafts true "删除drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /drafts/deleteTeamDrafts [delete]
export const deleteTeamDrafts = (params) => {
  return service({
    url: '/drafts/deleteTeamDrafts',
    method: 'delete',
    params
  })
}

// @Tags TeamDrafts
// @Summary 批量删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /drafts/deleteTeamDrafts [delete]
export const deleteTeamDraftsByIds = (params) => {
  return service({
    url: '/drafts/deleteTeamDraftsByIds',
    method: 'delete',
    params
  })
}

// @Tags TeamDrafts
// @Summary 更新drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamDrafts true "更新drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /drafts/updateTeamDrafts [put]
export const updateTeamDrafts = (data) => {
  return service({
    url: '/drafts/updateTeamDrafts',
    method: 'put',
    data
  })
}

// @Tags TeamDrafts
// @Summary 用id查询drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TeamDrafts true "用id查询drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /drafts/findTeamDrafts [get]
export const findTeamDrafts = (params) => {
  return service({
    url: '/drafts/findTeamDrafts',
    method: 'get',
    params
  })
}

// @Tags TeamDrafts
// @Summary 分页获取drafts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取drafts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /drafts/getTeamDraftsList [get]
export const getTeamDraftsList = (params) => {
  return service({
    url: '/drafts/getTeamDraftsList',
    method: 'get',
    params
  })
}

// @Tags TeamDrafts
// @Summary 不需要鉴权的drafts表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamDraftsSearch true "分页获取drafts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /drafts/getTeamDraftsPublic [get]
export const getTeamDraftsPublic = () => {
  return service({
    url: '/drafts/getTeamDraftsPublic',
    method: 'get',
  })
}
