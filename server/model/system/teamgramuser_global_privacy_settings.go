
// 自动生成模板UserGlobalPrivacySettings
package system
import (
	"time"
)

// userGlobalPrivacySettings表 结构体  UserGlobalPrivacySettings
type UserGlobalPrivacySettings struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  UserId  *int64 `json:"userId" form:"userId" gorm:"column:user_id;"`  //用户ID
  ArchiveAndMuteNewNoncontactPeers  *bool `json:"archiveAndMuteNewNoncontactPeers" form:"archiveAndMuteNewNoncontactPeers" gorm:"column:archive_and_mute_new_noncontact_peers;"`  //存档并静音新的非联系同伴
  KeepArchivedUnmuted  *bool `json:"keepArchivedUnmuted" form:"keepArchivedUnmuted" gorm:"column:keep_archived_unmuted;"`  //已存档且未静音
  KeepArchivedFolders  *bool `json:"keepArchivedFolders" form:"keepArchivedFolders" gorm:"column:keep_archived_folders;"`  //保留存档文件夹
  HideReadMarks  *bool `json:"hideReadMarks" form:"hideReadMarks" gorm:"column:hide_read_marks;"`  //隐藏已读标记
  NewNoncontactPeersRequirePremium  *bool `json:"newNoncontactPeersRequirePremium" form:"newNoncontactPeersRequirePremium" gorm:"column:new_noncontact_peers_require_premium;"`  //新非联系人需要特权才能联系
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //更新时间
}


// TableName userGlobalPrivacySettings表 UserGlobalPrivacySettings自定义表名 user_global_privacy_settings
func (UserGlobalPrivacySettings) TableName() string {
    return "user_global_privacy_settings"
}





