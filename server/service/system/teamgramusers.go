
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type TeamGramUsersService struct {}
// CreateTeamGramUsers 创建users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService) CreateTeamGramUsers(ctx context.Context, TeamgramUsers *system.TeamGramUsers) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(TeamgramUsers).Error
	return err
}

// DeleteTeamGramUsers 删除users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService)DeleteTeamGramUsers(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.TeamGramUsers{},"id = ?",id).Error
	return err
}

// DeleteTeamGramUsersByIds 批量删除users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService)DeleteTeamGramUsersByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.TeamGramUsers{},"id in ?",ids).Error
	return err
}

// UpdateTeamGramUsers 更新users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService)UpdateTeamGramUsers(ctx context.Context, TeamgramUsers system.TeamGramUsers) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamGramUsers{}).Where("id = ?",TeamgramUsers.Id).Updates(&TeamgramUsers).Error
	return err
}

// GetTeamGramUsers 根据id获取users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService)GetTeamGramUsers(ctx context.Context, id string) (TeamgramUsers system.TeamGramUsers, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&TeamgramUsers).Error
	return
}
// GetTeamGramUsersInfoList 分页获取users表记录
// Author [yourname](https://github.com/yourname)
func (TeamgramUsersService *TeamGramUsersService)GetTeamGramUsersInfoList(ctx context.Context, info systemReq.TeamGramUsersSearch) (list []system.TeamGramUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamGramUsers{})
    var TeamgramUserss []system.TeamGramUsers
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.UserType != nil {
        db = db.Where("user_type = ?", *info.UserType)
    }
    if info.AccessHash != nil {
        db = db.Where("access_hash = ?", *info.AccessHash)
    }
    if info.SecretKeyId != nil {
        db = db.Where("secret_key_id = ?", *info.SecretKeyId)
    }
    if info.FirstName != nil && *info.FirstName != "" {
        db = db.Where("first_name LIKE ?", "%"+ *info.FirstName+"%")
    }
    if info.LastName != nil && *info.LastName != "" {
        db = db.Where("last_name LIKE ?", "%"+ *info.LastName+"%")
    }
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username LIKE ?", "%"+ *info.Username+"%")
    }
    if info.Phone != nil && *info.Phone != "" {
        db = db.Where("phone LIKE ?", "%"+ *info.Phone+"%")
    }
    if info.CountryCode != nil && *info.CountryCode != "" {
        db = db.Where("country_code = ?", *info.CountryCode)
    }
    if info.Verified != nil {
        db = db.Where("verified = ?", *info.Verified)
    }
    if info.Support != nil {
        db = db.Where("support = ?", *info.Support)
    }
    if info.Scam != nil {
        db = db.Where("scam = ?", *info.Scam)
    }
    if info.Fake != nil {
        db = db.Where("fake = ?", *info.Fake)
    }
    if info.Premium != nil {
        db = db.Where("premium = ?", *info.Premium)
    }
    if info.PremiumExpireDate != nil {
        db = db.Where("premium_expire_date <> ?", *info.PremiumExpireDate)
    }
    if info.About != nil && *info.About != "" {
        db = db.Where("about LIKE ?", "%"+ *info.About+"%")
    }
    if info.State != nil {
        db = db.Where("state = ?", *info.State)
    }
    if info.IsBot != nil {
        db = db.Where("is_bot = ?", *info.IsBot)
    }
    if info.AccountDaysTtl != nil {
        db = db.Where("account_days_ttl = ?", *info.AccountDaysTtl)
    }
    if info.PhotoId != nil {
        db = db.Where("photo_id = ?", *info.PhotoId)
    }
    if info.Restricted != nil {
        db = db.Where("restricted = ?", *info.Restricted)
    }
    if info.RestrictionReason != nil && *info.RestrictionReason != "" {
        db = db.Where("restriction_reason LIKE ?", "%"+ *info.RestrictionReason+"%")
    }
    if info.ArchiveAndMuteNewNoncontactPeers != nil {
        db = db.Where("archive_and_mute_new_noncontact_peers = ?", *info.ArchiveAndMuteNewNoncontactPeers)
    }
    if info.EmojiStatusDocumentId != nil {
        db = db.Where("emoji_status_document_id = ?", *info.EmojiStatusDocumentId)
    }
    if info.EmojiStatusUntil != nil {
        db = db.Where("emoji_status_until = ?", *info.EmojiStatusUntil)
    }
    if info.StoriesMaxId != nil {
        db = db.Where("stories_max_id = ?", *info.StoriesMaxId)
    }
    if info.Color != nil {
        db = db.Where("color = ?", *info.Color)
    }
    if info.ColorBackgroundEmojiId != nil {
        db = db.Where("color_background_emoji_id = ?", *info.ColorBackgroundEmojiId)
    }
    if info.Birthday != nil && *info.Birthday != "" {
        db = db.Where("birthday LIKE ?", "%"+ *info.Birthday+"%")
    }
    if info.PersonalChannelId != nil {
        db = db.Where("personal_channel_id = ?", *info.PersonalChannelId)
    }
    if info.AuthorizationTtlDays != nil {
        db = db.Where("authorization_ttl_days = ?", *info.AuthorizationTtlDays)
    }
    if info.Deleted != nil {
        db = db.Where("deleted = ?", *info.Deleted)
    }
    if info.DeleteReason != nil && *info.DeleteReason != "" {
        db = db.Where("delete_reason LIKE ?", "%"+ *info.DeleteReason+"%")
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

	err = db.Find(&TeamgramUserss).Error
	return  TeamgramUserss, total, err
}
func (TeamgramUsersService *TeamGramUsersService)GetTeamGramUsersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
