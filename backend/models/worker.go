package models

import "errors"

type Worker struct {
	name  string
	email string
	phone string
	Id    Id
}

func NewWorker(name string, email string, phone string) (*Worker, error) {
	if name == "" || email == "" {
		return nil, errors.New("error, Name and Email")
	} else {
		return &Worker{
			name:  name,
			email: email,
			phone: phone,
		}, nil
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
