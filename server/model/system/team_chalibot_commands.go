
// 自动生成模板TeamChaliBotCommands
package system
import (
	"time"
)

// 业务机器人命令 结构体  TeamChaliBotCommands
type TeamChaliBotCommands struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  BotId  *int64 `json:"botId" form:"botId" gorm:"column:bot_id;"`  //机器人ID
  Command  *string `json:"command" form:"command" gorm:"column:command;size:128;"`  //命令
  Description  *string `json:"description" form:"description" gorm:"column:description;size:10240;"`  //描述
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //更新时间
}


// TableName 业务机器人命令 TeamChaliBotCommands自定义表名 bot_commands
func (TeamChaliBotCommands) TableName() string {
    return "bot_commands"
}





