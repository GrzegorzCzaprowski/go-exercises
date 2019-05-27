package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
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
	var url strings.Builder
	url.WriteString(address)
	url.WriteString("/users/")

	jsonValue, _ := json.Marshal(user) //zamiana danych na forma jsona
	//zbuforowanie jasona do bytów
	response, err := http.Post(url.String(), "aplication/json", bytes.NewBuffer(jsonValue))
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
	var url strings.Builder
	url.WriteString(address)
	url.WriteString("/users/")
	url.WriteString(strconv.FormatUint(id, 10))
	url.WriteString("/")
	println(url.String())

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url.String(), nil)

	response, _ := client.Do(req)
	defer response.Body.Close()
}

func getUser(id uint64, address string) {
	var url strings.Builder
	url.WriteString(address)
	url.WriteString("/users/")
	url.WriteString(strconv.FormatUint(id, 10))
	url.WriteString("/")

	response, _ := http.Get(url.String())

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
		println("choose action")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		command := strings.Split(input, " ")
		if command[0] == "adduser" {
			var user User

			print("enter ID (uint64): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			i, _ := strconv.ParseUint(input, 0, 64)
			user.ID = i
			println(user.ID)

			print("enter Name (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Name = input
			println(user.Name)

			print("enter Surname (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Surname = input
			println(user.Surname)

			print("enter Email (string): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			user.Email = input
			println(user.Email)

			postUser(user, flagServerAddress)
		} else if command[0] == "delete" {
			userID, _ := strconv.ParseUint(command[1], 10, 64)
			deleteUser(userID, flagServerAddress)
		} else if command[0] == "getuser" {
			println("weszlo")
			userID, _ := strconv.ParseUint(command[1], 10, 64)
			getUser(userID, flagServerAddress)
		} else {
			println("bad command")
		}
	}
}
