package po

import "camellia/tool/field"

type StaffPO struct {
	field.TablePrimaryIdField
	Uid int64 `gorm:"column:uid;not null;default: 0;comment:账号id;index:idx_uid;" json:"uid" form:"uid"`
	Email string `gorm:"column:email;not null;default: '';comment:员工邮箱;index:idx_email;" json:"email" form:"email"`
	Phone string `gorm:"column:phone;not null;default: '';comment:员工手机号;index:idx_phone;" json:"phone" form:"phone"`
	Name string `gorm:"column:name;not null;default: '';comment:员工姓名;" json:"name" form:"name"`
	Nickname string `gorm:"column:nickname;not null;default: '';comment:员工昵称;" json:"nickname" form:"nickname"`
	Avatar string `gorm:"column:avatar;not null;default: '';comment:员工头像(相对路径);" json:"avatar" form:"avatar"`
	Gender int8 `gorm:"column:gender;not null;default: 3;comment:员工性别: 1-男 2-女 3-未知;" json:"gender" form:"gender"`
	field.TableTimeField
}
