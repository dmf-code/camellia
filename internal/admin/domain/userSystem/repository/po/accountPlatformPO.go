package po

import "camellia/tool/field"

type AccountPlatformPO struct {
	field.TablePrimaryIdField
	Uid int64 `gorm:"column:uid;not null;default: 0;comment:账号id;index:idx_uid;" json:"uid" form:"uid"`
	PlatformId string `gorm:"column:platform_id;not null;default: '';comment:平台id;index:idx_platformId;" json:"platform_id" form:"platform_id"`
	PlatformToken string `gorm:"column:platform_token;not null;default: '';comment:平台access_token;" json:"platform_token" form:"platform_token"`
	Type string `gorm:"column:type;not null;default: '0';comment:平台类型 0:未知,1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter;" json:"type" form:"type"`
	Nickname string `gorm:"column:nickname;not null;default: '';comment:昵称;" json:"nickname" form:"nickname"`
	Avatar string `gorm:"column:avatar;not null;default: '';comment:头像;" json:"avatar" form:"avatar"`
	field.TableTimeField
}

func (AccountPlatformPO) TableName() string {
	return "account_platform"
}
