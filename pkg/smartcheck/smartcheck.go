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

type (
	RequestCreateSessionUserCredentials struct {
		UserID   string
		Password string
	}
	RequestCreateSessionUser struct {
		User RequestCreateSessionUserCredentials
	}

	RequestCreateSessionSamlCredentials struct {
		Response     string
		SelectedRole string
	}

	RequestCreateSessionSaml struct {
		Saml RequestCreateSessionSamlCredentials
	}

	ResponseUser struct {
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

	ResponseSessions struct {
		Id              string
		Href            string
		User            ResponseUser
		Token           string
		Created         string
		Updated         string
		Expires         string
		RoleSessionName string
	}
)

func main() {
	URL := "https://192.168.184.18:31616/api/sessions"
	fmt.Println("Calling API...")

	request := RequestCreateSessionUser{
		User: RequestCreateSessionUserCredentials{
			UserID:   "administrator",
			Password: "Zxcv7890!",
		},
	}
	requestJSON, err := json.Marshal(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(requestJSON))
	ignoreTLSError := true
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: ignoreTLSError},
	}
	client := &http.Client{Transport: transport}
	body := bytes.NewBuffer(requestJSON)
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
	//	fmt.Print(string(bodyBytes))
	var sessions ResponseSessions
	json.Unmarshal(bodyBytes, &sessions)
	fmt.Printf("token: %s\n", sessions.Token)

}
