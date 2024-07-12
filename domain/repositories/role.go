package repositories

import "tech/domain/models"

type Roles interface {
	Insert(role models.Role) (int, error)
	GetRole(id int) (*models.Role, error)
	All() ([]models.Role, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
}
