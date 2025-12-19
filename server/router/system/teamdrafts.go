package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamDraftsRouter struct {}

// InitTeamDraftsRouter 初始化 drafts表 路由信息
func (s *TeamDraftsRouter) InitTeamDraftsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	draftsRouter := Router.Group("drafts").Use(middleware.OperationRecord())
	draftsRouterWithoutRecord := Router.Group("drafts")
	draftsRouterWithoutAuth := PublicRouter.Group("drafts")
	{
		draftsRouter.POST("createTeamDrafts", draftsApi.CreateTeamDrafts)   // 新建drafts表
		draftsRouter.DELETE("deleteTeamDrafts", draftsApi.DeleteTeamDrafts) // 删除drafts表
		draftsRouter.DELETE("deleteTeamDraftsByIds", draftsApi.DeleteTeamDraftsByIds) // 批量删除drafts表
		draftsRouter.PUT("updateTeamDrafts", draftsApi.UpdateTeamDrafts)    // 更新drafts表
	}
	{
		draftsRouterWithoutRecord.GET("findTeamDrafts", draftsApi.FindTeamDrafts)        // 根据ID获取drafts表
		draftsRouterWithoutRecord.GET("getTeamDraftsList", draftsApi.GetTeamDraftsList)  // 获取drafts表列表
	}
	{
	    draftsRouterWithoutAuth.GET("getTeamDraftsPublic", draftsApi.GetTeamDraftsPublic)  // drafts表开放接口
	}
}
