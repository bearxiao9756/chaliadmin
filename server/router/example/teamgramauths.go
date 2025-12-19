package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthsRouter struct {}

// InitAuthsRouter 初始化 auths表 路由信息
func (s *AuthsRouter) InitAuthsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	authsRouter := Router.Group("auths").Use(middleware.OperationRecord())
	authsRouterWithoutRecord := Router.Group("auths")
	authsRouterWithoutAuth := PublicRouter.Group("auths")
	{
		authsRouter.POST("createAuths", authsApi.CreateAuths)   // 新建auths表
		authsRouter.DELETE("deleteAuths", authsApi.DeleteAuths) // 删除auths表
		authsRouter.DELETE("deleteAuthsByIds", authsApi.DeleteAuthsByIds) // 批量删除auths表
		authsRouter.PUT("updateAuths", authsApi.UpdateAuths)    // 更新auths表
	}
	{
		authsRouterWithoutRecord.GET("findAuths", authsApi.FindAuths)        // 根据ID获取auths表
		authsRouterWithoutRecord.GET("getAuthsList", authsApi.GetAuthsList)  // 获取auths表列表
	}
	{
	    authsRouterWithoutAuth.GET("getAuthsPublic", authsApi.GetAuthsPublic)  // auths表开放接口
	}
}
