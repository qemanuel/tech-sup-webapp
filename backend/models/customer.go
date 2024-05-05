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
	}
	return customer
}

func (cust *Customer) SetCustomerId(id map[int]*Customer) {

}

func (cust *Customer) UpdateCustomerInfo(email string, phone string) error {
	var err error
	if email == "" && phone == "" {
		err = errors.New("[Error]: Info details are not set")
	}
	if email != "" {
		cust.email = email
	}
	if phone != "" {
		cust.phone = phone
	}
	return err
}
