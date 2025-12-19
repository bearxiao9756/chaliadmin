import service from '@/utils/request'
// @Tags Auths
// @Summary 创建auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Auths true "创建auths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /auths/createAuths [post]
export const createAuths = (data) => {
  return service({
    url: '/auths/createAuths',
    method: 'post',
    data
  })
}

// @Tags Auths
// @Summary 删除auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Auths true "删除auths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /auths/deleteAuths [delete]
export const deleteAuths = (params) => {
  return service({
    url: '/auths/deleteAuths',
    method: 'delete',
    params
  })
}

// @Tags Auths
// @Summary 批量删除auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除auths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /auths/deleteAuths [delete]
export const deleteAuthsByIds = (params) => {
  return service({
    url: '/auths/deleteAuthsByIds',
    method: 'delete',
    params
  })
}

// @Tags Auths
// @Summary 更新auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Auths true "更新auths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /auths/updateAuths [put]
export const updateAuths = (data) => {
  return service({
    url: '/auths/updateAuths',
    method: 'put',
    data
  })
}

// @Tags Auths
// @Summary 用id查询auths表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Auths true "用id查询auths表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /auths/findAuths [get]
export const findAuths = (params) => {
  return service({
    url: '/auths/findAuths',
    method: 'get',
    params
  })
}

// @Tags Auths
// @Summary 分页获取auths表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取auths表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /auths/getAuthsList [get]
export const getAuthsList = (params) => {
  return service({
    url: '/auths/getAuthsList',
    method: 'get',
    params
  })
}

// @Tags Auths
// @Summary 不需要鉴权的auths表接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.AuthsSearch true "分页获取auths表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /auths/getAuthsPublic [get]
export const getAuthsPublic = () => {
  return service({
    url: '/auths/getAuthsPublic',
    method: 'get',
  })
}
