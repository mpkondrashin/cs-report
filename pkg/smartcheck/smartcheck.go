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
		PasswordChangeRequired bool
		Created                string
		Updated                string
	}

	ResponseCreateSession struct {
		Id              string
		Href            string
		User            ResponseUser
		Token           string
		Created         string
		Updated         string
		Expires         string
		RoleSessionName string
	}
	//ResponseDeleteSessionFields {

	//"fields": {
	//"property1": "string",
	//"property2": "string"
	//}
	//}
	ResponseDeleteSession struct {
		Message string
	}
)

type SmartCheck struct {
	url            string
	ignoreTLSError bool
}

func NewSmartCheck(url string, ignoreTLSError bool) *SmartCheck {
	return &SmartCheck{
		url:            url,
		ignoreTLSError: ignoreTLSError,
	}
}

func (s *SmartCheck) Request(req *http.Request) (*http.Response, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: s.ignoreTLSError},
	}
	client := &http.Client{Transport: transport}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Version", "2018-05-01")
	return client.Do(req)
}

func (s *SmartCheck) CreateSession(credentials interface{}) (*SmartCheckSession, error) {
	requestJSON, err := json.Marshal(credentials)
	if err != nil {
		return nil, err

	}
	body := bytes.NewBuffer(requestJSON)
	url := fmt.Sprintf("%s/sessions", s.url)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := s.Request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var session SmartCheckSession
	session.smartCheck = s
	err = json.Unmarshal(bodyBytes, &session.response)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

type SmartCheckSession struct {
	smartCheck *SmartCheck
	response   ResponseCreateSession
}

func (s *SmartCheckSession) Request(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.response.Token))
	return s.smartCheck.Request(req)
}
func (s *SmartCheckSession) Delete() error {
	fmt.Print(s.smartCheck.url, s.response.Id)

	url := fmt.Sprintf("%s/sessions/%s", s.smartCheck.url, s.response.Id)
	req, err := http.NewRequest("DELETE", url, nil)
	fmt.Println("A")
	if err != nil {
		return err
	}
	resp, err := s.Request(req)
	fmt.Println("A")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println("A")
	if err != nil {
		return err
	}
	if len(bodyBytes) == 0 {
		return nil
	}
	var response ResponseDeleteSession
	err = json.Unmarshal(bodyBytes, &response)
	fmt.Println("A")
	if err != nil {
		return err
	}
	if response.Message != "" {
		return fmt.Errorf("Delete: %s", response.Message)
	}
	fmt.Println("A")
	return nil
}

func main() {
	URL := "https://192.168.184.18:31616/api"
	sc := NewSmartCheck(URL, true)
	request := RequestCreateSessionUser{
		User: RequestCreateSessionUserCredentials{
			UserID:   "administrator",
			Password: "Zxcv7890!",
		},
	}
	fmt.Println("Create Session")
	session, err := sc.CreateSession(&request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete Session")
	err = session.Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(err)
}
