package bootstrap

import (
	"fmt"
	"lucian/internal/model"
	"lucian/tool/helper"
)

func InitTable() {
	db := helper.Db()
	fmt.Println(db.AutoMigrate(new(model.Menu)).Error)
	fmt.Println(db.AutoMigrate(new(model.Role)).Error)
	fmt.Println(db.AutoMigrate(new(model.Admin)).Error)
	fmt.Println(db.AutoMigrate(new(model.RoleMenu)).Error)
	fmt.Println(db.AutoMigrate(new(model.AdminRole)).Error)
	fmt.Println(db.AutoMigrate(new(model.Article)).Error)
	fmt.Println(db.AutoMigrate(new(model.Tag)).Error)
	fmt.Println(db.AutoMigrate(new(model.Category)).Error)
	fmt.Println(db.AutoMigrate(new(model.Tutorial)).Error)
	//fmt.Println(db.AutoMigrate(new(tutorial.MenuTutorial)).Error)
	fmt.Println(db.AutoMigrate(new(model.ContentTutorial)).Error)
	fmt.Println(db.AutoMigrate(new(model.Nav)).Error)
	fmt.Println(db.AutoMigrate(new(model.Bookmark)))
}
