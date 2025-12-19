import service from '@/utils/request'
// @Tags TeamChaliBotCommands
// @Summary 创建业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBotCommands true "创建业务机器人命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /TeamChalibotCommands/createTeamChaliBotCommands [post]
export const createTeamChaliBotCommands = (data) => {
  return service({
    url: '/TeamChalibotCommands/createTeamChaliBotCommands',
    method: 'post',
    data
  })
}

// @Tags TeamChaliBotCommands
// @Summary 删除业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBotCommands true "删除业务机器人命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TeamChalibotCommands/deleteTeamChaliBotCommands [delete]
export const deleteTeamChaliBotCommands = (params) => {
  return service({
    url: '/TeamChalibotCommands/deleteTeamChaliBotCommands',
    method: 'delete',
    params
  })
}

// @Tags TeamChaliBotCommands
// @Summary 批量删除业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除业务机器人命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TeamChalibotCommands/deleteTeamChaliBotCommands [delete]
export const deleteTeamChaliBotCommandsByIds = (params) => {
  return service({
    url: '/TeamChalibotCommands/deleteTeamChaliBotCommandsByIds',
    method: 'delete',
    params
  })
}

// @Tags TeamChaliBotCommands
// @Summary 更新业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.TeamChaliBotCommands true "更新业务机器人命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TeamChalibotCommands/updateTeamChaliBotCommands [put]
export const updateTeamChaliBotCommands = (data) => {
  return service({
    url: '/TeamChalibotCommands/updateTeamChaliBotCommands',
    method: 'put',
    data
  })
}

// @Tags TeamChaliBotCommands
// @Summary 用id查询业务机器人命令
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.TeamChaliBotCommands true "用id查询业务机器人命令"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TeamChalibotCommands/findTeamChaliBotCommands [get]
export const findTeamChaliBotCommands = (params) => {
  return service({
    url: '/TeamChalibotCommands/findTeamChaliBotCommands',
    method: 'get',
    params
  })
}

// @Tags TeamChaliBotCommands
// @Summary 分页获取业务机器人命令列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取业务机器人命令列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TeamChalibotCommands/getTeamChaliBotCommandsList [get]
export const getTeamChaliBotCommandsList = (params) => {
  return service({
    url: '/TeamChalibotCommands/getTeamChaliBotCommandsList',
    method: 'get',
    params
  })
}

// @Tags TeamChaliBotCommands
// @Summary 不需要鉴权的业务机器人命令接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.TeamChaliBotCommandsSearch true "分页获取业务机器人命令列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /TeamChalibotCommands/getTeamChaliBotCommandsPublic [get]
export const getTeamChaliBotCommandsPublic = () => {
  return service({
    url: '/TeamChalibotCommands/getTeamChaliBotCommandsPublic',
    method: 'get',
  })
}
