
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type UserGlobalPrivacySettingsSearch struct{
      Id  *int `json:"id" form:"id"` 
      UserId  *int `json:"userId" form:"userId"` 
      ArchiveAndMuteNewNoncontactPeers  *bool `json:"archiveAndMuteNewNoncontactPeers" form:"archiveAndMuteNewNoncontactPeers"` 
      KeepArchivedUnmuted  *bool `json:"keepArchivedUnmuted" form:"keepArchivedUnmuted"` 
      KeepArchivedFolders  *bool `json:"keepArchivedFolders" form:"keepArchivedFolders"` 
      HideReadMarks  *bool `json:"hideReadMarks" form:"hideReadMarks"` 
      NewNoncontactPeersRequirePremium  *bool `json:"newNoncontactPeersRequirePremium" form:"newNoncontactPeersRequirePremium"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
