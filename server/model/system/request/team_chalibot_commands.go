package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeamChaliBotCommandsSearch struct{
      Id  *int `json:"id" form:"id"` 
      BotId  *int `json:"botId" form:"botId"` 
      Command  *string `json:"command" form:"command"` 
      Description  *string `json:"description" form:"description"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
