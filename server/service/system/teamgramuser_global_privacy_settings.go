
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type UserGlobalPrivacySettingsService struct {}
// CreateUserGlobalPrivacySettings 创建userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService) CreateUserGlobalPrivacySettings(ctx context.Context, userGlobalPrivacySettings *system.UserGlobalPrivacySettings) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(userGlobalPrivacySettings).Error
	return err
}

// DeleteUserGlobalPrivacySettings 删除userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)DeleteUserGlobalPrivacySettings(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.UserGlobalPrivacySettings{},"id = ?",id).Error
	return err
}

// DeleteUserGlobalPrivacySettingsByIds 批量删除userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)DeleteUserGlobalPrivacySettingsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.UserGlobalPrivacySettings{},"id in ?",ids).Error
	return err
}

// UpdateUserGlobalPrivacySettings 更新userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)UpdateUserGlobalPrivacySettings(ctx context.Context, userGlobalPrivacySettings system.UserGlobalPrivacySettings) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.UserGlobalPrivacySettings{}).Where("id = ?",userGlobalPrivacySettings.Id).Updates(&userGlobalPrivacySettings).Error
	return err
}

// GetUserGlobalPrivacySettings 根据id获取userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)GetUserGlobalPrivacySettings(ctx context.Context, id string) (userGlobalPrivacySettings system.UserGlobalPrivacySettings, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&userGlobalPrivacySettings).Error
	return
}
// GetUserGlobalPrivacySettingsInfoList 分页获取userGlobalPrivacySettings表记录
// Author [yourname](https://github.com/yourname)
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)GetUserGlobalPrivacySettingsInfoList(ctx context.Context, info systemReq.UserGlobalPrivacySettingsSearch) (list []system.UserGlobalPrivacySettings, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.UserGlobalPrivacySettings{})
    var userGlobalPrivacySettingss []system.UserGlobalPrivacySettings
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.ArchiveAndMuteNewNoncontactPeers != nil {
        db = db.Where("archive_and_mute_new_noncontact_peers = ?", *info.ArchiveAndMuteNewNoncontactPeers)
    }
    if info.KeepArchivedUnmuted != nil {
        db = db.Where("keep_archived_unmuted = ?", *info.KeepArchivedUnmuted)
    }
    if info.KeepArchivedFolders != nil {
        db = db.Where("keep_archived_folders = ?", *info.KeepArchivedFolders)
    }
    if info.HideReadMarks != nil {
        db = db.Where("hide_read_marks = ?", *info.HideReadMarks)
    }
    if info.NewNoncontactPeersRequirePremium != nil {
        db = db.Where("new_noncontact_peers_require_premium = ?", *info.NewNoncontactPeersRequirePremium)
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

	err = db.Find(&userGlobalPrivacySettingss).Error
	return  userGlobalPrivacySettingss, total, err
}
func (userGlobalPrivacySettingsService *UserGlobalPrivacySettingsService)GetUserGlobalPrivacySettingsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
