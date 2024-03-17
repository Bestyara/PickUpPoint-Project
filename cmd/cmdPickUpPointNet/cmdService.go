package cmdPickUpPointNet

import (
	"PickUpPoint/cmd/cmdHelp"
	"PickUpPoint/internal/model"
	ServicePickUpPointNet "PickUpPoint/internal/service/PickUpPointNet"
	"context"
	"fmt"
	"log"
	"sync"
)

func PickUpPointNet(s ServicePickUpPointNet.Service) {
	var command string
	var ID int64
	var Name, Address, Contact string
	fmt.Println("Enter command (HelpNet - for some help):")
	ctx, cancel := context.WithCancel(context.Background())
	readchan := make(chan int64)
	writechan := make(chan model.PickUpPointNet)
	defer close(readchan)
	defer close(writechan)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go Read(s, ctx, readchan, &wg)
	go Write(s, ctx, writechan, &wg)
	for {
		fmt.Scan(&command)
		if command != "HelpNet" {
			switch command {
			case "Read":
				fmt.Println("Enter pick up point ID:")
				fmt.Scan(&ID)
				readchan <- ID
			case "Write":
				fmt.Println("Enter pick up point ID:")
				fmt.Scan(&ID)
				fmt.Println("Enter pick up point Name:")
				fmt.Scan(&Name)
				fmt.Println("Enter pick up point Address:")
				fmt.Scan(&Address)
				fmt.Println("Enter pick up point Contact:")
				fmt.Scan(&Contact)
				writechan <- model.PickUpPointNet{ID: ID, Name: Name, Address: Address, Contact: Contact}
			case "Exit":
				cancel()
				wg.Wait()
			default:
				fmt.Println("Incorrect command, try again!")
			}
		} else {
			cmdHelp.HelpNet()
		}
	}
}

// Read gets data output in some format
func Read(s ServicePickUpPointNet.Service, ctx context.Context, readchan <-chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case ID := <-readchan:
			data, err := s.Read(ID)
			if err != nil {
				log.Println(err)
			} else {
				returnFormatOutput(data)
			}
		}
	}
}

// Write puts data in storage file
func Write(s ServicePickUpPointNet.Service, ctx context.Context, writechan <-chan model.PickUpPointNet, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case pupnet := <-writechan:
			err := s.Write(pupnet.ID, pupnet.Name, pupnet.Address, pupnet.Contact)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Info about Pick Up Point have been successfully added")
			}
		}
	}
}
