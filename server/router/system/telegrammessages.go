package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MessagesRouter struct {}

// InitMessagesRouter 初始化 messages表 路由信息
func (s *MessagesRouter) InitMessagesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	messagesRouter := Router.Group("messages").Use(middleware.OperationRecord())
	messagesRouterWithoutRecord := Router.Group("messages")
	messagesRouterWithoutAuth := PublicRouter.Group("messages")
	{
		messagesRouter.POST("createMessages", messagesApi.CreateMessages)   // 新建messages表
		messagesRouter.DELETE("deleteMessages", messagesApi.DeleteMessages) // 删除messages表
		messagesRouter.DELETE("deleteMessagesByIds", messagesApi.DeleteMessagesByIds) // 批量删除messages表
		messagesRouter.PUT("updateMessages", messagesApi.UpdateMessages)    // 更新messages表
	}
	{
		messagesRouterWithoutRecord.GET("findMessages", messagesApi.FindMessages)        // 根据ID获取messages表
		messagesRouterWithoutRecord.GET("getMessagesList", messagesApi.GetMessagesList)  // 获取messages表列表
	}
	{
	    messagesRouterWithoutAuth.GET("getMessagesPublic", messagesApi.GetMessagesPublic)  // messages表开放接口
	}
}
