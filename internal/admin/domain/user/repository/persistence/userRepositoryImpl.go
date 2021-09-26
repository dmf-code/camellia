package persistence

import (
	"camellia/internal/admin/domain/user/repository/mapper"
	"camellia/internal/admin/domain/user/repository/po"
)

type UserRepositoryImpl struct {
	userDao mapper.UserDao
}

func (impl UserRepositoryImpl) Insert(userPO po.UserPO) {
	impl.userDao.Insert(userPO)
}

func (impl UserRepositoryImpl) Update() {

}

func (impl UserRepositoryImpl) FindById() po.UserPO {

	return po.UserPO{}
}

func (impl UserRepositoryImpl) Delete() {

}
