package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {
	fmt.Println(numAstronauts())
}

func computeHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return strings.ToUpper(fmt.Sprintf("%x",hash.Sum(nil)))
}

func numAstronauts() int {
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return people1.Number
}