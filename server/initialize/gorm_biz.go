package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	teamgramDBDb := global.GetGlobalDBByDBName("teamgramDB")
	teamgramDBDb.AutoMigrate(system.TeamgramDrafts{})
	return nil
}
