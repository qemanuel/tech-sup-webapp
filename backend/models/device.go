package models

import (
	"errors"
)

type Device struct {
	owner  *Customer
	serial int
	kind   string
	brand  string
	model  string
	id     int
}

func NewDevice(serial int, owner *Customer, kind string, brand string, model string) (*Device, error) {
	if kind == "" || brand == "" || model == "" || owner == nil {
		return nil, errors.New("[Error]: Customer, Kind, Brand and Model must be set")
	} else {
		return &Device{
			owner:  owner,
			kind:   kind,
			serial: serial,
			brand:  brand,
			model:  model,
		}, nil
	}
}

func GetDevice(dev *Device) Device {
	device := Device{
		owner:  dev.owner,
		kind:   dev.kind,
		serial: dev.serial,
		brand:  dev.brand,
		model:  dev.model,
	}
	return device
}

func (dev *Device) UpdateDeviceSpec(serial int, kind string, brand string, model string) error {
	var err error
	if kind == "" && brand == "" && model == "" && serial == 0 {
		err = errors.New("[Error]: Specs are not set")
	}
	if kind != "" {
		dev.kind = kind
	}
	if brand != "" {
		dev.brand = brand
	}
	if model != "" {
		dev.model = model
	}
	if serial != 0 {
		dev.serial = serial
	}
	return err
}
