package persistence

import (
	"camellia/internal/admin/domain/userSystem/repository/mapper"
	"camellia/internal/admin/domain/user/repository/po"
)

type RoleRepositoryImpl struct {
	roleDao mapper.RoleDao
}

func (impl RoleRepositoryImpl) Insert(userPO po.UserPO) {
	impl.roleDao.Insert(userPO)
}

func (impl RoleRepositoryImpl) Update() {

}

func (impl RoleRepositoryImpl) FindById() po.UserPO {

	return po.UserPO{}
}

func (impl RoleRepositoryImpl) Delete() {

}
