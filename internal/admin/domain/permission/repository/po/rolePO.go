package po

import (
	"camellia/tool/field"
)

type RolePO struct {
	field.TablePrimaryIdField
	//TODO 之后把pid改为雪花算法id
	Pid string `gorm:"column:pid;default: 0;comment:'0-根节点';unique_index;" json:"pid" form:"pid"`
	Name string `gorm:"column:name;size:32;not null;comment:'名称';unique_index;" json:"name" form:"name"`
	Memo string `gorm:"column:memo;size:64;comment:'备注'" json:"memo" form:"memo"`
	Sequence uint64 `gorm:"column:sequence;not null;default: 5;comment:'排序值'" json:"sequence" form:"sequence"`
	field.TableTimeField
}
