package get_user

import (
	"fmt"
	"tech/domain"
	"tech/domain/models"
)

type Request struct {
	Id int `json:"id"`
}

type Response struct {
	User *models.User `json:"user"`
}

func Run(c domain.Context, r Request) (*Response, error) {
	if err := validate(c); err != nil {
		return nil, fmt.Errorf("unable to validate case [get_user] due [%s]", err)
	}

	user, err := c.Connection().User().GetUser(r.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to get user with id [%d] due [%s]", r.Id, err)
	}

	return &Response{User: user}, nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
