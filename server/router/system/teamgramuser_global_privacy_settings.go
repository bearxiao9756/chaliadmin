package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserGlobalPrivacySettingsRouter struct {}

// InitUserGlobalPrivacySettingsRouter 初始化 userGlobalPrivacySettings表 路由信息
func (s *UserGlobalPrivacySettingsRouter) InitUserGlobalPrivacySettingsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userGlobalPrivacySettingsRouter := Router.Group("userGlobalPrivacySettings").Use(middleware.OperationRecord())
	userGlobalPrivacySettingsRouterWithoutRecord := Router.Group("userGlobalPrivacySettings")
	userGlobalPrivacySettingsRouterWithoutAuth := PublicRouter.Group("userGlobalPrivacySettings")
	{
		userGlobalPrivacySettingsRouter.POST("createUserGlobalPrivacySettings", userGlobalPrivacySettingsApi.CreateUserGlobalPrivacySettings)   // 新建userGlobalPrivacySettings表
		userGlobalPrivacySettingsRouter.DELETE("deleteUserGlobalPrivacySettings", userGlobalPrivacySettingsApi.DeleteUserGlobalPrivacySettings) // 删除userGlobalPrivacySettings表
		userGlobalPrivacySettingsRouter.DELETE("deleteUserGlobalPrivacySettingsByIds", userGlobalPrivacySettingsApi.DeleteUserGlobalPrivacySettingsByIds) // 批量删除userGlobalPrivacySettings表
		userGlobalPrivacySettingsRouter.PUT("updateUserGlobalPrivacySettings", userGlobalPrivacySettingsApi.UpdateUserGlobalPrivacySettings)    // 更新userGlobalPrivacySettings表
	}
	{
		userGlobalPrivacySettingsRouterWithoutRecord.GET("findUserGlobalPrivacySettings", userGlobalPrivacySettingsApi.FindUserGlobalPrivacySettings)        // 根据ID获取userGlobalPrivacySettings表
		userGlobalPrivacySettingsRouterWithoutRecord.GET("getUserGlobalPrivacySettingsList", userGlobalPrivacySettingsApi.GetUserGlobalPrivacySettingsList)  // 获取userGlobalPrivacySettings表列表
	}
	{
	    userGlobalPrivacySettingsRouterWithoutAuth.GET("getUserGlobalPrivacySettingsPublic", userGlobalPrivacySettingsApi.GetUserGlobalPrivacySettingsPublic)  // userGlobalPrivacySettings表开放接口
	}
}
