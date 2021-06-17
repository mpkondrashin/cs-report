package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	//	"log"
	"net/http"
	//	"os"
)

type (
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
	req.Header.Add("Content-Type", "application/json")
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
	url := fmt.Sprintf("%s/sessions/%s", s.smartCheck.url, s.response.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := s.Request(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(bodyBytes) == 0 {
		return nil
	}
	var response ResponseDeleteSession
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}
	if response.Message != "" {
		return fmt.Errorf("Delete: %s", response.Message)
	}
	return nil
}

type ListScansParameters struct {
	Expand     string
	Cursor     string
	Limit      int64
	Registry   string
	Repository string
	Tag        string
	Digest     string
	Exact      bool
	Status     string
}

func (l *ListScansParameters) Query() string {
	var sb strings.Builder
	if l.Expand != "" {
		sb.WriteString("expand=")
		sb.WriteString(l.Expand)
		sb.WriteString("&")
	}
	if l.Cursor != "" {
		sb.WriteString("cursor=")
		sb.WriteString(l.Cursor)
		sb.WriteString("&")
	}
	if l.Limit != 0 {
		sb.WriteString("limit=")
		sb.WriteString(strconv.FormatInt(l.Limit, 10))
		sb.WriteString("&")
	}
	if l.Registry != "" {
		sb.WriteString("registry=")
		sb.WriteString(l.Registry)
		sb.WriteString("&")
	}
	if l.Repository != "" {
		sb.WriteString("repository=")
		sb.WriteString(l.Repository)
		sb.WriteString("&")
	}
	if l.Tag != "" {
		sb.WriteString("tag=")
		sb.WriteString(l.Tag)
		sb.WriteString("&")
	}
	if l.Digest != "" {
		sb.WriteString("digest=")
		sb.WriteString(l.Digest)
		sb.WriteString("&")
	}
	if l.Exact {
		sb.WriteString("exact=true&")
	}
	if l.Status != "" {
		sb.WriteString("status=")
		sb.WriteString(l.Status)
		sb.WriteString("&")
	}
	result := sb.String()
	if len(result) == 0 {
		return ""
	}
	return result[:len(result)-1]
}

func (s *SmartCheckSession) ListScans(parameters *ListScansParameters) (*ResponseListScans, error) {
	url := fmt.Sprintf("%s/scans?limit=1", s.smartCheck.url)
	req, err := http.NewRequest("GET", url, nil)
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
	if len(bodyBytes) == 0 {
		return nil, fmt.Errorf("Empty response")
	}
	var response ResponseListScans
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *SmartCheckSession) List(method, baseURL, parameters, key string, body io.Reader) {
	url := fmt.Sprintf("%s?%s", baseURL, parameters)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
		//return nil, err
	}
	//go func() {
	for {
		resp, err := s.Request(req)
		if err != nil {
			panic(err)
			//return nil, err
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
			//	return nil, err
		}
		if len(bodyBytes) == 0 {
			panic(fmt.Errorf("Empty response"))
			//	return nil, fmt.Errorf("Empty response")
		}

		var response map[string]interface{}
		err = json.Unmarshal([]byte(bodyBytes), &response)
		list, ok := response[key].([]interface{})
		if !ok {
			panic(fmt.Errorf("%s is not a list", key))
		}
		for n, each := range list {
			fmt.Printf("%d\n%v\n\n\n", n, each)
		}
		cursor, ok := response["next"]
		if !ok {
			break
		}
		url = fmt.Sprintf("%s?cursor=%s", baseURL, cursor)
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			panic(err)
		}
	}
	//	return &response, nil

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
	//	r := RequestCreateSession{}
	//	r.User.UserID = "administrator"
	//	r.User.Password = "Zxcv7890!"
	fmt.Println("Create Session")
	session, err := sc.CreateSession(&request)
	if err != nil {
		fmt.Println(err)
		return
	}
	listScansParameters := ListScansParameters{
		Expand:     "",
		Cursor:     "",
		Limit:      1,
		Registry:   "",
		Repository: "",
		Tag:        "",
		Digest:     "",
		Exact:      false,
		Status:     "",
	}

	resp, err := session.ListScans(&listScansParameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp.Scans)
	fmt.Printf("%d\n", len(resp.Scans))
	s, _ := json.MarshalIndent(resp.Scans, "", "\t")
	fmt.Print(string(s))

	session.List("GET", "https://192.168.184.18:31616/api/scans", "", "scans", nil)

	fmt.Println("Delete Session")
	err = session.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done")
}
