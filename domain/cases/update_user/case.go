package update_user

import (
	"fmt"
	"tech/domain"
)

type Request struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	RoleId     int    `json:"roleId"`
}

func Run(c domain.Context, r Request) error {
	if err := validate(c); err != nil {
		return fmt.Errorf("unable to validate case [update_user] due [%s]", err)
	}

	userUpdates := make(map[string]interface{})
	if r.Login != "" {
		userUpdates["Login"] = r.Login
	}

	if r.FirstName != "" {
		userUpdates["FirstName"] = r.FirstName
	}

	if r.SecondName != "" {
		userUpdates["SecondName"] = r.SecondName
	}

	_, err := c.Connection().Role().GetRole(r.RoleId)
	if err != nil {
		return fmt.Errorf("role [%d] does not exist", r.RoleId)
	}

	if r.RoleId > 0 {
		userUpdates["RoleId"] = r.RoleId
	}

	if err := c.Connection().User().Update(r.Id, userUpdates); err != nil {
		return fmt.Errorf("unable to update user [%d] due [%s]", r.Id, err)
	}

	return nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
