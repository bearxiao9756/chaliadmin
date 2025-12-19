import service from '@/utils/request'
// @Tags UserGlobalPrivacySettings
// @Summary 创建userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGlobalPrivacySettings true "创建userGlobalPrivacySettings表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userGlobalPrivacySettings/createUserGlobalPrivacySettings [post]
export const createUserGlobalPrivacySettings = (data) => {
  return service({
    url: '/userGlobalPrivacySettings/createUserGlobalPrivacySettings',
    method: 'post',
    data
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 删除userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGlobalPrivacySettings true "删除userGlobalPrivacySettings表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGlobalPrivacySettings/deleteUserGlobalPrivacySettings [delete]
export const deleteUserGlobalPrivacySettings = (params) => {
  return service({
    url: '/userGlobalPrivacySettings/deleteUserGlobalPrivacySettings',
    method: 'delete',
    params
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 批量删除userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userGlobalPrivacySettings表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGlobalPrivacySettings/deleteUserGlobalPrivacySettings [delete]
export const deleteUserGlobalPrivacySettingsByIds = (params) => {
  return service({
    url: '/userGlobalPrivacySettings/deleteUserGlobalPrivacySettingsByIds',
    method: 'delete',
    params
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 更新userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserGlobalPrivacySettings true "更新userGlobalPrivacySettings表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userGlobalPrivacySettings/updateUserGlobalPrivacySettings [put]
export const updateUserGlobalPrivacySettings = (data) => {
  return service({
    url: '/userGlobalPrivacySettings/updateUserGlobalPrivacySettings',
    method: 'put',
    data
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 用id查询userGlobalPrivacySettings表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserGlobalPrivacySettings true "用id查询userGlobalPrivacySettings表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userGlobalPrivacySettings/findUserGlobalPrivacySettings [get]
export const findUserGlobalPrivacySettings = (params) => {
  return service({
    url: '/userGlobalPrivacySettings/findUserGlobalPrivacySettings',
    method: 'get',
    params
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 分页获取userGlobalPrivacySettings表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userGlobalPrivacySettings表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userGlobalPrivacySettings/getUserGlobalPrivacySettingsList [get]
export const getUserGlobalPrivacySettingsList = (params) => {
  return service({
    url: '/userGlobalPrivacySettings/getUserGlobalPrivacySettingsList',
    method: 'get',
    params
  })
}

// @Tags UserGlobalPrivacySettings
// @Summary 不需要鉴权的userGlobalPrivacySettings表接口
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.UserGlobalPrivacySettingsSearch true "分页获取userGlobalPrivacySettings表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGlobalPrivacySettings/getUserGlobalPrivacySettingsPublic [get]
export const getUserGlobalPrivacySettingsPublic = () => {
  return service({
    url: '/userGlobalPrivacySettings/getUserGlobalPrivacySettingsPublic',
    method: 'get',
  })
}
