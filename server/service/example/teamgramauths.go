
package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
)

type AuthsService struct {}
// CreateAuths 创建auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService) CreateAuths(ctx context.Context, auths *example.Auths) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(auths).Error
	return err
}

// DeleteAuths 删除auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService)DeleteAuths(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&example.Auths{},"id = ?",id).Error
	return err
}

// DeleteAuthsByIds 批量删除auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService)DeleteAuthsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]example.Auths{},"id in ?",ids).Error
	return err
}

// UpdateAuths 更新auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService)UpdateAuths(ctx context.Context, auths example.Auths) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&example.Auths{}).Where("id = ?",auths.Id).Updates(&auths).Error
	return err
}

// GetAuths 根据id获取auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService)GetAuths(ctx context.Context, id string) (auths example.Auths, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&auths).Error
	return
}
// GetAuthsInfoList 分页获取auths表记录
// Author [yourname](https://github.com/yourname)
func (authsService *AuthsService)GetAuthsInfoList(ctx context.Context, info exampleReq.AuthsSearch) (list []example.Auths, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&example.Auths{})
    var authss []example.Auths
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.AuthKeyId != nil {
        db = db.Where("auth_key_id = ?", *info.AuthKeyId)
    }
    if info.LangPack != nil && *info.LangPack != "" {
        db = db.Where("lang_pack LIKE ?", "%"+ *info.LangPack+"%")
    }
    if info.ApiId != nil {
        db = db.Where("api_id = ?", *info.ApiId)
    }
    if info.ClientIp != nil && *info.ClientIp != "" {
        db = db.Where("client_ip LIKE ?", "%"+ *info.ClientIp+"%")
    }
    if info.DeviceModel != nil && *info.DeviceModel != "" {
        db = db.Where("device_model LIKE ?", "%"+ *info.DeviceModel+"%")
    }
    if info.SystemLangCode != nil && *info.SystemLangCode != "" {
        db = db.Where("system_lang_code LIKE ?", "%"+ *info.SystemLangCode+"%")
    }
    if info.SystemVersion != nil && *info.SystemVersion != "" {
        db = db.Where("system_version LIKE ?", "%"+ *info.SystemVersion+"%")
    }
    if info.AppVersion != nil && *info.AppVersion != "" {
        db = db.Where("app_version LIKE ?", "%"+ *info.AppVersion+"%")
    }
    if info.Proxy != nil && *info.Proxy != "" {
        db = db.Where("proxy LIKE ?", "%"+ *info.Proxy+"%")
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

	err = db.Find(&authss).Error
	return  authss, total, err
}
func (authsService *AuthsService)GetAuthsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
