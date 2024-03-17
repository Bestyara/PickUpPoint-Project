package PickUpPoint

import (
	"PickUpPoint/internal/model"
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"time"
)

type Storage struct {
	filename string
	mp       []PickUpPointDto
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

// AcceptOrderFromCourier adds an order to storage
func (s *Storage) AcceptOrderFromCourier(OrderID int64, UserID int64, date time.Time) error {
	data := s.mp

	for _, v := range data {
		if OrderID == v.OrderID {
			return errors.New("order already has been received by PVZ")
		}
	}

	if !checkDate(date) {
		return errors.New("invalid date")
	}

	data = append(data, PickUpPointDto{OrderID: OrderID,
		UserID:     UserID,
		IsGiven:    false,
		IsReturned: false,
		ExpireDate: date})
	s.mp = data
	err := s.writeData(data)
	if err != nil {
		return err
	}
	return nil
}

// ReturnToCourier return the order to courier
func (s *Storage) ReturnToCourier(OrderID int64) error {
	data := s.mp
	count := 0
	for i, v := range data {
		if v.OrderID == OrderID {
			if checkDate(v.ExpireDate) && !v.IsGiven { //if the date is correct and recipient didn't achieve that order
				data = append(data[:i], data[i+1:]...)
				count++
			} else if v.IsGiven {
				return errors.New("order have been already achieved")
			} else if !checkDate(v.ExpireDate) {
				return errors.New("invalid date")
			}
			break
		}
	}
	if count == 0 {
		return errors.New("there is no order with such id")
	}

	s.mp = data
	err := s.writeData(data)
	if err != nil {
		return err
	}
	return nil
}

// GiveOrderToClient give an order to the client
func (s *Storage) GiveOrderToClient(OrdersID []int64) error {
	data := sliceToMap(s) //It is simplier to work with Map
	userid := data[OrdersID[0]].UserID

	for _, v := range OrdersID {
		if !(data[v].IsGiven == false && data[v].IsReturned == false && checkDate(data[v].ExpireDate) && data[v].UserID == userid) {
			return errors.New("order can not be given to client")
		}
	}
	for _, v := range OrdersID {
		data[v] = PickUpPointDto{OrderID: data[v].OrderID,
			UserID:     data[v].UserID,
			IsGiven:    true,
			IsReturned: data[v].IsReturned,
			ExpireDate: data[v].ExpireDate}
	}

	err := s.writeData(mapToSlice(data))
	if err != nil {
		return err
	}
	return nil
}

// ReturnOrderList gives all saved data
func (s *Storage) ReturnOrderList(UserID int64, n string) ([]model.PickUpPoint, error) {
	data := dtoToModel(s.mp)
	var userData []model.PickUpPoint
	for _, v := range data {
		if v.UserID == UserID {
			userData = append(userData, v)
		}
	}
	if len(userData) == 0 {
		return nil, errors.New("user not found")
	}
	if n == "all" {
		return userData, nil
	} else {
		nint, _ := strconv.Atoi(n)
		if nint > len(userData) {
			return nil, errors.New("slice is shorter than n")
		}
		return userData[len(userData)-nint:], nil
	}
}

// ReturnOrder deletes order with concrete id
func (s *Storage) ReturnOrder(UserID int64, OrderID int64) error {
	data := s.mp

	var dataid int = -1
	for i, v := range data {
		if v.UserID == UserID && v.OrderID == OrderID {
			dataid = i
		}
	}

	if dataid == -1 {
		return errors.New("there is no such an order or client is not in database")
	}
	if time.Now().After(data[dataid].ExpireDate) {
		return errors.New("invalid date of return")
	}
	data[dataid].IsReturned = true

	err := s.writeData(data)
	if err != nil {
		return err
	}
	return nil
}

// ListReturn return a list of returned orders
func (s *Storage) ListReturn() ([]model.PickUpPoint, error) { //completed
	data := dtoToModel(s.mp)
	var returnList []model.PickUpPoint
	for _, v := range data {
		if v.IsReturned {
			returnList = append(returnList, v)
		}
	}
	if len(returnList) == 0 {
		return nil, errors.New("the list of returned orders is empty")
	}
	return returnList, nil
}

// writeData writes data in file
func (s *Storage) writeData(data []PickUpPointDto) error {
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

// dtoToModel converts data from dto struct to model (because it could be more semantic right)
func dtoToModel(dto []PickUpPointDto) []model.PickUpPoint {
	var mod []model.PickUpPoint
	for _, v := range dto {
		mod = append(mod, model.PickUpPoint{
			OrderID:    v.OrderID,
			UserID:     v.UserID,
			IsGiven:    v.IsGiven,
			IsReturned: v.IsReturned,
			ExpireDate: v.ExpireDate,
		})
	}
	return mod
}

// checkDate checks if the expiry date is before the current date
func checkDate(date time.Time) bool {
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// Map gives us a map, because map suits better in some cases
func sliceToMap(s *Storage) map[int64]PickUpPointDto {
	m := make(map[int64]PickUpPointDto, len(s.mp))
	for _, v := range s.mp {
		m[v.OrderID] = v
	}
	return m
}

// MapToSlice converts map to slice
func mapToSlice(m map[int64]PickUpPointDto) []PickUpPointDto {
	var s []PickUpPointDto
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
