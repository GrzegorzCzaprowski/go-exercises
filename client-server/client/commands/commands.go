package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

type Client int

func (c Client) DeleteUser(id uint64, address string) {
	url := fmt.Sprintf("%s/users/%d/", address, id)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url, nil)

	response, _ := client.Do(req)
	defer response.Body.Close()
}

func (c Client) PostUser(user User, address string) {
	var url string
	url = fmt.Sprintf("%s/users/", address)

	jsonValue, _ := json.Marshal(user)
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

func (c Client) GetUser(id uint64, address string) {
	url := fmt.Sprintf("%s/users/%d/", address, id)

	response, _ := http.Get(url)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
