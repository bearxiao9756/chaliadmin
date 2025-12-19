
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type TeamChaliBotCommandsService struct {}
// CreateTeamChaliBotCommands 创建业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService) CreateTeamChaliBotCommands(ctx context.Context, TeamChalibotCommands *system.TeamChaliBotCommands) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Create(TeamChalibotCommands).Error
	return err
}

// DeleteTeamChaliBotCommands 删除业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)DeleteTeamChaliBotCommands(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&system.TeamChaliBotCommands{},"id = ?",id).Error
	return err
}

// DeleteTeamChaliBotCommandsByIds 批量删除业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)DeleteTeamChaliBotCommandsByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Delete(&[]system.TeamChaliBotCommands{},"id in ?",ids).Error
	return err
}

// UpdateTeamChaliBotCommands 更新业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)UpdateTeamChaliBotCommands(ctx context.Context, TeamChalibotCommands system.TeamChaliBotCommands) (err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamChaliBotCommands{}).Where("id = ?",TeamChalibotCommands.Id).Updates(&TeamChalibotCommands).Error
	return err
}

// GetTeamChaliBotCommands 根据id获取业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)GetTeamChaliBotCommands(ctx context.Context, id string) (TeamChalibotCommands system.TeamChaliBotCommands, err error) {
	err = global.MustGetGlobalDBByDBName("teamgramDB").Where("id = ?", id).First(&TeamChalibotCommands).Error
	return
}
// GetTeamChaliBotCommandsInfoList 分页获取业务机器人命令记录
// Author [yourname](https://github.com/yourname)
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)GetTeamChaliBotCommandsInfoList(ctx context.Context, info systemReq.TeamChaliBotCommandsSearch) (list []system.TeamChaliBotCommands, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("teamgramDB").Model(&system.TeamChaliBotCommands{})
    var TeamChalibotCommandss []system.TeamChaliBotCommands
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
    if info.BotId != nil {
        db = db.Where("bot_id = ?", *info.BotId)
    }
    if info.Command != nil && *info.Command != "" {
        db = db.Where("command = ?", *info.Command)
    }
    if info.Description != nil && *info.Description != "" {
        db = db.Where("description LIKE ?", "%"+ *info.Description+"%")
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

	err = db.Find(&TeamChalibotCommandss).Error
	return  TeamChalibotCommandss, total, err
}
func (TeamChalibotCommandsService *TeamChaliBotCommandsService)GetTeamChaliBotCommandsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
