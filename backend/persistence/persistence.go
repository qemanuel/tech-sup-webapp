package persistence

import (
	"errors"

	"github.com/qemanuel/tech-sup-webapp/backend/models"
)

func SetId(object models.Id, id int) error {
	var err error
	if id == 0 {
		err = errors.New("[Error]: Id missing")
	} else if object.Id == 0 {
		err = errors.New("[Error]: Id already set")
	} else {
		object.Id = id
	}
	return err
}
