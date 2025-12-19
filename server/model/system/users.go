
// 自动生成模板Users
package system
import (
	"time"
)

// users表 结构体  Users
type Users struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //用户ID
  UserType  *int32 `json:"userType" form:"userType" gorm:"column:user_type;"`  //userType字段
  AccessHash  *int64 `json:"accessHash" form:"accessHash" gorm:"column:access_hash;"`  //accessHash字段
  SecretKeyId  *int64 `json:"secretKeyId" form:"secretKeyId" gorm:"column:secret_key_id;"`  //secretKey用户ID
  FirstName  *string `json:"firstName" form:"firstName" gorm:"column:first_name;size:64;"`  //firstName字段
  LastName  *string `json:"lastName" form:"lastName" gorm:"column:last_name;size:64;"`  //lastName字段
  Username  *string `json:"username" form:"username" gorm:"column:username;size:64;"`  //username字段
  Phone  *string `json:"phone" form:"phone" gorm:"column:phone;size:255;"`  //phone字段
  CountryCode  *string `json:"countryCode" form:"countryCode" gorm:"column:country_code;size:3;"`  //countryCode字段
  Verified  *bool `json:"verified" form:"verified" gorm:"column:verified;"`  //verified字段
  Support  *bool `json:"support" form:"support" gorm:"column:support;"`  //support字段
  Scam  *bool `json:"scam" form:"scam" gorm:"column:scam;"`  //scam字段
  Fake  *bool `json:"fake" form:"fake" gorm:"column:fake;"`  //fake字段
  Premium  *bool `json:"premium" form:"premium" gorm:"column:premium;"`  //premium字段
  PremiumExpireDate  *int64 `json:"premiumExpireDate" form:"premiumExpireDate" gorm:"column:premium_expire_date;"`  //premiumExpireDate字段
  About  *string `json:"about" form:"about" gorm:"column:about;size:128;"`  //about字段
  State  *int32 `json:"state" form:"state" gorm:"column:state;"`  //state字段
  IsBot  *bool `json:"isBot" form:"isBot" gorm:"column:is_bot;"`  //isBot字段
  AccountDaysTtl  *int32 `json:"accountDaysTtl" form:"accountDaysTtl" gorm:"column:account_days_ttl;"`  //accountDaysTtl字段
  PhotoId  *int64 `json:"photoId" form:"photoId" gorm:"column:photo_id;"`  //photo用户ID
  Restricted  *bool `json:"restricted" form:"restricted" gorm:"column:restricted;"`  //restricted字段
  RestrictionReason  *string `json:"restrictionReason" form:"restrictionReason" gorm:"column:restriction_reason;size:128;"`  //restrictionReason字段
  ArchiveAndMuteNewNoncontactPeers  *bool `json:"archiveAndMuteNewNoncontactPeers" form:"archiveAndMuteNewNoncontactPeers" gorm:"column:archive_and_mute_new_noncontact_peers;"`  //archiveAndMuteNewNoncontactPeers字段
  EmojiStatusDocumentId  *int64 `json:"emojiStatusDocumentId" form:"emojiStatusDocumentId" gorm:"column:emoji_status_document_id;"`  //emojiStatusDocument用户ID
  EmojiStatusUntil  *int32 `json:"emojiStatusUntil" form:"emojiStatusUntil" gorm:"column:emoji_status_until;"`  //emojiStatusUntil字段
  StoriesMaxId  *int32 `json:"storiesMaxId" form:"storiesMaxId" gorm:"column:stories_max_id;"`  //storiesMax用户ID
  Color  *int32 `json:"color" form:"color" gorm:"column:color;"`  //color字段
  ColorBackgroundEmojiId  *int64 `json:"colorBackgroundEmojiId" form:"colorBackgroundEmojiId" gorm:"column:color_background_emoji_id;"`  //colorBackgroundEmoji用户ID
  ProfileColor  *int32 `json:"profileColor" form:"profileColor" gorm:"column:profile_color;"`  //profileColor字段
  ProfileColorBackgroundEmojiId  *int64 `json:"profileColorBackgroundEmojiId" form:"profileColorBackgroundEmojiId" gorm:"column:profile_color_background_emoji_id;"`  //profileColorBackgroundEmoji用户ID
  Birthday  *string `json:"birthday" form:"birthday" gorm:"column:birthday;"`  //birthday字段
  PersonalChannelId  *int64 `json:"personalChannelId" form:"personalChannelId" gorm:"column:personal_channel_id;"`  //personalChannel用户ID
  AuthorizationTtlDays  *int32 `json:"authorizationTtlDays" form:"authorizationTtlDays" gorm:"column:authorization_ttl_days;"`  //authorizationTtlDays字段
  Deleted  *bool `json:"deleted" form:"deleted" gorm:"column:deleted;"`  //deleted字段
  DeleteReason  *string `json:"deleteReason" form:"deleteReason" gorm:"column:delete_reason;size:128;"`  //deleteReason字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
}


// TableName users表 Users自定义表名 users
func (Users) TableName() string {
    return "users"
}





