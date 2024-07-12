package postgresql_driver

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tech/domain"
	"tech/domain/repositories"
)

type connection struct {
	db *gorm.DB

	userRepository  repositories.User
	rolesRepository repositories.Roles
}

func makeConnection(db *gorm.DB) *connection {
	return &connection{
		db:              db,
		userRepository:  &userRepository{db},
		rolesRepository: &roleRepository{db},
	}
}

func Make(user, password, host, port, database string) (domain.Connection, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		database,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open database due [%s]", err)
	}

	// Получить объект sql.DB для использования его методов
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to get DB object due [%s]", err)
	}

	// Ping
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping DB due [%s]", err)
	}

	return makeConnection(db), nil
}

func (c *connection) User() repositories.User {
	return c.userRepository
}

func (c *connection) Role() repositories.Roles {
	return c.rolesRepository
}
