
// 自动生成模板Messages
package system
import (
	"time"
)

// messages表 结构体  Messages
type Messages struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  UserId  *int64 `json:"userId" form:"userId" gorm:"column:user_id;"`  //用户ID
  UserMessageBoxId  *int32 `json:"userMessageBoxId" form:"userMessageBoxId" gorm:"column:user_message_box_id;"`  //消盒子序号
  DialogId1  *int64 `json:"dialogId1" form:"dialogId1" gorm:"column:dialog_id1;"`  //发送者会话编号
  DialogId2  *int64 `json:"dialogId2" form:"dialogId2" gorm:"column:dialog_id2;"`  //接收者会话编号
  DialogMessageId  *int64 `json:"dialogMessageId" form:"dialogMessageId" gorm:"column:dialog_message_id;"`  //消息ID
  SenderUserId  *int64 `json:"senderUserId" form:"senderUserId" gorm:"column:sender_user_id;"`  //发送者ID
  PeerType  *int32 `json:"peerType" form:"peerType" gorm:"column:peer_type;"`  //消息类型
  PeerId  *int64 `json:"peerId" form:"peerId" gorm:"column:peer_id;"`  //消息类型ID
  RandomId  *int64 `json:"randomId" form:"randomId" gorm:"column:random_id;"`  //幂等随机数
  MessageFilterType  *int32 `json:"messageFilterType" form:"messageFilterType" gorm:"column:message_filter_type;"`  //消息过滤类型
  Message  *string `json:"message" form:"message" gorm:"column:message;size:6000;"`  //消息内容
  Mentioned  *bool `json:"mentioned" form:"mentioned" gorm:"column:mentioned;"`  //是否提及
  MediaUnread  *bool `json:"mediaUnread" form:"mediaUnread" gorm:"column:media_unread;"`  //未读媒体
  Pinned  *bool `json:"pinned" form:"pinned" gorm:"column:pinned;"`  //是否置顶
  HasReaction  *bool `json:"hasReaction" form:"hasReaction" gorm:"column:has_reaction;"`  //是否反应
  Reaction  *string `json:"reaction" form:"reaction" gorm:"column:reaction;size:16;"`  //反应内容
  ReactionDate  *int64 `json:"reactionDate" form:"reactionDate" gorm:"column:reaction_date;"`  //反应时间
  ReactionUnread  *bool `json:"reactionUnread" form:"reactionUnread" gorm:"column:reaction_unread;"`  //未读反应
  Date2  *int64 `json:"date2" form:"date2" gorm:"column:date2;"`  //时间
  TtlPeriod  *int32 `json:"ttlPeriod" form:"ttlPeriod" gorm:"column:ttl_period;"`  //过期时间
  SavedPeerType  *int32 `json:"savedPeerType" form:"savedPeerType" gorm:"column:saved_peer_type;"`  //保存类型
  SavedPeerId  *int64 `json:"savedPeerId" form:"savedPeerId" gorm:"column:saved_peer_id;"`  //保存类型ID
  Deleted  *bool `json:"deleted" form:"deleted" gorm:"column:deleted;"`  //是否删除
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //更新时间
}


// TableName messages表 Messages自定义表名 messages
func (Messages) TableName() string {
    return "messages"
}





