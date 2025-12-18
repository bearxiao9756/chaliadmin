
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type UsersService struct {}
// CreateUsers 创建users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) CreateUsers(ctx context.Context, users *system.Users) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(users).Error
	return err
}

// DeleteUsers 删除users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsers(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.Users{},"id = ?",id).Error
	return err
}

// DeleteUsersByIds 批量删除users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsersByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.Users{},"id in ?",ids).Error
	return err
}

// UpdateUsers 更新users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)UpdateUsers(ctx context.Context, users system.Users) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.Users{}).Where("id = ?",users.Id).Updates(&users).Error
	return err
}

// GetUsers 根据id获取users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsers(ctx context.Context, id string) (users system.Users, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&users).Error
	return
}
// GetUsersInfoList 分页获取users表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsersInfoList(ctx context.Context, info systemReq.UsersSearch) (list []system.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.Users{})
    var userss []system.Users
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userss).Error
	return  userss, total, err
}
func (usersService *UsersService)GetUsersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
