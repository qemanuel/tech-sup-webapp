package models

import (
	"errors"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

type Incidence struct {
	persistence.Record `mapstructure:",squash"`
	JobId              string `mapstructure:"job_id" json:"job_id"`
	Body               string `mapstructure:"body" json:"body"`
	NextId             string `mapstructure:"next_id" json:"next_id"`
	AuthorId           string `mapstructure:"author_id" json:"author_id"`
}

func NewIncidence(jobId string, body string, nextId string, authorId string) (Incidence, error) {
	if jobId == "" || body == "" || authorId == "" {
		return Incidence{}, errors.New("[Error]: job-id, body and author-id are required")
	}
	return Incidence{
		JobId:    jobId,
		Body:     body,
		NextId:   nextId,
		AuthorId: authorId,
	}, nil
}
