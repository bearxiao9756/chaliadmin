
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TeamChaliBotsSearch struct{
      Id  *int `json:"id" form:"id"` 
      BotId  *int `json:"botId" form:"botId"` 
      BotType  *int `json:"botType" form:"botType"` 
      CreatorUserId  *int `json:"creatorUserId" form:"creatorUserId"` 
      BotChatHistory  *bool `json:"botChatHistory" form:"botChatHistory"` 
      BotNochats  *bool `json:"botNochats" form:"botNochats"` 
      Verified  *bool `json:"verified" form:"verified"` 
      BotInlineGeo  *bool `json:"botInlineGeo" form:"botInlineGeo"` 
      BotInfoVersion  *int `json:"botInfoVersion" form:"botInfoVersion"` 
      CreatedAt  *time.Time `json:"createdAt" form:"createdAt"` 
      UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt"` 
    request.PageInfo
}
