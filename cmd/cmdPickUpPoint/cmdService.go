package cmdPickUpPoint

import (
	"PickUpPoint/internal/service/PickUpPoint"
	"fmt"
	"log"
	"strconv"
	"time"
)

// AcceptOrderFromCourier gets input and adds new order data
func AcceptOrderFromCourier(serv PickUpPoint.Service) {
	var UserID, OrderID int64
	var ExpireDate time.Time
	fmt.Println("Enter User ID:")
	fmt.Scan(&UserID)
	fmt.Println("Enter Order ID:")
	fmt.Scan(&OrderID)
	fmt.Println("Enter expire date (yyyy-mm-dd):")
	var strdate string
	fmt.Scan(&strdate)
	ExpireDate, err := time.Parse("2006-01-02", strdate)
	if err != nil {
		log.Println("Can not parse a date", err)
	}
	err = serv.AcceptOrderFromCourier(OrderID, UserID, ExpireDate)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Order has been successfully created!")
	}
}

// ReturnToCourier gets data and delete it from list
func ReturnToCourier(serv PickUpPoint.Service) {
	var OrderID int64
	fmt.Println("Enter Order ID:")
	fmt.Scan(&OrderID)
	err := serv.ReturnToCourier(OrderID)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Order has been successfully returned to courier!")
	}
}

// GiveOrderToClient gives order to client (change give flag)
func GiveOrderToClient(serv PickUpPoint.Service) {
	//var OrderID int
	var s string
	var OrdersID []int64
	fmt.Println("Enter Order ID to receive: (exit - end input):")
	for {
		fmt.Scan(&s)
		if s == "exit" {
			break
		} else {
			OrderID, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Println("Can not convert string to int", err)
				return
			}
			OrdersID = append(OrdersID, OrderID)
		}
	}
	err := serv.GiveOrderToClient(OrdersID)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Orders has been successfully given to the client!")
	}
}

// ReturnOrderList gets user id and returns all his returns
func ReturnOrderList(serv PickUpPoint.Service) {
	var n string
	var UserID int64
	fmt.Println("Enter User ID:")
	fmt.Scan(&UserID)
	fmt.Println("Enter how many last order you need to print (all - all orders):")
	fmt.Scan(&n)
	data, err := serv.ReturnOrderList(UserID, n)
	if err != nil {
		log.Println(err)
	} else {
		returnFormatList(data)
	}
}

// ReturnOrder gets user id and order id and change flag of such user and order to true
func ReturnOrder(serv PickUpPoint.Service) {
	var UserID, OrderID int64
	fmt.Println("Enter User ID:")
	fmt.Scan(&UserID)
	fmt.Println("Enter Order ID:")
	fmt.Scan(&OrderID)
	err := serv.ReturnOrder(UserID, OrderID)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Order has been successfully returned to PVZ!")
	}
}

// ListReturn return all data
func ListReturn(serv PickUpPoint.Service) {
	data, err := serv.ListReturn()
	if err != nil {
		log.Println(err)
	} else {
		returnFormatList(data)
	}
}
