package PickUpPointNet

import (
	"PickUpPoint/internal/model"
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

type Storage struct {
	filename string
	mp       []PickUpPointNetDto
	mx       sync.RWMutex
}

// New initialize Storage object
func New(filename string) (Storage, error) {
	f, err := os.OpenFile(filename, os.O_CREATE, 0777)
	if err != nil {
		return Storage{}, err
	}

	var s Storage
	reader := bufio.NewReader(f)
	data, err := io.ReadAll(reader)
	if err != nil {
		return Storage{}, err
	}
	json.Unmarshal(data, &s.mp)
	if err != nil {
		return Storage{}, err
	}

	return Storage{filename: filename, mp: s.mp}, nil
}

// Read gives a data
func (s *Storage) Read(ID int64) (model.PickUpPointNet, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for _, v := range s.mp {
		if v.ID == ID {
			return dtoToModel([]PickUpPointNetDto{v})[0], nil
		}
	}
	return model.PickUpPointNet{}, errors.New("there is no pick up point with such ID")
}

// Write writes some data in file
func (s *Storage) Write(ID int64, Name string, Address string, Contact string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	data := s.mp
	for _, v := range data {
		if v.ID == ID {
			return errors.New("there is already a pick up point with same id")
		}
	}
	data = append(data, PickUpPointNetDto{
		ID:      ID,
		Name:    Name,
		Address: Address,
		Contact: Contact,
	})
	s.mp = data
	err := s.writeData(data)
	if err != nil {
		return err
	}
	return nil
}

// dtoToModel converts data from dto struct to model (because it could be more semantic right)
func dtoToModel(dto []PickUpPointNetDto) []model.PickUpPointNet {
	var mod []model.PickUpPointNet
	for _, v := range dto {
		mod = append(mod, model.PickUpPointNet{
			ID:      v.ID,
			Name:    v.Name,
			Address: v.Address,
			Contact: v.Contact,
		})
	}
	return mod
}

// writeData writes data in file
func (s *Storage) writeData(data []PickUpPointNetDto) error {
	filedata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.filename, filedata, 0777)
	if err != nil {
		return err
	}
	return nil
}
