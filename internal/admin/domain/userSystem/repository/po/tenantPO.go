package po

import "camellia/tool/field"

type TenantPO struct {
	field.TablePrimaryIdField
	Name string `gorm:"column:name;size:64;not null;comment:企业名称;index:idx_name,unique;" json:"name" form:"name"`
	field.TableTimeField
}
