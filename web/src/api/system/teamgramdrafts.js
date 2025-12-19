import service from '@/utils/request'
// @Tags TeamgramDrafts
// @Summary 创建drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamgramDrafts true "创建drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /teamgramdrafts/createTeamgramDrafts [post]
export const createTeamgramDrafts = (data) => {
  return service({
    url: '/teamgramdrafts/createTeamgramDrafts',
    method: 'post',
    data
  })
}

// @Tags TeamgramDrafts
// @Summary 删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamgramDrafts true "删除drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamgramdrafts/deleteTeamgramDrafts [delete]
export const deleteTeamgramDrafts = (params) => {
  return service({
    url: '/teamgramdrafts/deleteTeamgramDrafts',
    method: 'delete',
    params
  })
}

// @Tags TeamgramDrafts
// @Summary 批量删除drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamgramdrafts/deleteTeamgramDrafts [delete]
export const deleteTeamgramDraftsByIds = (params) => {
  return service({
    url: '/teamgramdrafts/deleteTeamgramDraftsByIds',
    method: 'delete',
    params
  })
}

// @Tags TeamgramDrafts
// @Summary 更新drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamgramDrafts true "更新drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamgramdrafts/updateTeamgramDrafts [put]
export const updateTeamgramDrafts = (data) => {
  return service({
    url: '/teamgramdrafts/updateTeamgramDrafts',
    method: 'put',
    data
  })
}

// @Tags TeamgramDrafts
// @Summary 用id查询drafts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TeamgramDrafts true "用id查询drafts表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamgramdrafts/findTeamgramDrafts [get]
export const findTeamgramDrafts = (params) => {
  return service({
    url: '/teamgramdrafts/findTeamgramDrafts',
    method: 'get',
    params
  })
}

// @Tags TeamgramDrafts
// @Summary 分页获取drafts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取drafts表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamgramdrafts/getTeamgramDraftsList [get]
export const getTeamgramDraftsList = (params) => {
  return service({
    url: '/teamgramdrafts/getTeamgramDraftsList',
    method: 'get',
    params
  })
}

// @Tags TeamgramDrafts
// @Summary 不需要鉴权的drafts表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamgramDraftsSearch true "分页获取drafts表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /teamgramdrafts/getTeamgramDraftsPublic [get]
export const getTeamgramDraftsPublic = () => {
  return service({
    url: '/teamgramdrafts/getTeamgramDraftsPublic',
    method: 'get',
  })
}
