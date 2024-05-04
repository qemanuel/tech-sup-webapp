package models

import (
	"errors"
)

type Job struct {
	device       *Device
	status       string
	observations string
	history      *IncidenceList
	Id           Id
	customer     *Customer
	responsible  *Worker
	autor        *Worker
}

func NewJob(device *Device, reason string, observations string, customer *Customer, responsible *Worker, autor *Worker) (*Job, error) {
	if reason == "" || customer == nil || autor == nil || device == nil {
		return nil, errors.New("[Error]: there are missing inputs")
	}
	if responsible == nil {
		responsible = autor
	}

	return &Job{
		device:       device,
		status:       "ingressed",
		observations: observations,
		customer:     customer,
		responsible:  responsible,
		autor:        autor,
	}, nil
}

func (job *Job) UpdateJobStatus(status string) error {
	var err error
	switch status {
	case "working":
		if job.status != "ingressed" && job.status != "waiting" {
			err = errors.New("[Error]: status invalid")
		} else {
			job.status = status
		}
	case "waiting":
		if job.status != "ingressed" && job.status != "working" {
			err = errors.New("[Error]: status invalid")
		} else {
			job.status = status
		}
	case "egressed":
		if job.status != "ingressed" && job.status != "working" && job.status != "waiting" {
			err = errors.New("[Error]: status invalid")
		} else {
			job.status = status
		}
	default:
		err = errors.New("[Error]: status invalid")
	}
	return err
}

func (job *Job) UpdateJobResponsible(responsible *Worker) error {
	var err error
	if responsible == nil {
		err = errors.New("[Error]: Responsible can't be nil")
	} else {
		job.responsible = responsible
	}
	return err
}

func (job *Job) UpdateJobCustomer(customer *Customer) error {
	var err error
	if customer == nil {
		err = errors.New("[Error]: Customer can't be nil")
	} else {
		job.customer = customer
	}
	return err
}
