package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamGramUsersRouter struct {}

// InitTeamGramUsersRouter 初始化 users表 路由信息
func (s *TeamGramUsersRouter) InitTeamGramUsersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	TeamgramUsersRouter := Router.Group("TeamgramUsers").Use(middleware.OperationRecord())
	TeamgramUsersRouterWithoutRecord := Router.Group("TeamgramUsers")
	TeamgramUsersRouterWithoutAuth := PublicRouter.Group("TeamgramUsers")
	{
		TeamgramUsersRouter.POST("createTeamGramUsers", TeamgramUsersApi.CreateTeamGramUsers)   // 新建users表
		TeamgramUsersRouter.DELETE("deleteTeamGramUsers", TeamgramUsersApi.DeleteTeamGramUsers) // 删除users表
		TeamgramUsersRouter.DELETE("deleteTeamGramUsersByIds", TeamgramUsersApi.DeleteTeamGramUsersByIds) // 批量删除users表
		TeamgramUsersRouter.PUT("updateTeamGramUsers", TeamgramUsersApi.UpdateTeamGramUsers)    // 更新users表
	}
	{
		TeamgramUsersRouterWithoutRecord.GET("findTeamGramUsers", TeamgramUsersApi.FindTeamGramUsers)        // 根据ID获取users表
		TeamgramUsersRouterWithoutRecord.GET("getTeamGramUsersList", TeamgramUsersApi.GetTeamGramUsersList)  // 获取users表列表
	}
	{
	    TeamgramUsersRouterWithoutAuth.GET("getTeamGramUsersPublic", TeamgramUsersApi.GetTeamGramUsersPublic)  // users表开放接口
	}
}
