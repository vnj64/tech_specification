package get_all_users

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Response struct {
	Users []models.User `json:"users"`
}

func Run(c domain.Context) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [get_all_users] due [%s]", err)
	}

	users, err := c.Connection().User().All()
	if err != nil {
		return nil, fmt.Errorf("unable to fetch users due [%s]", err)
	}

	return &Response{Users: users}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
