package main

import (
	"PickUpPoint/cmd/cmdHelp"
	"PickUpPoint/cmd/cmdPickUpPoint"
	"PickUpPoint/cmd/cmdPickUpPointNet"
	ServicePickUpPoint "PickUpPoint/internal/service/PickUpPoint"
	ServicePickUpPointNet "PickUpPoint/internal/service/PickUpPointNet"
	StoragePickUpPoint "PickUpPoint/internal/storage/PickUpPoint"
	StoragePickUpPointNet "PickUpPoint/internal/storage/PickUpPointNet"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter command (Help - for some help):")
	var command string
	fmt.Scan(&command)
	storpup, err := StoragePickUpPoint.New("orderStorage")
	if err != nil {
		log.Println("Can not connect to pick up point storage!", err)
		return
	}
	servpup := ServicePickUpPoint.New(&storpup)

	storpupnet, err := StoragePickUpPointNet.New("orderStorageNet")
	if err != nil {
		log.Println("Can not connect to pick up point net storage!", err)
		return
	}
	servpupnet := ServicePickUpPointNet.New(&storpupnet)

	if command != "Help" {
		switch command {
		case "AcceptOrderFromCourier":
			cmdPickUpPoint.AcceptOrderFromCourier(servpup)
		case "ReturnToCourier":
			cmdPickUpPoint.ReturnToCourier(servpup)
		case "GiveOrderToClient":
			cmdPickUpPoint.GiveOrderToClient(servpup)
		case "ReturnOrderList":
			cmdPickUpPoint.ReturnOrderList(servpup)
		case "ReturnOrder":
			cmdPickUpPoint.ReturnOrder(servpup)
		case "ListReturn":
			cmdPickUpPoint.ListReturn(servpup)
		case "PickUpPointNet":
			cmdPickUpPointNet.PickUpPointNet(servpupnet)
		default:
			fmt.Println("Incorrect command, try again!")
		}
	} else {
		cmdHelp.Help()
	}
}
