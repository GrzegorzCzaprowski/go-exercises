package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

func postUser(user User, address string) {
	var url string
	url = fmt.Sprintf("%s/users/", address)

	jsonValue, _ := json.Marshal(user) //zamiana danych na forma jsona
	//zbuforowanie jasona do bytów
	response, err := http.Post(url, "aplication/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func deleteUser(id uint64, address string) {
	url := fmt.Sprintf("%s/users/%d/", address, id)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url, nil)

	response, _ := client.Do(req)
	defer response.Body.Close()
}

func getUser(id uint64, address string) {
	url := fmt.Sprintf("%s/users/%d/", address, id)

	response, _ := http.Get(url)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func main() {
	var flagServerAddress string
	flag.StringVar(&flagServerAddress, "addr", "http://localhost:8000", "server address")
	flag.Parse()

	for {
		fmt.Println("choose action")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		command := strings.Split(input, " ")
		if command[0] == "adduser" {
			var user User

			fmt.Print("enter ID (uint64): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			i, _ := strconv.ParseUint(input, 0, 64)
			user.ID = i
			fmt.Println(user.ID)

			fmt.Print("enter Name (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Name = input
			fmt.Println(user.Name)

			fmt.Print("enter Surname (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Surname = input
			fmt.Println(user.Surname)

			fmt.Print("enter Email (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Email = input
			fmt.Println(user.Email)

			postUser(user, flagServerAddress)
		} else if command[0] == "delete" {
			userID, _ := strconv.ParseUint(command[1], 10, 64)
			deleteUser(userID, flagServerAddress)
		} else if command[0] == "getuser" {
			userID, _ := strconv.ParseUint(command[1], 10, 64)
			getUser(userID, flagServerAddress)
		} else {
			fmt.Println("bad command")
		}
	}
}
