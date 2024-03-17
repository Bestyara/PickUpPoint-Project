package cmdPickUpPoint

import (
	"PickUpPoint/internal/model"
	"fmt"
)

func returnFormatList(data []model.PickUpPoint) {
	for i, v := range data {
		fmt.Printf("----------Record №%d----------\n", i+1)
		fmt.Println("User ID:", v.UserID)
		fmt.Println("OrderID:", v.OrderID)
		fmt.Println("ExpireDate:", v.ExpireDate)
		fmt.Println("Status of order:")
		if v.IsGiven {
			fmt.Println("User has achieved his order")
		} else {
			fmt.Println("User has not achieved his order")
		}
		if v.IsReturned {
			fmt.Println("User has returned his order")
		} else {
			fmt.Println("User has not returned his order")
		}
	}
	fmt.Println("------------------------------")
}
