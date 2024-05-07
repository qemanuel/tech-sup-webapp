package models

import (
	"errors"
)

type Worker struct {
	name  string
	email string
	phone string
	id    int
}

func NewWorker(name string, email string, phone string) (*Worker, error) {
	if name == "" || email == "" {
		return nil, errors.New("error, Name and Email")
	} else {
		worker := &Worker{
			name:  name,
			email: email,
			phone: phone,
		}
		return worker, nil
	}
}

func GetWorker(work *Worker) Worker {
	worker := Worker{
		name:  work.name,
		email: work.email,
		phone: work.phone,
		id:    work.id,
	}
	return worker
}

func (work *Worker) SetId(id int) error {
	if work.id != 0 {
		return errors.New("[Error]: ID already assigned")
	} else {
		work.id = id
		return nil
	}
}

func (work *Worker) GetId() int {
	return work.id
}

func (work *Worker) String() []string {
	return []string{
		work.name,
		work.email,
		work.phone,
	}
}

func (worker *Worker) Update(name string, email string, phone string) error {
	if name == "" && email == "" && phone == "" {
		return errors.New("[Error]: Update details are not set")
	}
	if name != "" {
		worker.name = name
	}
	if email != "" {
		worker.email = email
	}
	if phone != "" {
		worker.phone = phone
	}
	return nil
}
