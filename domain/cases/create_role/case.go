package create_role

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Request struct {
	Role models.Role `json:"role"`
}

type Response struct {
	RoleId int `json:"roleId"`
}

func Run(c domain.Context, r Request) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [create_role] due [%s]", err)
	}

	roleId, err := c.Connection().Role().Insert(r.Role)
	if err != nil {
		return nil, fmt.Errorf("unable to create role due [%s]", err)
	}

	r.Role.RoleId = roleId

	return &Response{RoleId: r.Role.RoleId}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
