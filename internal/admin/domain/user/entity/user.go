package entity

import "camellia/tool/field"

type User struct {
	field.TablePrimaryIdField
	RoleId uint64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;comment:'角色ID'"`
	Password string `gorm:"column:password;size:128;not null;comment:'密码';" json:"password"`
	Username string `gorm:"column:username;size:32;not null;unique_index; comment:'用户名';" json:"username"`
	field.TableTimeField
}


