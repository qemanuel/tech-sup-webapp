package models

import (
	"errors"
	"time"
)

type Incidence struct {
	device    *Device
	timeStamp time.Time
	body      string
	next      *Incidence
	autor     *Worker
}

type IncidenceList struct {
	head   *Incidence
	Length int
}

func NewIncidenceList() *IncidenceList {
	return &IncidenceList{
		head:   nil,
		Length: 0,
	}
}

func (list *IncidenceList) AddIncidence(device *Device, body string, autor *Worker) (*Incidence, error) {
	if list == nil {
		return nil, errors.New("[Error]: IncidenceList missing")
	}
	if device == nil || autor == nil || body == "" {
		return nil, errors.New("[Error]: Device, Autor and Body must be set")
	}
	incidence := &Incidence{
		device:    device,
		timeStamp: time.Now(),
		body:      body,
		next:      nil,
		autor:     autor,
	}
	if list.head == nil {
		list.head = incidence
	} else {
		iterator := list.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = incidence
	}
	list.Length += 1
	return incidence, nil
}

func (list *IncidenceList) GetIncidence(find *Incidence) Incidence {
	var incidence Incidence
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		if iterator == find {
			incidence = Incidence{
				device:    iterator.device,
				timeStamp: iterator.timeStamp,
				body:      iterator.body,
				next:      iterator.next,
				autor:     iterator.autor,
			}
		}
	}
	return incidence
}

func (list *IncidenceList) RemoveIncidence(incidence *Incidence) {
	var previous *Incidence
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		if iterator == incidence {
			if list.head == iterator {
				list.head = iterator.next
			} else {
				previous.next = iterator.next
				return
			}
		}
		previous = iterator
	}
}

func (list *IncidenceList) GetAllIncidences() []Incidence {
	IncidenceSlice := make([]Incidence, 0, list.Length)
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		incidence := Incidence{
			device:    iterator.device,
			timeStamp: iterator.timeStamp,
			body:      iterator.body,
			next:      iterator.next,
			autor:     iterator.autor,
		}
		IncidenceSlice = append(IncidenceSlice, incidence)
	}
	return IncidenceSlice
}
