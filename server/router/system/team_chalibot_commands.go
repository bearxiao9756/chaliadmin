package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamChaliBotCommandsRouter struct {}

// InitTeamChaliBotCommandsRouter 初始化 业务机器人命令 路由信息
func (s *TeamChaliBotCommandsRouter) InitTeamChaliBotCommandsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	TeamChalibotCommandsRouter := Router.Group("TeamChalibotCommands").Use(middleware.OperationRecord())
	TeamChalibotCommandsRouterWithoutRecord := Router.Group("TeamChalibotCommands")
	TeamChalibotCommandsRouterWithoutAuth := PublicRouter.Group("TeamChalibotCommands")
	{
		TeamChalibotCommandsRouter.POST("createTeamChaliBotCommands", TeamChalibotCommandsApi.CreateTeamChaliBotCommands)   // 新建业务机器人命令
		TeamChalibotCommandsRouter.DELETE("deleteTeamChaliBotCommands", TeamChalibotCommandsApi.DeleteTeamChaliBotCommands) // 删除业务机器人命令
		TeamChalibotCommandsRouter.DELETE("deleteTeamChaliBotCommandsByIds", TeamChalibotCommandsApi.DeleteTeamChaliBotCommandsByIds) // 批量删除业务机器人命令
		TeamChalibotCommandsRouter.PUT("updateTeamChaliBotCommands", TeamChalibotCommandsApi.UpdateTeamChaliBotCommands)    // 更新业务机器人命令
	}
	{
		TeamChalibotCommandsRouterWithoutRecord.GET("findTeamChaliBotCommands", TeamChalibotCommandsApi.FindTeamChaliBotCommands)        // 根据ID获取业务机器人命令
		TeamChalibotCommandsRouterWithoutRecord.GET("getTeamChaliBotCommandsList", TeamChalibotCommandsApi.GetTeamChaliBotCommandsList)  // 获取业务机器人命令列表
	}
	{
	    TeamChalibotCommandsRouterWithoutAuth.GET("getTeamChaliBotCommandsPublic", TeamChalibotCommandsApi.GetTeamChaliBotCommandsPublic)  // 业务机器人命令开放接口
	}
}
