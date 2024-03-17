package cmdPickUpPointNet

import (
	"PickUpPoint/internal/model"
	"fmt"
)

func returnFormatOutput(data model.PickUpPointNet) {
	fmt.Println("----------Pick-Up-Point----------")
	fmt.Println("ID:", data.ID)
	fmt.Println("Name:", data.Name)
	fmt.Println("Address:", data.Address)
	fmt.Println("Contacts:", data.Contact)
	fmt.Println("------------------------------")
}
