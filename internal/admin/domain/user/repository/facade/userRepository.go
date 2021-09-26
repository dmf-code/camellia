package facade

import "camellia/internal/admin/domain/user/repository/po"

type UserRepository interface {
	Insert(userPO po.UserPO)
	Update(userPO po.UserPO)
	FindById(userPO po.UserPO) po.UserPO
	Delete(userPO po.UserPO)
}
