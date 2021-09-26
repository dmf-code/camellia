package field

import "camellia/tool/format"
// 表主键
type TablePrimaryIdField struct {
	Id	uint64	`gorm:"primary_key" json:"id"`
}

// 表时间字段
type TableTimeField struct {
	CreatedAt format.JSONTime  `json:"created_at"`
	UpdatedAt format.JSONTime  `json:"updated_at"`
	DeletedAt *format.JSONTime `json:"-"`
}
