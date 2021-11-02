package po

import "camellia/tool/field"

type MemberPO struct {
	field.TablePrimaryIdField
	Uid int64 `gorm:"column:uid;not null;default: '0';comment:'账号id';index:idx-uid;" json:"uid" form:"uid"`
	Nickname string `gorm:"column:nickname;not null;default: '';comment:'昵称';" json:"nickname" form:"nickname"`
	Avatar string `gorm:"column:avatar;not null;default: '';comment:'头像(相对路径)';" json:"avatar" form:"avatar"`
	Gender int8 `gorm:"column:gender;not null;default: '0';comment:'性别:0-未知 1-男生 2-女生';" json:"gender" form:"gender"`
	Role int8 `gorm:"column:create_ip_at;not null;default: '';comment:'创建ip';" json:"create_ip_at" form:"create_ip_at"`
	LastLoginAt int32 `gorm:"column:last_login_at;not null;default: '0';comment:'最后一次登录时间';" json:"last_login_at" form:"last_login_at"`
	LastLoginIpAt string `gorm:"column:last_login_ip_at;not null;default: '';comment:'最后一次登录ip';" json:"last_login_ip_at" form:"last_login_ip_at"`
	LoginTimes int32 `gorm:"column:login_times;not null;default: '0';comment:'登录次数';" json:"login_times" form:"login_times"`
	Status int8 `gorm:"column:status;not null;default: '0';comment:'状态 1:enable, 0:disable, -1:deleted';" json:"status" form:"status"`
	field.TableTimeField
}
