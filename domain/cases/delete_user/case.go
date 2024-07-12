package delete_user

import (
	"fmt"
	"tech/domain"
)

type Request struct {
	Id int `json:"id"`
}

func Run(c domain.Context, r Request) error {
	if err := validate(c); err != nil {
		return fmt.Errorf("unable to validate case [delete_user] due [%s]", err)
	}

	_, err := c.Connection().User().GetUser(r.Id)
	if err != nil {
		return fmt.Errorf("user [%d] does not exists due [%s]", r.Id, err)
	}

	if err := c.Connection().User().Delete(r.Id); err != nil {
		return fmt.Errorf("unable to delete user [%d] due [%s]", r.Id, err)
	}

	return nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
