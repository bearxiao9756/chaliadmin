// 自动生成模板Users
package system

import (
	"time"
)

// users表 结构体  Users
type Users struct {
	Id                               *int64     `json:"id" form:"id" gorm:"primarykey;column:id;"`                                                                                     //用户ID
	UserType                         *int32     `json:"userType" form:"userType" gorm:"column:user_type;"`                                                                             //用户类型
	AccessHash                       *int64     `json:"accessHash" form:"accessHash" gorm:"column:access_hash;"`                                                                       //accessHash
	SecretKeyId                      *int64     `json:"secretKeyId" form:"secretKeyId" gorm:"column:secret_key_id;"`                                                                   //secretKey
	FirstName                        *string    `json:"firstName" form:"firstName" gorm:"column:first_name;size:64;"`                                                                  //姓名-姓名
	LastName                         *string    `json:"lastName" form:"lastName" gorm:"column:last_name;size:64;"`                                                                     //姓名-姓
	Username                         *string    `json:"username" form:"username" gorm:"column:username;size:64;"`                                                                      //姓名-氏
	Phone                            *string    `json:"phone" form:"phone" gorm:"column:phone;size:255;"`                                                                              //手机
	CountryCode                      *string    `json:"countryCode" form:"countryCode" gorm:"column:country_code;size:3;"`                                                             //国家编码
	Verified                         *bool      `json:"verified" form:"verified" gorm:"column:verified;"`                                                                              //是否验证
	Support                          *bool      `json:"support" form:"support" gorm:"column:support;"`                                                                                 //是否支持
	Scam                             *bool      `json:"scam" form:"scam" gorm:"column:scam;"`                                                                                          //是否欺诈
	Fake                             *bool      `json:"fake" form:"fake" gorm:"column:fake;"`                                                                                          //是否假冒
	Premium                          *bool      `json:"premium" form:"premium" gorm:"column:premium;"`                                                                                 //是否特权
	PremiumExpireDate                *int64     `json:"premiumExpireDate" form:"premiumExpireDate" gorm:"column:premium_expire_date;"`                                                 //特权时间
	About                            *string    `json:"about" form:"about" gorm:"column:about;size:128;"`                                                                              //个人简介
	State                            *int32     `json:"state" form:"state" gorm:"column:state;"`                                                                                       //用户状态
	IsBot                            *bool      `json:"isBot" form:"isBot" gorm:"column:is_bot;"`                                                                                      //是否机器人
	AccountDaysTtl                   *int32     `json:"accountDaysTtl" form:"accountDaysTtl" gorm:"column:account_days_ttl;"`                                                          //自动删除时间
	PhotoId                          *int64     `json:"photoId" form:"photoId" gorm:"column:photo_id;"`                                                                                //头像ID
	Restricted                       *bool      `json:"restricted" form:"restricted" gorm:"column:restricted;"`                                                                        //是否封禁
	RestrictionReason                *string    `json:"restrictionReason" form:"restrictionReason" gorm:"column:restriction_reason;size:128;"`                                         //封禁原因
	ArchiveAndMuteNewNoncontactPeers *bool      `json:"archiveAndMuteNewNoncontactPeers" form:"archiveAndMuteNewNoncontactPeers" gorm:"column:archive_and_mute_new_noncontact_peers;"` //是否陌生人静音
	EmojiStatusDocumentId            *int64     `json:"emojiStatusDocumentId" form:"emojiStatusDocumentId" gorm:"column:emoji_status_document_id;"`                                    //表情文档ID
	EmojiStatusUntil                 *int32     `json:"emojiStatusUntil" form:"emojiStatusUntil" gorm:"column:emoji_status_until;"`                                                    //表情到期时间
	StoriesMaxId                     *int32     `json:"storiesMaxId" form:"storiesMaxId" gorm:"column:stories_max_id;"`                                                                //最新的故事 ID
	Color                            *int32     `json:"color" form:"color" gorm:"column:color;"`                                                                                       //主颜色
	ColorBackgroundEmojiId           *int64     `json:"colorBackgroundEmojiId" form:"colorBackgroundEmojiId" gorm:"column:color_background_emoji_id;"`                                 //表情背景色ID
	ProfileColor                     *int32     `json:"profileColor" form:"profileColor" gorm:"column:profile_color;"`                                                                 //个人页面主色
	ProfileColorBackgroundEmojiId    *int64     `json:"profileColorBackgroundEmojiId" form:"profileColorBackgroundEmojiId" gorm:"column:profile_color_background_emoji_id;"`           //个人页面表情背景色ID
	Birthday                         *string    `json:"birthday" form:"birthday" gorm:"column:birthday;"`                                                                              //生日
	PersonalChannelId                *int64     `json:"personalChannelId" form:"personalChannelId" gorm:"column:personal_channel_id;"`                                                 //个人频道ID
	AuthorizationTtlDays             *int32     `json:"authorizationTtlDays" form:"authorizationTtlDays" gorm:"column:authorization_ttl_days;"`                                        //Token有效期
	Deleted                          *bool      `json:"deleted" form:"deleted" gorm:"column:deleted;"`                                                                                 //是否删除
	DeleteReason                     *string    `json:"deleteReason" form:"deleteReason" gorm:"column:delete_reason;size:128;"`                                                        //删除原因
	CreatedAt                        *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                                                                          //创建时间
	UpdatedAt                        *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                                                                          //更新时间
}

// TableName users表 Users自定义表名 users
func (Users) TableName() string {
	return "users"
}
