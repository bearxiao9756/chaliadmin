import service from '@/utils/request'
// @Tags TeamGramUsers
// @Summary 创建users表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamGramUsers true "创建users表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TeamgramUsers/createTeamGramUsers [post]
export const createTeamGramUsers = (data) => {
  return service({
    url: '/TeamgramUsers/createTeamGramUsers',
    method: 'post',
    data
  })
}

// @Tags TeamGramUsers
// @Summary 删除users表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamGramUsers true "删除users表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TeamgramUsers/deleteTeamGramUsers [delete]
export const deleteTeamGramUsers = (params) => {
  return service({
    url: '/TeamgramUsers/deleteTeamGramUsers',
    method: 'delete',
    params
  })
}

// @Tags TeamGramUsers
// @Summary 批量删除users表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除users表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TeamgramUsers/deleteTeamGramUsers [delete]
export const deleteTeamGramUsersByIds = (params) => {
  return service({
    url: '/TeamgramUsers/deleteTeamGramUsersByIds',
    method: 'delete',
    params
  })
}

// @Tags TeamGramUsers
// @Summary 更新users表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamGramUsers true "更新users表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TeamgramUsers/updateTeamGramUsers [put]
export const updateTeamGramUsers = (data) => {
  return service({
    url: '/TeamgramUsers/updateTeamGramUsers',
    method: 'put',
    data
  })
}

// @Tags TeamGramUsers
// @Summary 用id查询users表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TeamGramUsers true "用id查询users表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TeamgramUsers/findTeamGramUsers [get]
export const findTeamGramUsers = (params) => {
  return service({
    url: '/TeamgramUsers/findTeamGramUsers',
    method: 'get',
    params
  })
}

// @Tags TeamGramUsers
// @Summary 分页获取users表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取users表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TeamgramUsers/getTeamGramUsersList [get]
export const getTeamGramUsersList = (params) => {
  return service({
    url: '/TeamgramUsers/getTeamGramUsersList',
    method: 'get',
    params
  })
}

// @Tags TeamGramUsers
// @Summary 不需要鉴权的users表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamGramUsersSearch true "分页获取users表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TeamgramUsers/getTeamGramUsersPublic [get]
export const getTeamGramUsersPublic = () => {
  return service({
    url: '/TeamgramUsers/getTeamGramUsersPublic',
    method: 'get',
  })
}
