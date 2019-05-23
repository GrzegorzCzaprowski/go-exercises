package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Area struct {
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	City      string  `json:"city"`
	Postal    string  `json:"zip"`
	Loc       float32 `json:"lat"`
	Lon       float32 `json:"lon"`
	Hostname  string  `json:"org"`
	IPaddress string  `json:"query"`
}

func main() {
	var flagIP string
	flag.StringVar(&flagIP, "ip", "", "area's IP")
	var flagGeo bool
	flag.BoolVar(&flagGeo, "geo", false, "only area's geographic data")
	flag.Parse()

	var url strings.Builder
	url.WriteString("http://IP-api.com/json/")
	url.WriteString(flagIP)
	url.WriteString("?fields=country,regionName,city,zip,lat,lon,org,query")

	resp, err := http.Get(url.String())
	if err != nil {
		log.Fatalln(err)
	}

	areaJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var area Area
	err = json.Unmarshal(areaJson, &area)
	if err != nil {
		log.Fatalln(err)
	}

	if flagGeo {
		fmt.Printf("City: %s\n", area.City)
		fmt.Printf("Region: %s\n", area.Region)
		fmt.Printf("Country: %s\n", area.Country)
		fmt.Printf("Loc: %f, %f\n", area.Loc, area.Lon)
	} else {
		fmt.Printf("IP address: %s\n", area.IPaddress)
		fmt.Printf("Organization: %s\n", area.Hostname)
		fmt.Printf("City: %s\n", area.City)
		fmt.Printf("Region: %s\n", area.Region)
		fmt.Printf("Country: %s\n", area.Country)
		fmt.Printf("Loc: %f, %f\n", area.Loc, area.Lon)
		fmt.Printf("Postal: %s\n", area.Postal)
	}
}
