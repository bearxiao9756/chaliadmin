package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	teamgramDBDb := global.GetGlobalDBByDBName("teamgramDB")
	teamgramDBDb.AutoMigrate()
	return nil
}
