import service from '@/utils/request'
// @Tags Messages
// @Summary 创建messages表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Messages true "创建messages表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /messages/createMessages [post]
export const createMessages = (data) => {
  return service({
    url: '/messages/createMessages',
    method: 'post',
    data
  })
}

// @Tags Messages
// @Summary 删除messages表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Messages true "删除messages表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /messages/deleteMessages [delete]
export const deleteMessages = (params) => {
  return service({
    url: '/messages/deleteMessages',
    method: 'delete',
    params
  })
}

// @Tags Messages
// @Summary 批量删除messages表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除messages表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /messages/deleteMessages [delete]
export const deleteMessagesByIds = (params) => {
  return service({
    url: '/messages/deleteMessagesByIds',
    method: 'delete',
    params
  })
}

// @Tags Messages
// @Summary 更新messages表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Messages true "更新messages表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /messages/updateMessages [put]
export const updateMessages = (data) => {
  return service({
    url: '/messages/updateMessages',
    method: 'put',
    data
  })
}

// @Tags Messages
// @Summary 用id查询messages表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Messages true "用id查询messages表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /messages/findMessages [get]
export const findMessages = (params) => {
  return service({
    url: '/messages/findMessages',
    method: 'get',
    params
  })
}

// @Tags Messages
// @Summary 分页获取messages表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取messages表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /messages/getMessagesList [get]
export const getMessagesList = (params) => {
  return service({
    url: '/messages/getMessagesList',
    method: 'get',
    params
  })
}

// @Tags Messages
// @Summary 不需要鉴权的messages表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.MessagesSearch true "分页获取messages表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /messages/getMessagesPublic [get]
export const getMessagesPublic = () => {
  return service({
    url: '/messages/getMessagesPublic',
    method: 'get',
  })
}
