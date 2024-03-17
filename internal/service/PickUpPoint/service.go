package PickUpPoint

import (
	"PickUpPoint/internal/model"
	"errors"
	"strconv"
	"time"
)

type stor interface {
	AcceptOrderFromCourier(int64, int64, time.Time) error
	ReturnToCourier(int64) error
	GiveOrderToClient([]int64) error
	ReturnOrderList(int64, string) ([]model.PickUpPoint, error)
	ReturnOrder(int64, int64) error
	ListReturn() ([]model.PickUpPoint, error)
}

type Service struct {
	s stor
}

func New(s stor) Service {
	return Service{s: s}
}

// AcceptOrderFromCourier checks data for validity and returns method from storage
func (s *Service) AcceptOrderFromCourier(OrderID int64, UserID int64, date time.Time) error {
	if OrderID < 0 {
		return errors.New("negative order id")
	}
	if UserID < 0 {
		return errors.New("negative user id")
	}
	return s.s.AcceptOrderFromCourier(OrderID, UserID, date)
}

// ReturnToCourier checks data for validity and returns method from storage
func (s *Service) ReturnToCourier(OrderID int64) error {
	if OrderID < 0 {
		return errors.New("negative order id")
	}
	return s.s.ReturnToCourier(OrderID)
}

// GiveOrderToClient checks data for validity and returns method from storage
func (s *Service) GiveOrderToClient(OrdersID []int64) error {
	if len(OrdersID) == 0 {
		return errors.New("empty list")
	}
	return s.s.GiveOrderToClient(OrdersID)
}

// ReturnOrderList checks data for validity and returns method from storage
func (s *Service) ReturnOrderList(UserID int64, n string) ([]model.PickUpPoint, error) {
	if UserID < 0 {
		return nil, errors.New("negative user id")
	}
	_, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}
	return s.s.ReturnOrderList(UserID, n)
}

// ReturnOrder checks data for validity and returns method from storage
func (s *Service) ReturnOrder(UserID int64, OrderID int64) error {
	if OrderID < 0 {
		return errors.New("negative order id")
	}
	if UserID < 0 {
		return errors.New("negative user id")
	}
	return s.s.ReturnOrder(UserID, OrderID)
}

// ListReturn returns method from storage
func (s *Service) ListReturn() ([]model.PickUpPoint, error) {
	return s.s.ListReturn()
}
