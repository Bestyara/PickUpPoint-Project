package PickUpPointNet

import (
	"PickUpPoint/internal/model"
	"errors"
)

type stor interface {
	Read(ID int64) (model.PickUpPointNet, error)
	Write(ID int64, Name string, Address string, Contact string) error
}

type Service struct {
	s stor
}

func New(s stor) Service {
	return Service{s: s}
}

// Read checks data for validity and returns method from storage
func (s *Service) Read(ID int64) (model.PickUpPointNet, error) {
	if ID < 0 {
		return model.PickUpPointNet{}, errors.New("negative user id")
	}
	return s.s.Read(ID)
}

// Write checks data for validity and returns method from storage
func (s *Service) Write(ID int64, Name string, Address string, Contact string) error {
	if ID < 0 {
		return errors.New("negative user id")
	}
	if Name == "" {
		return errors.New("empty name field")
	}
	if Address == "" {
		return errors.New("empty address field")
	}
	if Contact == "" {
		return errors.New("empty contact field")
	}
	return s.s.Write(ID, Name, Address, Contact)
}
