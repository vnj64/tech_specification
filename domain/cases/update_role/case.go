package update_role

import (
	"fmt"
	"tech/domain"
	"time"
)

type Request struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Run(c domain.Context, r Request) error {
	if err := validate(c); err != nil {
		return fmt.Errorf("unable to validate case [update_role] due [%s]", err)
	}

	_, err := c.Connection().Role().GetRole(r.Id)
	if err != nil {
		return fmt.Errorf("role [%d] does not exists due [%s]", r.Id, err)
	}

	roleUpdates := make(map[string]interface{})
	if r.Name != "" {
		roleUpdates["Name"] = r.Name
	}

	if r.Description != "" {
		roleUpdates["Description"] = r.Description
	}
	roleUpdates["UpdatedAt"] = time.Now()

	if err := c.Connection().Role().Update(r.Id, roleUpdates); err != nil {
		return fmt.Errorf("unable to update role [%d] due [%s]", r.Id, err)
	}

	return nil
}

func validate(c domain.Context) error {
	return domain.ValidateContext(c)
}
