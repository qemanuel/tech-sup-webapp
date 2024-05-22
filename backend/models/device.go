package models

import (
	"errors"
)

type Device struct {
	Brand   string `mapstructure:"brand" json:"brand"`
	Kind    string `mapstructure:"kind" json:"kind"`
	Model   string `mapstructure:"model" json:"model"`
	OwnerId string `mapstructure:"owner_id" json:"owner_id"`
	Serial  string `mapstructure:"serial" json:"serial"`
	Id      string `mapstructure:"id" json:"id"`
}

func NewDevice(brand string, kind string, model string, owner *Customer, serial string) (*Device, error) {
	if brand == "" || kind == "" || model == "" || owner == nil {
		return nil, errors.New("[Error]: Customer, Kind, Brand and Model must be set")
	} else {
		return &Device{
			Brand:   brand,
			Kind:    kind,
			Model:   model,
			OwnerId: owner.Id,
			Serial:  serial,
		}, nil
	}
}

func GetDevice(dev *Device) Device {
	device := Device{
		Brand:   dev.Brand,
		Kind:    dev.Kind,
		Model:   dev.Model,
		OwnerId: dev.OwnerId,
		Serial:  dev.Serial,
	}
	return device
}

func (dev *Device) SetId(id string) error {
	if dev.Id != "" {
		return errors.New("[Error]: ID already assigned")
	} else {
		dev.Id = id
		return nil
	}
}

func (dev *Device) GetId() string {
	return dev.Id
}

func (dev *Device) String() []string {
	return []string{
		dev.Brand,
		dev.Kind,
		dev.Model,
		dev.OwnerId,
		dev.Serial,
	}
}

func (dev *Device) Update(brand string, kind string, model string, serial string) error {
	var err error
	if kind == "" && brand == "" && model == "" && serial == "" {
		err = errors.New("[Error]: Specs are not set")
	}
	if kind != "" {
		dev.Kind = kind
	}
	if brand != "" {
		dev.Brand = brand
	}
	if model != "" {
		dev.Model = model
	}
	if serial != "" {
		dev.Serial = serial
	}
	return err
}
