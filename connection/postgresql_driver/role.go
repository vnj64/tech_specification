package postgresql_driver

import (
	"gorm.io/gorm"
	"tech/domain/models"
	"time"
)

type roleRepository struct {
	db *gorm.DB
}

type role struct {
	RoleId      int `gorm:"primary_key"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r role) model() *models.Role {
	return &models.Role{
		RoleId:      r.RoleId,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func makeRole(r models.Role) role {
	return role{
		RoleId:      r.RoleId,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func (r *roleRepository) Insert(role models.Role) (int, error) {
	rm := makeRole(role)

	result := r.db.Create(&rm)
	if result.Error != nil {
		return 0, result.Error
	}

	return rm.RoleId, nil
}

func (r *roleRepository) GetRole(id int) (*models.Role, error) {
	var result role

	if err := r.db.Take(&result, `"role_id" = ?`, id).Error; err != nil {
		return nil, err
	}

	return result.model(), nil
}

func (r *roleRepository) All() ([]models.Role, error) {
	var result []role

	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}

	out := make([]models.Role, len(result))

	for i, rm := range result {
		out[i] = *rm.model()
	}

	return out, nil
}

func (r *roleRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(role{RoleId: id}).Updates(updates).Error
}

func (r *roleRepository) Delete(roleId int) error {
	return r.db.Delete(role{RoleId: roleId}).Error
}
