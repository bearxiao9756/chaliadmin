
// 自动生成模板TeamDrafts
package system
import (
	"time"
	"gorm.io/datatypes"
)

// drafts表 结构体  TeamDrafts
type TeamDrafts struct {
  Id  *int32 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  UserId  *int32 `json:"userId" form:"userId" gorm:"column:user_id;"`  //用户ID
  PeerDialogId  *int64 `json:"peerDialogId" form:"peerDialogId" gorm:"column:peer_dialog_id;"`  //会话ID
  DraftType  *int32 `json:"draftType" form:"draftType" gorm:"column:draft_type;"`  //草稿类型
  DraftMessageData  datatypes.JSON `json:"draftMessageData" form:"draftMessageData" gorm:"column:draft_message_data;" swaggertype:"array,object"`  //草稿消息数据
  Date2  *int64 `json:"date2" form:"date2" gorm:"column:date2;"`  //date2字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //更新时间
}


// TableName drafts表 TeamDrafts自定义表名 drafts
func (TeamDrafts) TableName() string {
    return "drafts"
}





