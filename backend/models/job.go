package models

import (
	"errors"
	"slices"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

type Job struct {
	persistence.Record `mapstructure:",squash"`
	DeviceId           string `mapstructure:"device_id" json:"device_id"`
	Status             string `mapstructure:"status" json:"status"`
	Reason             string `mapstructure:"reason" json:"reason"`
	Observations       string `mapstructure:"observations" json:"observations"`
	AuthorId           string `mapstructure:"author_id" json:"author_id"`
	AssignedId         string `mapstructure:"assigned_id" json:"assigned_id"`
}

func NewJob(deviceId string, reason string, observations string, status string, assignedId string, authorId string) (Job, error) {

	validStatus := []string{"ingressed", "in-progress", "on-hold", "finished", "egressed"}
	if !slices.Contains(validStatus, status) {
		return Job{}, errors.New("[Error]: Status invalid")
	}
	if deviceId == "" || reason == "" || authorId == "" {
		return Job{}, errors.New("[Error]: device-id, reason and author-id are required")
	}

	return Job{
		DeviceId:     deviceId,
		Status:       status,
		Reason:       reason,
		Observations: observations,
		AuthorId:     authorId,
		AssignedId:   assignedId,
	}, nil
}
