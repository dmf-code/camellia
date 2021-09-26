package po

import (
	"camellia/tool/field"
)

type RolePO struct {
	field.TablePrimaryIdField
	UserId uint64 `gorm:"column:user_id;not null;comment:'用户id';unique_index;" json:"user_id" form:"user_id"`
	Name string `gorm:"column:name;size:32;not null;comment:'名称';unique_index;" json:"name" form:"name"`
	Memo string `gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	Sequence uint64 `gorm:"column:sequence;not null;default: 5;comment:'排序值'" json:"sequence" form:"sequence"`
	field.TableTimeField
}
