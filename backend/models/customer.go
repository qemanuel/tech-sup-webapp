package models

import (
	"errors"
)

type Customer struct {
	Name  string `mapstructure:"name" json:"name"`
	Email string `mapstructure:"email" json:"email"`
	Phone string `mapstructure:"phone" json:"phone"`
	Id    string `mapstructure:"id" json:"id"`
}

func NewCustomer(name string, email string, phone string) (*Customer, error) {
	if name == "" || email == "" || phone == "" {
		return nil, errors.New("error, Name, Email and Phone must be set")
	} else {
		return &Customer{
			Name:  name,
			Email: email,
			Phone: phone,
		}, nil
	}
}

func GetCustomer(cust *Customer) Customer {
	customer := Customer{
		Name:  cust.Name,
		Email: cust.Email,
		Phone: cust.Phone,
		Id:    cust.Id,
	}
	return customer
}

func (cust *Customer) SetId(id string) error {
	if cust.Id != "" {
		return errors.New("[Error]: ID already assigned")
	} else {
		cust.Id = id
		return nil
	}
}

func (cust *Customer) GetId() string {
	return cust.Id
}

func (cust *Customer) String() []string {
	return []string{
		cust.Name,
		cust.Email,
		cust.Phone,
	}
}

func (cust *Customer) Update(name string, email string, phone string) error {
	if name == "" && email == "" && phone == "" {
		return errors.New("[Error]: Update details are not set")
	}
	if name != "" {
		cust.Name = name
	}
	if email != "" {
		cust.Email = email
	}
	if phone != "" {
		cust.Phone = phone
	}
	return nil
}
