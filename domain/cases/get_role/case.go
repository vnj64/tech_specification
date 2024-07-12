package get_role

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Request struct {
	Id int `json:"id"`
}

type Response struct {
	Role *models.Role `json:"role"`
}

func Run(c domain.Context, r Request) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [get_role] due [%s]", err)
	}

	role, err := c.Connection().Role().GetRole(r.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to get role with id [%d] due [%s]", r.Id, err)
	}

	return &Response{Role: role}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
