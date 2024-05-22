package models

import (
	"errors"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

type Customer struct {
	persistence.Record `mapstructure:",squash"`
	Name               string `mapstructure:"name" json:"name"`
	Email              string `mapstructure:"email" json:"email"`
	Phone              string `mapstructure:"phone" json:"phone"`
}

func NewCustomer(name string, email string, phone string) (*Customer, error) {
	if name == "" || email == "" || phone == "" {
		return nil, errors.New("error, Name, Email and Phone must be set")
	} else {
		return &Customer{
			Record: persistence.Record{},
			Name:   name,
			Email:  email,
			Phone:  phone,
		}, nil
	}
}
