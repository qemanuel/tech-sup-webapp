package models

import (
	"errors"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

type Worker struct {
	persistence.Record `mapstructure:",squash"`
	Name               string `mapstructure:"name" json:"name"`
	Email              string `mapstructure:"email" json:"email"`
	Phone              string `mapstructure:"phone" json:"phone"`
}

func NewWorker(name string, email string, phone string) (Worker, error) {
	if name == "" || email == "" {
		return Worker{}, errors.New("error, Name and Email")
	} else {
		worker := Worker{
			Record: persistence.Record{},
			Name:   name,
			Email:  email,
			Phone:  phone,
		}
		return worker, nil
	}
}
