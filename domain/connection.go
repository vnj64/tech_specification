package domain

import "tech/domain/repositories"

type Connection interface {
	User() repositories.User
	Role() repositories.Roles
}
