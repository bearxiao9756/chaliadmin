package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MessagesSearch struct{
      Id  *int `json:"id" form:"id"` 
      UserId  *int `json:"userId" form:"userId"` 
      UserMessageBoxId  *int `json:"userMessageBoxId" form:"userMessageBoxId"` 
      DialogId1  *int `json:"dialogId1" form:"dialogId1"` 
      DialogId2  *int `json:"dialogId2" form:"dialogId2"` 
      DialogMessageId  *int `json:"dialogMessageId" form:"dialogMessageId"` 
      SenderUserId  *int `json:"senderUserId" form:"senderUserId"` 
      PeerType  *int `json:"peerType" form:"peerType"` 
      PeerId  *int `json:"peerId" form:"peerId"` 
      RandomId  *int `json:"randomId" form:"randomId"` 
      MessageFilterType  *int `json:"messageFilterType" form:"messageFilterType"` 
      Message  *string `json:"message" form:"message"` 
      Mentioned  *bool `json:"mentioned" form:"mentioned"` 
      Pinned  *bool `json:"pinned" form:"pinned"` 
      HasReaction  *bool `json:"hasReaction" form:"hasReaction"` 
      Reaction  *string `json:"reaction" form:"reaction"` 
      ReactionUnread  *bool `json:"reactionUnread" form:"reactionUnread"` 
      StartDate2  *int  `json:"startDate2" form:"startDate2"`
EndDate2  *int  `json:"endDate2" form:"endDate2"`
      Deleted  *bool `json:"deleted" form:"deleted"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
