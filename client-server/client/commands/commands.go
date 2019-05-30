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

type Client struct {
	Address string
	ID      uint64
	User    User
}

func (c Client) DeleteUser() {
	url := fmt.Sprintf("%s/users/%d/", c.Address, c.ID)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url, nil)

	response, _ := client.Do(req)
	defer response.Body.Close()
}

func (c Client) PostUser() {
	var url string
	url = fmt.Sprintf("%s/users/", c.Address)

	jsonValue, _ := json.Marshal(c.User)
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

func (c Client) GetUser() {
	url := fmt.Sprintf("%s/users/%d/", c.Address, c.ID)

	response, _ := http.Get(url)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
