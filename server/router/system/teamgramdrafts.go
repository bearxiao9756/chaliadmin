package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamgramDraftsRouter struct {}

// InitTeamgramDraftsRouter 初始化 drafts表 路由信息
func (s *TeamgramDraftsRouter) InitTeamgramDraftsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	teamgramdraftsRouter := Router.Group("teamgramdrafts").Use(middleware.OperationRecord())
	teamgramdraftsRouterWithoutRecord := Router.Group("teamgramdrafts")
	teamgramdraftsRouterWithoutAuth := PublicRouter.Group("teamgramdrafts")
	{
		teamgramdraftsRouter.POST("createTeamgramDrafts", teamgramdraftsApi.CreateTeamgramDrafts)   // 新建drafts表
		teamgramdraftsRouter.DELETE("deleteTeamgramDrafts", teamgramdraftsApi.DeleteTeamgramDrafts) // 删除drafts表
		teamgramdraftsRouter.DELETE("deleteTeamgramDraftsByIds", teamgramdraftsApi.DeleteTeamgramDraftsByIds) // 批量删除drafts表
		teamgramdraftsRouter.PUT("updateTeamgramDrafts", teamgramdraftsApi.UpdateTeamgramDrafts)    // 更新drafts表
	}
	{
		teamgramdraftsRouterWithoutRecord.GET("findTeamgramDrafts", teamgramdraftsApi.FindTeamgramDrafts)        // 根据ID获取drafts表
		teamgramdraftsRouterWithoutRecord.GET("getTeamgramDraftsList", teamgramdraftsApi.GetTeamgramDraftsList)  // 获取drafts表列表
	}
	{
	    teamgramdraftsRouterWithoutAuth.GET("getTeamgramDraftsPublic", teamgramdraftsApi.GetTeamgramDraftsPublic)  // drafts表开放接口
	}
}
