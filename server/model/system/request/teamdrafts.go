
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type TeamDraftsSearch struct{
      UserId  *int `json:"userId" form:"userId"` 
      DraftType  *int `json:"draftType" form:"draftType"` 
      DraftMessageData  string `json:"draftMessageData" form:"draftMessageData"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
