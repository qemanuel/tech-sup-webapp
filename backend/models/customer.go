package models

import (
	"errors"
)

type Customer struct {
	name  string
	email string
	phone string
	id    int
}

func NewCustomer(name string, email string, phone string) (*Customer, error) {
	if name == "" || email == "" || phone == "" {
		return nil, errors.New("error, Name, Email and Phone must be set")
	} else {
		return &Customer{
			name:  name,
			email: email,
			phone: phone,
		}, nil
	}
}

func GetCustomer(cust *Customer) Customer {
	customer := Customer{
		name:  cust.name,
		email: cust.email,
		phone: cust.phone,
		id:    cust.id,
	}
	return customer
}

func (cust *Customer) SetId(id int) error {
	if id != 0 {
		return errors.New("[Error]: ID already assigned")
	} else {
		cust.id = id
	}
	return nil
}

func (cust *Customer) GetId() int {
	return cust.id
}

func (cust *Customer) String() []string {
	return []string{
		cust.name,
		cust.email,
		cust.phone,
	}
}

func (cust *Customer) Update(name string, email string, phone string) error {
	if name == "" && email == "" && phone == "" {
		return errors.New("[Error]: Update details are not set")
	}
	if name != "" {
		cust.name = name
	}
	if email != "" {
		cust.email = email
	}
	if phone != "" {
		cust.phone = phone
	}
	return nil
}
