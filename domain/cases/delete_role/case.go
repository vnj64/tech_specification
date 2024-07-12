package delete_role

import (
	"fmt"
	"tech/domain"
)

type Request struct {
	Id int `json:"id"`
}

func Run(c domain.Context, r Request) error {
	if err := validate(c); err != nil {
		return fmt.Errorf("unable to validate case [delete_role] due [%s]", err)
	}

	_, err := c.Connection().Role().GetRole(r.Id)
	if err != nil {
		return fmt.Errorf("role [%d] does not exist due [%s]", r.Id, err)
	}

	if err := c.Connection().Role().Delete(r.Id); err != nil {
		return fmt.Errorf("unable to delete role [%d] due [%s]", r.Id, err)
	}

	return nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
