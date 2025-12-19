
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type TeamgramDraftsService struct {}
// CreateTeamgramDrafts 创建drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService) CreateTeamgramDrafts(ctx context.Context, teamgramdrafts *system.TeamgramDrafts) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(teamgramdrafts).Error
	return err
}

// DeleteTeamgramDrafts 删除drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService)DeleteTeamgramDrafts(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.TeamgramDrafts{},"id = ?",id).Error
	return err
}

// DeleteTeamgramDraftsByIds 批量删除drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService)DeleteTeamgramDraftsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.TeamgramDrafts{},"id in ?",ids).Error
	return err
}

// UpdateTeamgramDrafts 更新drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService)UpdateTeamgramDrafts(ctx context.Context, teamgramdrafts system.TeamgramDrafts) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamgramDrafts{}).Where("id = ?",teamgramdrafts.Id).Updates(&teamgramdrafts).Error
	return err
}

// GetTeamgramDrafts 根据id获取drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService)GetTeamgramDrafts(ctx context.Context, id string) (teamgramdrafts system.TeamgramDrafts, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&teamgramdrafts).Error
	return
}
// GetTeamgramDraftsInfoList 分页获取drafts表记录
// Author [yourname](https://github.com/yourname)
func (teamgramdraftsService *TeamgramDraftsService)GetTeamgramDraftsInfoList(ctx context.Context, info systemReq.TeamgramDraftsSearch) (list []system.TeamgramDrafts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamgramDrafts{})
    var teamgramdraftss []system.TeamgramDrafts
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.PeerDialogId != nil {
        db = db.Where("peer_dialog_id = ?", *info.PeerDialogId)
    }
    if info.DraftType != nil {
        db = db.Where("draft_type = ?", *info.DraftType)
    }
    if info.DraftMessageData != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
    if info.Date2 != nil {
        db = db.Where("date2 = ?", *info.Date2)
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

	err = db.Find(&teamgramdraftss).Error
	return  teamgramdraftss, total, err
}
func (teamgramdraftsService *TeamgramDraftsService)GetTeamgramDraftsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
