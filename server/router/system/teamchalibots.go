package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamChaliBotsRouter struct {}

// InitTeamChaliBotsRouter 初始化 bots表 路由信息
func (s *TeamChaliBotsRouter) InitTeamChaliBotsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	teamchalibotsRouter := Router.Group("teamchalibots").Use(middleware.OperationRecord())
	teamchalibotsRouterWithoutRecord := Router.Group("teamchalibots")
	teamchalibotsRouterWithoutAuth := PublicRouter.Group("teamchalibots")
	{
		teamchalibotsRouter.POST("createTeamChaliBots", teamchalibotsApi.CreateTeamChaliBots)   // 新建bots表
		teamchalibotsRouter.DELETE("deleteTeamChaliBots", teamchalibotsApi.DeleteTeamChaliBots) // 删除bots表
		teamchalibotsRouter.DELETE("deleteTeamChaliBotsByIds", teamchalibotsApi.DeleteTeamChaliBotsByIds) // 批量删除bots表
		teamchalibotsRouter.PUT("updateTeamChaliBots", teamchalibotsApi.UpdateTeamChaliBots)    // 更新bots表
	}
	{
		teamchalibotsRouterWithoutRecord.GET("findTeamChaliBots", teamchalibotsApi.FindTeamChaliBots)        // 根据ID获取bots表
		teamchalibotsRouterWithoutRecord.GET("getTeamChaliBotsList", teamchalibotsApi.GetTeamChaliBotsList)  // 获取bots表列表
	}
	{
	    teamchalibotsRouterWithoutAuth.GET("getTeamChaliBotsPublic", teamchalibotsApi.GetTeamChaliBotsPublic)  // bots表开放接口
	}
}
