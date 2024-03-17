package PickUpPoint

import "time"

type PickUpPointDto struct {
	UserID     int64
	OrderID    int64
	IsGiven    bool
	IsReturned bool
	ExpireDate time.Time
}
