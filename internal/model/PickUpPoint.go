package model

import "time"

type PickUpPoint struct {
	UserID     int64
	OrderID    int64
	IsGiven    bool
	IsReturned bool
	ExpireDate time.Time
}
