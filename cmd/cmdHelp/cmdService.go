package cmdHelp

import "fmt"

func Help() {
	fmt.Printf(`-------------PickUpPoint-usage-guide----------------
AcceptOrderFromCourier       Accepts an order from courier
ReturnToCourier              Returns an order to courier
GiveOrderToClient            Gives an order to recepient
ReturnOrderList              Prints n last orders of user
ReturnOrder                  Returns an order to pick up point
ListReturn                   Prints list of returns
PickUpPointNet	             Enter net mode (configure info about the PickUpPoints)`)
}

func HelpNet() {
	fmt.Printf(`-------------PickUpPointNet-usage-guide----------------
Read						Read info about pick up point
Write						Write info about pick up point
Exit                        End entering commands
`)
}
