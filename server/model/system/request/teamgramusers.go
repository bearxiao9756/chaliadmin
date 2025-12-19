
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type TeamGramUsersSearch struct{
      Id  *int `json:"id" form:"id"` 
      UserType  *int `json:"userType" form:"userType"` 
      AccessHash  *int `json:"accessHash" form:"accessHash"` 
      SecretKeyId  *int `json:"secretKeyId" form:"secretKeyId"` 
      FirstName  *string `json:"firstName" form:"firstName"` 
      LastName  *string `json:"lastName" form:"lastName"` 
      Username  *string `json:"username" form:"username"` 
      Phone  *string `json:"phone" form:"phone"` 
      CountryCode  *string `json:"countryCode" form:"countryCode"` 
      Verified  *bool `json:"verified" form:"verified"` 
      Support  *bool `json:"support" form:"support"` 
      Scam  *bool `json:"scam" form:"scam"` 
      Fake  *bool `json:"fake" form:"fake"` 
      Premium  *bool `json:"premium" form:"premium"` 
      PremiumExpireDate  *int `json:"premiumExpireDate" form:"premiumExpireDate"` 
      About  *string `json:"about" form:"about"` 
      State  *int `json:"state" form:"state"` 
      IsBot  *bool `json:"isBot" form:"isBot"` 
      AccountDaysTtl  *int `json:"accountDaysTtl" form:"accountDaysTtl"` 
      PhotoId  *int `json:"photoId" form:"photoId"` 
      Restricted  *bool `json:"restricted" form:"restricted"` 
      RestrictionReason  *string `json:"restrictionReason" form:"restrictionReason"` 
      ArchiveAndMuteNewNoncontactPeers  *bool `json:"archiveAndMuteNewNoncontactPeers" form:"archiveAndMuteNewNoncontactPeers"` 
      EmojiStatusDocumentId  *int `json:"emojiStatusDocumentId" form:"emojiStatusDocumentId"` 
      EmojiStatusUntil  *int `json:"emojiStatusUntil" form:"emojiStatusUntil"` 
      StoriesMaxId  *int `json:"storiesMaxId" form:"storiesMaxId"` 
      Color  *int `json:"color" form:"color"` 
      ColorBackgroundEmojiId  *int `json:"colorBackgroundEmojiId" form:"colorBackgroundEmojiId"` 
      Birthday  *string `json:"birthday" form:"birthday"` 
      PersonalChannelId  *int `json:"personalChannelId" form:"personalChannelId"` 
      AuthorizationTtlDays  *int `json:"authorizationTtlDays" form:"authorizationTtlDays"` 
      Deleted  *bool `json:"deleted" form:"deleted"` 
      DeleteReason  *string `json:"deleteReason" form:"deleteReason"` 
      CreatedAtRange  []time.Time  `json:"createdAtRange" form:"createdAtRange[]"`
      UpdatedAtRange  []time.Time  `json:"updatedAtRange" form:"updatedAtRange[]"`
    request.PageInfo
}
