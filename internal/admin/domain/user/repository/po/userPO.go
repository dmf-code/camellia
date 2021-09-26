package po

import (
	"camellia/tool/field"
)

type UserPO struct {
	field.TablePrimaryIdField
	Password string `gorm:"column:password;size:128;not null;comment:'密码';" json:"password"`
	Username string `gorm:"column:username;size:32;not null;unique_index; comment:'用户名';" json:"username"`
	Avatar string `gorm:"column:avatar;size:255;default: ''; comment:'头像';" json:"avatar"`
	Phone string `gorm:"column:phone;size:16;default: null; comment:'手机号';" json:"phone"`
	Email string `gorm:"column:email;size:32;default: null;unique_index; comment:'邮箱';" json:"email"`
	field.TableTimeField
}
