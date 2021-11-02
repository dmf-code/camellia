package po

import "camellia/tool/field"

type AccountPO struct {
	field.TablePrimaryIdField
	Email string `gorm:"column:email;not null;default: '';comment:'邮箱';index:idx-email;" json:"email" form:"email"`
	Phone string `gorm:"column:phone;not null;default: '';comment:'手机号';index:idx-phone;" json:"phone" form:"phone"`
	Username string `gorm:"column:username;not null;default: '';comment:'用户名';index:idx-username;" json:"username" form:"username"`
	Password string `gorm:"column:password;not null;default: '';comment:'密码';" json:"password" form:"password"`
	CreateIpAt string `gorm:"column:create_ip_at;not null;default: '';comment:'创建ip';" json:"create_ip_at" form:"create_ip_at"`
	LastLoginAt int32 `gorm:"column:last_login_at;not null;default: '0';comment:'最后一次登录时间';" json:"last_login_at" form:"last_login_at"`
	LastLoginIpAt string `gorm:"column:last_login_ip_at;not null;default: '';comment:'最后一次登录ip';" json:"last_login_ip_at" form:"last_login_ip_at"`
	LoginTimes int32 `gorm:"column:login_times;not null;default: '0';comment:'登录次数';" json:"login_times" form:"login_times"`
	Status int8 `gorm:"column:status;not null;default: '0';comment:'状态 1:enable, 0:disable, -1:deleted';" json:"status" form:"status"`
	field.TableTimeField
}

type AccountPlatformPO struct {
	field.TablePrimaryIdField
	Uid int64 `gorm:"column:uid;not null;default: '0';comment:'账号id';index:idx-uid;" json:"uid" form:"uid"`
	PlatformId string `gorm:"column:platform_id;not null;default: '';comment:'平台id';index:idx-platform_id;" json:"platform_id" form:"platform_id"`
	PlatformToken string `gorm:"column:platform_token;not null;default: '';comment:'平台access_token';" json:"platform_token" form:"platform_token"`
	Type string `gorm:"column:type;not null;default: '0';comment:'平台类型 0:未知,1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter';" json:"type" form:"type"`
	Nickname string `gorm:"column:nickname;not null;default: '';comment:'昵称';" json:"nickname" form:"nickname"`
	Avatar string `gorm:"column:avatar;not null;default: '';comment:'头像';" json:"avatar" form:"avatar"`
	field.TableTimeField
}
