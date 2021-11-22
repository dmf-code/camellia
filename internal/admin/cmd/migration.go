package cmd

import (
	"camellia/internal/admin/domain/userSystem/repository/po"
	"camellia/tool/database"
	"camellia/tool/exception"
	"github.com/spf13/cobra"
)

var MigrationCmd = &cobra.Command{
	Use:   "Migration",
	Short: "系统数据初始化",
	Run: func(cmd *cobra.Command, args []string) {
		initTableData()
	},
}

func initTableData() {
	var err error
	db := database.GetInstance()

	if db.Migrator().HasTable("account") == false {
		err = db.Migrator().AutoMigrate(&po.AccountPO{})
		exception.GetIns().Throw(err)
	}

	if db.Migrator().HasTable("account_platform") == false {
		err = db.Migrator().AutoMigrate(&po.AccountPlatformPO{})
		exception.GetIns().Throw(err)
	}

	if db.Migrator().HasTable("member") == false {
		err = db.Migrator().AutoMigrate(&po.MemberPO{})
		exception.GetIns().Throw(err)
	}

	if db.Migrator().HasTable("role") == false {
		err = db.Migrator().AutoMigrate(&po.RolePO{})
		exception.GetIns().Throw(err)
	}

	if db.Migrator().HasTable("staff") == false {
		err = db.Migrator().AutoMigrate(&po.StaffPO{})
		exception.GetIns().Throw(err)
	}

	if db.Migrator().HasTable("tenant") == false {
		err = db.Migrator().AutoMigrate(&po.TenantPO{})
		exception.GetIns().Throw(err)
	}

}

