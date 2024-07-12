package get_all_roles

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Response struct {
	Roles []models.Role `json:"roles"`
}

func Run(c domain.Context) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [get_all_roles] due [%s]", err)
	}

	roles, err := c.Connection().Role().All()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve roles due [%s]", err)
	}

	return &Response{Roles: roles}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
