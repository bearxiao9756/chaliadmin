
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type TeamChaliBotsService struct {}
// CreateTeamChaliBots 创建bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService) CreateTeamChaliBots(ctx context.Context, teamchalibots *system.TeamChaliBots) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(teamchalibots).Error
	return err
}

// DeleteTeamChaliBots 删除bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService)DeleteTeamChaliBots(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.TeamChaliBots{},"id = ?",id).Error
	return err
}

// DeleteTeamChaliBotsByIds 批量删除bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService)DeleteTeamChaliBotsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.TeamChaliBots{},"id in ?",ids).Error
	return err
}

// UpdateTeamChaliBots 更新bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService)UpdateTeamChaliBots(ctx context.Context, teamchalibots system.TeamChaliBots) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamChaliBots{}).Where("id = ?",teamchalibots.Id).Updates(&teamchalibots).Error
	return err
}

// GetTeamChaliBots 根据id获取bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService)GetTeamChaliBots(ctx context.Context, id string) (teamchalibots system.TeamChaliBots, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&teamchalibots).Error
	return
}
// GetTeamChaliBotsInfoList 分页获取bots表记录
// Author [yourname](https://github.com/yourname)
func (teamchalibotsService *TeamChaliBotsService)GetTeamChaliBotsInfoList(ctx context.Context, info systemReq.TeamChaliBotsSearch) (list []system.TeamChaliBots, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamChaliBots{})
    var teamchalibotss []system.TeamChaliBots
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.BotId != nil {
        db = db.Where("bot_id = ?", *info.BotId)
    }
    if info.BotType != nil {
        db = db.Where("bot_type = ?", *info.BotType)
    }
    if info.CreatorUserId != nil {
        db = db.Where("creator_user_id = ?", *info.CreatorUserId)
    }
    if info.BotChatHistory != nil {
        db = db.Where("bot_chat_history = ?", *info.BotChatHistory)
    }
    if info.BotNochats != nil {
        db = db.Where("bot_nochats = ?", *info.BotNochats)
    }
    if info.Verified != nil {
        db = db.Where("verified = ?", *info.Verified)
    }
    if info.BotInlineGeo != nil {
        db = db.Where("bot_inline_geo = ?", *info.BotInlineGeo)
    }
    if info.BotInfoVersion != nil {
        db = db.Where("bot_info_version = ?", *info.BotInfoVersion)
    }
    if info.CreatedAt != nil {
        db = db.Where("created_at = ?", *info.CreatedAt)
    }
    if info.UpdatedAt != nil {
        db = db.Where("updated_at = ?", *info.UpdatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&teamchalibotss).Error
	return  teamchalibotss, total, err
}
func (teamchalibotsService *TeamChaliBotsService)GetTeamChaliBotsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
