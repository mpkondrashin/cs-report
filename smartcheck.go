package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	//	"log"
	"github.com/deiu/linkheader"
	"net/http"
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
	url := fmt.Sprintf("%s/api/sessions", s.url)
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
	bodyBytes, err := io.ReadAll(resp.Body)
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
	url := fmt.Sprintf("%s/api/sessions/%s", s.smartCheck.url, s.response.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := s.Request(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
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
	bodyBytes, err := io.ReadAll(resp.Body)
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

func (s *SmartCheckSession) List(url, key string, body io.Reader) chan []byte {
	out := make(chan []byte, 100)
	go func() {
		uri := s.smartCheck.url + url
		req, err := http.NewRequest("GET", uri, body)
		if err != nil {
			panic(fmt.Errorf("NewRequest %s: %w", uri, err))
		}
		for {
			resp, err := s.Request(req)
			if err != nil {
				panic(fmt.Errorf("Request %s: %w", uri, err))
				//return nil, err
			}
			defer resp.Body.Close()
			bodyBytes, err := io.ReadAll(resp.Body)
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
			if err != nil {
				panic(fmt.Errorf("%s\n%w", string(bodyBytes), err))
			}
			list, ok := response[key].([]interface{})
			if !ok {
				panic(fmt.Errorf("%s\n%s is not a list",
					string(bodyBytes), key))
			}
			for _, each := range list {
				//fmt.Printf("\n\n%v\n\n", reflect.TypeOf(each))
				js, err := json.Marshal(each)
				if err != nil {
					panic(err)
				}
				out <- js
				//	fmt.Printf("%d\n%v\n\n\n", n, each)
			}

			linkHeader := resp.Header.Get("Link")
			//fmt.Println("link Header: ", linkHeader)
			if linkHeader == "" {
				break
			}
			linkMap := lh.ParseHeader(linkHeader)
			linkNext, ok := linkMap["next"]
			if !ok {
				break
			}
			linkHref, ok := linkNext["href"]
			if !ok {
				break
			}
			url = fmt.Sprintf("%s%s", s.smartCheck.url, linkHref)
			//fmt.Println("URL", url)
			req, err = http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}
		}
		close(out)
	}()
	return out
	//	return &response, nil

}

func (s *SmartCheckSession) ListRegistries() chan *ResponseRegistry {
	out := make(chan *ResponseRegistry, 100)
	go func() {
		regChan := s.List("/api/registries", "registries", nil)
		for reg := range regChan {
			var response ResponseRegistry
			err := json.Unmarshal(reg, &response)
			if err != nil {
				panic(err)
			}
			out <- &response
		}
		close(out)
	}()
	return out
}

func (s *SmartCheckSession) ListRegistryImages(registryId string) chan *ResponseImage {
	out := make(chan *ResponseImage, 100)
	go func() {
		path := fmt.Sprintf("/api/registries/%s/images", registryId)
		regChan := s.List(path, "images", nil)
		for reg := range regChan {
			var response ResponseImage
			err := json.Unmarshal(reg, &response)
			if err != nil {
				panic(err)
			}
			out <- &response
		}
		close(out)
	}()
	return out
}

func (s *SmartCheckSession) ImageLastScan(image *ResponseImage) *ResponseScan {
	query := fmt.Sprintf("/api/scans?limit=1&registry=%s&repository=%s&tag=%s&digest=%s&exact=true&",
		image.Registry, image.Repository, image.Tag, image.Digest)

	//sb.WriteString("status=")
	//fmt.Println("Query: ", query)
	//fmt.Printf("image: %+v", image)
	scanChan := s.List(query, "scans", nil)
	scan := <-scanChan
	var response ResponseScan
	err := json.Unmarshal(scan, &response)
	if err != nil {
		panic(err)
	}
	return &response
}

func (s *SmartCheckSession) ListMalwareFindings(query string) chan *ResponseLayerMalware {
	out := make(chan *ResponseLayerMalware, 100)
	go func() {
		responseChan := s.List(query, "malware", nil)
		for respJson := range responseChan {
			var response ResponseLayerMalware
			err := json.Unmarshal(respJson, &response)
			if err != nil {
				panic(err)
			}
			out <- &response
		}
		close(out)
	}()
	return out
}

func (s *SmartCheckSession) ListVulnerabilitiesFindings(query string) chan *ResponseLayerVulnerabilities {
	out := make(chan *ResponseLayerVulnerabilities, 100)
	go func() {
		responseChan := s.List(query, "vulnerabilities", nil)
		for respJson := range responseChan {
			var response ResponseLayerVulnerabilities
			err := json.Unmarshal(respJson, &response)
			if err != nil {
				panic(err)
			}
			out <- &response
		}
		close(out)
	}()
	return out
}

func (s *SmartCheckSession) ListContentsFindings(query string) chan *ResponseLayerContents {
	out := make(chan *ResponseLayerContents, 100)
	go func() {
		responseChan := s.List(query, "contents", nil)
		for respJson := range responseChan {
			var response ResponseLayerContents
			err := json.Unmarshal(respJson, &response)
			if err != nil {
				panic(err)
			}
			out <- &response
		}
		close(out)
	}()
	return out
}
