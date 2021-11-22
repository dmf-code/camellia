package po

import "camellia/tool/field"

type AccountPO struct {
	field.TablePrimaryIdField
	TenantId int32 `gorm:"column:tenant_id;not null;default: 0;comment:租户id;" json:"tenant_id" form:"tenant_id"`
	Email string `gorm:"column:email;not null;default: '';comment:邮箱;index:idx_email;" json:"email" form:"email"`
	Phone string `gorm:"column:phone;not null;default: '';comment:手机号;index:idx_phone;" json:"phone" form:"phone"`
	Username string `gorm:"column:username;not null;default: '';comment:用户名;index:idx_username;" json:"username" form:"username"`
	Password string `gorm:"column:password;not null;default: '';comment:密码;" json:"password" form:"password"`
	CreateIpAt string `gorm:"column:create_ip_at;not null;default: '';comment:创建ip;" json:"create_ip_at" form:"create_ip_at"`
	LastLoginAt int32 `gorm:"column:last_login_at;not null;default: 0;comment:最后一次登录时间;" json:"last_login_at" form:"last_login_at"`
	LastLoginIpAt string `gorm:"column:last_login_ip_at;not null;default: '';comment:最后一次登录ip;" json:"last_login_ip_at" form:"last_login_ip_at"`
	LoginTimes int32 `gorm:"column:login_times;not null;default: 0;comment:登录次数;" json:"login_times" form:"login_times"`
	Status int8 `gorm:"column:status;not null;default: 2;comment:状态 1:disable, 2:enable, 3:deleted;" json:"status" form:"status"`
	field.TableTimeField
}

func (AccountPO) TableName() string {
	return "account"
}
