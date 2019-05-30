package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/GrzegorzCzaprowski/go-exercises/client-server/client/commands"
)

func adduser() commands.User {
	var user commands.User

	fmt.Print("enter ID (uint64): ")
	var id string
	fmt.Scanln(&id)
	user.ID, _ = strconv.ParseUint(id, 0, 64)

	fmt.Print("enter Name (string): ")
	var name string
	fmt.Scanln(&name)
	user.Name = name

	fmt.Print("enter Surname (string): ")
	var surname string
	fmt.Scanln(&surname)
	user.Surname = surname

	fmt.Print("enter Email (string): ")
	var email string
	fmt.Scanln(&email)
	user.Email = email

	return user
}

func choose(input1, input2, flagServerAddress string) {
	var client commands.Client

	client.Address = flagServerAddress
	client.ID, _ = strconv.ParseUint(input2, 10, 64)

	if input1 == "adduser" {
		client.User = adduser()
		client.PostUser()
	} else if input1 == "delete" {
		client.DeleteUser()
	} else if input1 == "getuser" {
		client.GetUser()
	} else {
		fmt.Println("bad command")
	}
}

func main() {
	var flagServerAddress string
	flag.StringVar(&flagServerAddress, "addr", "http://localhost:8000", "server address")
	flag.Parse()

	for {
		fmt.Println("choose action")
		var input1 string
		var input2 string
		fmt.Scanln(&input1, &input2)

		choose(input1, input2, flagServerAddress)
	}
}
