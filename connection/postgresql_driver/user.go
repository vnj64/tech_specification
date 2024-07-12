package postgresql_driver

import (
	"gorm.io/gorm"
	"tech/domain/models"
	"time"
)

type userRepository struct {
	db *gorm.DB
}

type user struct {
	UserId     int `gorm:"primary_key"`
	Login      string
	FirstName  string
	SecondName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	RoleId     int
}

func makeUser(u models.User) user {
	return user{
		UserId:     u.UserId,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		RoleId:     u.RoleId,
	}
}

func (u user) model() *models.User {
	return &models.User{
		UserId:     u.UserId,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		RoleId:     u.RoleId,
	}
}

func (r *userRepository) Insert(user models.User) (int, error) {
	u := makeUser(user)

	result := r.db.Create(&u)
	if result.Error != nil {
		return 0, result.Error
	}

	return u.UserId, nil
}

func (r *userRepository) GetUser(id int) (*models.User, error) {
	var result user

	if err := r.db.Take(&result, `"user_id" = ?`, id).Error; err != nil {
		return nil, err
	}

	return result.model(), nil
}

func (r *userRepository) All() ([]models.User, error) {
	var result []user

	if err := r.db.Find(&result).Error; err != nil {
		return nil, err
	}

	out := make([]models.User, len(result))

	for i, u := range result {
		out[i] = *u.model()
	}

	return out, nil
}

func (r *userRepository) Update(id int, updates map[string]interface{}) error {
	return r.db.Model(user{UserId: id}).Updates(updates).Error
}

func (r *userRepository) Delete(userId int) error {
	return r.db.Delete(user{UserId: userId}).Error
}
