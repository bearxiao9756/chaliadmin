import service from '@/utils/request'
// @Tags TeamChaliBots
// @Summary 创建bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBots true "创建bots表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /teamchalibots/createTeamChaliBots [post]
export const createTeamChaliBots = (data) => {
  return service({
    url: '/teamchalibots/createTeamChaliBots',
    method: 'post',
    data
  })
}

// @Tags TeamChaliBots
// @Summary 删除bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBots true "删除bots表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamchalibots/deleteTeamChaliBots [delete]
export const deleteTeamChaliBots = (params) => {
  return service({
    url: '/teamchalibots/deleteTeamChaliBots',
    method: 'delete',
    params
  })
}

// @Tags TeamChaliBots
// @Summary 批量删除bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除bots表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamchalibots/deleteTeamChaliBots [delete]
export const deleteTeamChaliBotsByIds = (params) => {
  return service({
    url: '/teamchalibots/deleteTeamChaliBotsByIds',
    method: 'delete',
    params
  })
}

// @Tags TeamChaliBots
// @Summary 更新bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBots true "更新bots表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamchalibots/updateTeamChaliBots [put]
export const updateTeamChaliBots = (data) => {
  return service({
    url: '/teamchalibots/updateTeamChaliBots',
    method: 'put',
    data
  })
}

// @Tags TeamChaliBots
// @Summary 用id查询bots表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TeamChaliBots true "用id查询bots表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamchalibots/findTeamChaliBots [get]
export const findTeamChaliBots = (params) => {
  return service({
    url: '/teamchalibots/findTeamChaliBots',
    method: 'get',
    params
  })
}

// @Tags TeamChaliBots
// @Summary 分页获取bots表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取bots表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamchalibots/getTeamChaliBotsList [get]
export const getTeamChaliBotsList = (params) => {
  return service({
    url: '/teamchalibots/getTeamChaliBotsList',
    method: 'get',
    params
  })
}

// @Tags TeamChaliBots
// @Summary 不需要鉴权的bots表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamChaliBotsSearch true "分页获取bots表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /teamchalibots/getTeamChaliBotsPublic [get]
export const getTeamChaliBotsPublic = () => {
  return service({
    url: '/teamchalibots/getTeamChaliBotsPublic',
    method: 'get',
  })
}
