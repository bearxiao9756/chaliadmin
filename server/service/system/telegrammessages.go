
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type MessagesService struct {}
// CreateMessages 创建messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService) CreateMessages(ctx context.Context, messages *system.Messages) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(messages).Error
	return err
}

// DeleteMessages 删除messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService)DeleteMessages(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.Messages{},"id = ?",id).Error
	return err
}

// DeleteMessagesByIds 批量删除messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService)DeleteMessagesByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.Messages{},"id in ?",ids).Error
	return err
}

// UpdateMessages 更新messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService)UpdateMessages(ctx context.Context, messages system.Messages) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.Messages{}).Where("id = ?",messages.Id).Updates(&messages).Error
	return err
}

// GetMessages 根据id获取messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService)GetMessages(ctx context.Context, id string) (messages system.Messages, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&messages).Error
	return
}
// GetMessagesInfoList 分页获取messages表记录
// Author [yourname](https://github.com/yourname)
func (messagesService *MessagesService)GetMessagesInfoList(ctx context.Context, info systemReq.MessagesSearch) (list []system.Messages, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.Messages{})
    var messagess []system.Messages
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.UserMessageBoxId != nil {
        db = db.Where("user_message_box_id = ?", *info.UserMessageBoxId)
    }
    if info.DialogId1 != nil {
        db = db.Where("dialog_id1 = ?", *info.DialogId1)
    }
    if info.DialogId2 != nil {
        db = db.Where("dialog_id2 = ?", *info.DialogId2)
    }
    if info.DialogMessageId != nil {
        db = db.Where("dialog_message_id = ?", *info.DialogMessageId)
    }
    if info.SenderUserId != nil {
        db = db.Where("sender_user_id = ?", *info.SenderUserId)
    }
    if info.PeerType != nil {
        db = db.Where("peer_type = ?", *info.PeerType)
    }
    if info.PeerId != nil {
        db = db.Where("peer_id = ?", *info.PeerId)
    }
    if info.RandomId != nil {
        db = db.Where("random_id = ?", *info.RandomId)
    }
    if info.MessageFilterType != nil {
        db = db.Where("message_filter_type = ?", *info.MessageFilterType)
    }
    if info.Message != nil && *info.Message != "" {
        db = db.Where("message = ?", *info.Message)
    }
    if info.Mentioned != nil {
        db = db.Where("mentioned = ?", *info.Mentioned)
    }
    if info.Pinned != nil {
        db = db.Where("pinned = ?", *info.Pinned)
    }
    if info.HasReaction != nil {
        db = db.Where("has_reaction = ?", *info.HasReaction)
    }
    if info.Reaction != nil && *info.Reaction != "" {
        db = db.Where("reaction = ?", *info.Reaction)
    }
    if info.ReactionUnread != nil {
        db = db.Where("reaction_unread = ?", *info.ReactionUnread)
    }
	if info.StartDate2 != nil && info.EndDate2 != nil {
		db = db.Where("date2 BETWEEN ? AND ? ", *info.StartDate2, *info.EndDate2)
	}
    if info.Deleted != nil {
        db = db.Where("deleted = ?", *info.Deleted)
    }
			if len(info.CreatedAtRange) == 2 {
				db = db.Where("created_at BETWEEN ? AND ? ", info.CreatedAtRange[0], info.CreatedAtRange[1])
			}
			if len(info.UpdatedAtRange) == 2 {
				db = db.Where("updated_at BETWEEN ? AND ? ", info.UpdatedAtRange[0], info.UpdatedAtRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&messagess).Error
	return  messagess, total, err
}
func (messagesService *MessagesService)GetMessagesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
