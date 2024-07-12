package create_user

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Request struct {
	User models.User `json:"user"`
}

type Response struct {
	Id int `json:"id"`
}

func Run(c domain.Context, r Request) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [create_user] due [%s]", err)
	}

	id, err := c.Connection().User().Insert(r.User)
	if err != nil {
		return nil, fmt.Errorf("unable to create user due [%s]", err)
	}

	return &Response{Id: id}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
