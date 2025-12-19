
// 自动生成模板Auths
package example
import (
	"time"
)

// auths表 结构体  Auths
type Auths struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //编号
  AuthKeyId  *int64 `json:"authKeyId" form:"authKeyId" gorm:"column:auth_key_id;"`  //authKeyId
  LangPack  *string `json:"langPack" form:"langPack" gorm:"column:lang_pack;size:255;"`  //设备类型
  ApiId  *int32 `json:"apiId" form:"apiId" gorm:"column:api_id;"`  //apiId
  ClientIp  *string `json:"clientIp" form:"clientIp" gorm:"column:client_ip;size:32;"`  //设备IP
  DeviceModel  *string `json:"deviceModel" form:"deviceModel" gorm:"column:device_model;size:255;"`  //设备名称
  SystemLangCode  *string `json:"systemLangCode" form:"systemLangCode" gorm:"column:system_lang_code;size:255;"`  //系统语言代码
  SystemVersion  *string `json:"systemVersion" form:"systemVersion" gorm:"column:system_version;size:255;"`  //系统版本
  AppVersion  *string `json:"appVersion" form:"appVersion" gorm:"column:app_version;size:255;"`  //软件版本
  Proxy  *string `json:"proxy" form:"proxy" gorm:"column:proxy;size:512;"`  //代理
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //更新时间
}


// TableName auths表 Auths自定义表名 auths
func (Auths) TableName() string {
    return "auths"
}





