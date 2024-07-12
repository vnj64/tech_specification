package repositories

import "tech/domain/models"

type User interface {
	Insert(user models.User) (int, error)
	GetUser(id int) (*models.User, error)
	All() ([]models.User, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
}
