package models

import (
	"errors"
)

type Device struct {
	brand  string
	kind   string
	model  string
	owner  *Customer
	serial string
	id     int
}

func NewDevice(brand string, kind string, model string, owner *Customer, serial string) (*Device, error) {
	if brand == "" || kind == "" || model == "" || owner == nil {
		return nil, errors.New("[Error]: Customer, Kind, Brand and Model must be set")
	} else {
		return &Device{
			brand:  brand,
			kind:   kind,
			model:  model,
			owner:  owner,
			serial: serial,
		}, nil
	}
}

func GetDevice(dev *Device) Device {
	device := Device{
		brand:  dev.brand,
		kind:   dev.kind,
		model:  dev.model,
		owner:  dev.owner,
		serial: dev.serial,
	}
	return device
}

func (dev *Device) SetId(id int) error {
	if id != 0 {
		return errors.New("[Error]: ID already assigned")
	} else {
		dev.id = id
	}
	return nil
}

func (dev *Device) GetId() int {
	return dev.id
}

func (dev *Device) String() []string {
	return []string{
		dev.brand,
		dev.kind,
		dev.model,
		dev.owner.name,
		dev.serial,
	}
}

func (dev *Device) Update(brand string, kind string, model string, serial string) error {
	var err error
	if kind == "" && brand == "" && model == "" && serial == "" {
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
	if serial != "" {
		dev.serial = serial
	}
	return err
}
