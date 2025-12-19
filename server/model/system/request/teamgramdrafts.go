
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type TeamgramDraftsSearch struct{
      Id  *int `json:"id" form:"id"` 
      UserId  *int `json:"userId" form:"userId"` 
      PeerDialogId  *int `json:"peerDialogId" form:"peerDialogId"` 
      DraftType  *int `json:"draftType" form:"draftType"` 
      DraftMessageData  string `json:"draftMessageData" form:"draftMessageData"` 
      Date2  *int `json:"date2" form:"date2"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
