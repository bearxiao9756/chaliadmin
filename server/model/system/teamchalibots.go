
// 自动生成模板TeamChaliBots
package system
import (
	"time"
)

// bots表 结构体  TeamChaliBots
type TeamChaliBots struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  BotId  *int64 `json:"botId" form:"botId" gorm:"column:bot_id;"`  //机器人ID
  BotType  *int32 `json:"botType" form:"botType" gorm:"column:bot_type;"`  //机器人类型
  CreatorUserId  *int64 `json:"creatorUserId" form:"creatorUserId" gorm:"column:creator_user_id;"`  //创作者用户 ID
  Token  *string `json:"token" form:"token" gorm:"column:token;size:128;"`  //令牌
  Description  *string `json:"description" form:"description" gorm:"column:description;size:10240;"`  //描述
  BotChatHistory  *bool `json:"botChatHistory" form:"botChatHistory" gorm:"column:bot_chat_history;"`  //机器人聊天记录
  BotNochats  *bool `json:"botNochats" form:"botNochats" gorm:"column:bot_nochats;"`  //机器人无聊天
  Verified  *bool `json:"verified" form:"verified" gorm:"column:verified;"`  //是否验证
  BotInlineGeo  *bool `json:"botInlineGeo" form:"botInlineGeo" gorm:"column:bot_inline_geo;"`  //botInlineGeo字段
  BotInfoVersion  *int32 `json:"botInfoVersion" form:"botInfoVersion" gorm:"column:bot_info_version;"`  //botInfoVersion字段
  BotInlinePlaceholder  *string `json:"botInlinePlaceholder" form:"botInlinePlaceholder" gorm:"column:bot_inline_placeholder;size:128;"`  //botInlinePlaceholder字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建于
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //创建于
}


// TableName bots表 TeamChaliBots自定义表名 bots
func (TeamChaliBots) TableName() string {
    return "bots"
}





