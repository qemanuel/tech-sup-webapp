package models

import (
	"errors"
)

type Worker struct {
	Name  string `mapstructure:"name" json:"name"`
	Email string `mapstructure:"email" json:"email"`
	Phone string `mapstructure:"phone" json:"phone"`
	Id    string `mapstructure:"id" json:"id"`
}

func NewWorker(name string, email string, phone string) (*Worker, error) {
	if name == "" || email == "" {
		return nil, errors.New("error, Name and Email")
	} else {
		worker := &Worker{
			Name:  name,
			Email: email,
			Phone: phone,
		}
		return worker, nil
	}
}

func GetWorker(work *Worker) Worker {
	worker := Worker{
		Name:  work.Name,
		Email: work.Email,
		Phone: work.Phone,
		Id:    work.Id,
	}
	return worker
}

func (work *Worker) SetId(id string) error {
	if work.Id != "" {
		return errors.New("[Error]: ID already assigned")
	} else {
		work.Id = id
		return nil
	}
}

func (work *Worker) GetId() string {
	return work.Id
}

func (work *Worker) String() []string {
	return []string{
		//fmt.Sprint(work.Id),
		work.Name,
		work.Email,
		work.Phone,
	}
}

func (worker *Worker) Update(name string, email string, phone string) error {
	if name == "" && email == "" && phone == "" {
		return errors.New("[Error]: Update details are not set")
	}
	if name != "" {
		worker.Name = name
	}
	if email != "" {
		worker.Email = email
	}
	if phone != "" {
		worker.Phone = phone
	}
	return nil
}
