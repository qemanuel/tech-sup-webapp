package models

import (
	"errors"

	"github.com/qemanuel/tech-sup-webapp/backend/persistence"
)

type Device struct {
	persistence.Record `mapstructure:",squash"`
	Brand              string `mapstructure:"brand" json:"brand"`
	Kind               string `mapstructure:"kind" json:"kind"`
	Model              string `mapstructure:"model" json:"model"`
	OwnerId            string `mapstructure:"owner_id" json:"owner_id"`
	Serial             string `mapstructure:"serial" json:"serial"`
}

func NewDevice(brand string, kind string, model string, ownerId string, serial string) (*Device, error) {
	if brand == "" || kind == "" || model == "" || ownerId == "" {
		return nil, errors.New("[Error]: Customer, Kind, Brand and Model must be set")
	} else {
		return &Device{
			Record:  persistence.Record{},
			Brand:   brand,
			Kind:    kind,
			Model:   model,
			OwnerId: ownerId,
			Serial:  serial,
		}, nil
	}
}
