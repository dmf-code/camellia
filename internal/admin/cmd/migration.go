package cmd

import (
	"camellia/internal/admin/domain/userSystem/repository/po"
	"camellia/tool/database"
	"camellia/tool/exception"
	"fmt"
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
	db := database.GetInstance()

	fmt.Println(db)

	err := db.Migrator().AutoMigrate(&po.AccountPO{})
	if err != nil {
		exception.GetIns().Throw(err)
	}

	err = db.Migrator().AutoMigrate(&po.MemberPO{})
	if err != nil {
		exception.GetIns().Throw(err)
	}

	err = db.Migrator().AutoMigrate(&po.RolePO{})
	if err != nil {
		exception.GetIns().Throw(err)
	}

	err = db.Migrator().AutoMigrate(&po.StaffPO{})
	if err != nil {
		exception.GetIns().Throw(err)
	}

	err = db.Migrator().AutoMigrate(&po.TenantPO{})
	if err != nil {
		exception.GetIns().Throw(err)
	}
}

