package system

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type TeamDraftsService struct {}
// CreateTeamDrafts 创建drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService) CreateTeamDrafts(ctx context.Context, drafts *system.TeamDrafts) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(drafts).Error
	return err
}

// DeleteTeamDrafts 删除drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService)DeleteTeamDrafts(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.TeamDrafts{},"id = ?",id).Error
	return err
}

// DeleteTeamDraftsByIds 批量删除drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService)DeleteTeamDraftsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.TeamDrafts{},"id in ?",ids).Error
	return err
}

// UpdateTeamDrafts 更新drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService)UpdateTeamDrafts(ctx context.Context, drafts system.TeamDrafts) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamDrafts{}).Where("id = ?",drafts.Id).Updates(&drafts).Error
	return err
}

// GetTeamDrafts 根据id获取drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService)GetTeamDrafts(ctx context.Context, id string) (drafts system.TeamDrafts, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&drafts).Error
	return
}
// GetTeamDraftsInfoList 分页获取drafts表记录
// Author [yourname](https://github.com/yourname)
func (draftsService *TeamDraftsService)GetTeamDraftsInfoList(ctx context.Context, info systemReq.TeamDraftsSearch) (list []system.TeamDrafts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamDrafts{})
    var draftss []system.TeamDrafts
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.UserId != nil {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.DraftType != nil {
        db = db.Where("draft_type = ?", *info.DraftType)
    }
	if info.DraftMessageData != "" {
		db = db.Where("draft_message_data = ?", info.DraftMessageData)
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

	err = db.Find(&draftss).Error
	return  draftss, total, err
}
func (draftsService *TeamDraftsService)GetTeamDraftsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
