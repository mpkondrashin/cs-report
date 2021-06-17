package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//	"log"
	"net/http"
	//	"os"
)

type ResponseUser struct {
	Id                     string
	Href                   string
	UserID                 string
	Name                   string
	Description            string
	Role                   string
	PasswordChangeRequired string
	Created                string
	Updated                string
}

type ResponseSessions struct {
	Id              string
	Href            string
	User            ResponseUser
	Token           string
	Created         string
	Updated         string
	Expires         string
	RoleSessionName string
}

func main() {
	URL := "https://192.168.184.18:31616/api/sessions"
	fmt.Println("Calling API...")

	requestBody := `{
  "user": {
    "userID": "administrator",
    "password": "Zxcv7890!"
  }
}`
	ignoreTLSError := true
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: ignoreTLSError},
	}
	client := &http.Client{Transport: transport}
	body := bytes.NewBufferString(requestBody)
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(bodyBytes))
	var sessions ResponseSessions
	json.Unmarshal(bodyBytes, &sessions)
	fmt.Printf("token: %s\n", sessions.Token)

}
