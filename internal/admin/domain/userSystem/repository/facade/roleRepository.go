package facade

import "camellia/internal/admin/domain/userSystem/repository/po"

type RoleRepository interface {
	Insert(rolePO po.RolePO)
	Update(rolePO po.RolePO)
	FindById(rolePO po.RolePO) po.RolePO
	Delete(rolePO po.RolePO)
}
