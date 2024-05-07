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
	}
	return worker
}

func (work *Worker) SetWorkerId(id int) error {
	if work.id != 0 {
		return errors.New("[Error]: The worker already has an ID")
	} else {
		work.id = id
		return nil
	}
}

func (work *Worker) GetWorkerId() int {
	return work.id
}

func (work *Worker) StringWorker() []string {
	return []string{
		work.name,
		work.email,
		work.phone,
	}
}

func (worker *Worker) UpdateWorkerInfo(email string, phone string) error {
	var err error
	if email == "" && phone == "" {
		err = errors.New("[Error]: Info details are not set")
	}
	if email != "" {
		worker.email = email
	}
	if phone != "" {
		worker.phone = phone
	}
	return err
}
