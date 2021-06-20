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
	//	"reflect"
	//"github.com/swhite24/link"
	"github.com/deiu/linkheader"
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

/*
func (s *SmartCheckSession) List(method, baseURL, parameters, key string, body io.Reader) chan []byte {
	out := make(chan []byte, 100)
	go func() {
		url := fmt.Sprintf("%s/%s?%s", s.smartCheck.url, baseURL, parameters)
		//fmt.Println(url)
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			panic(err)
			//return nil, err
		}
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
			if err != nil {
				panic(err)
			}
			//fmt.Printf("\n\n%v\n\n", response)
			list, ok := response[key].([]interface{})
			if !ok {
				//fmt.Printf(string(bodyBytes))
				panic(fmt.Errorf("%s is not a list", key))
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
			cursor, ok := response["next"]
			if !ok {
				//fmt.Println("======= NO NEXT ======")
				break
			}
			url = fmt.Sprintf("%s/%s?cursor=%s", s.smartCheck.url, baseURL, cursor)
			req, err = http.NewRequest(method, url, nil)
			if err != nil {
				panic(err)
			}
		}
		close(out)
	}()
	return out
	//	return &response, nil

}
*/
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

func main() {
	URL := "https://192.168.184.18:31616"
	sc := NewSmartCheck(URL, true)
	request := RequestCreateSessionUser{
		User: RequestCreateSessionUserCredentials{
			UserID:   "administrator",
			Password: "Zxcv7890!",
		},
	}
	//fmt.Println("Create Session")
	session, err := sc.CreateSession(&request)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = session.Delete()
		if err != nil {
			panic(err)
		}
	}()
	for r := range session.ListRegistries() {
		fmt.Println("Registry:", r.ID)
		for im := range session.ListRegistryImages(r.ID) {
			fmt.Println("Image:", im.ID) //, im.Tag, im.Registry, im.Repository, im.Status)
			scan := session.ImageLastScan(im)
			for _, layer := range scan.Details.Results {
				//fmt.Println("Result:")
				if layer.Malware != "" {
					for malware := range session.ListMalwareFindings(layer.Malware) {
						name := malware.Icrc.Name
						url := malware.Icrc.URL
						if malware.Trendx.Found.Name != "" {
							name = malware.Trendx.Found.Name
							url = malware.Trendx.Found.URL
						}
						fmt.Println(scan.Details.Completed)
						fmt.Println(scan.Name)
						fmt.Printf("%s %s (%s)\n", malware.Filename, name, url)
						/*
							JSON, err := json.MarshalIndent(malware, "", "  ")
							if err != nil {
								panic(err)
							}
							fmt.Printf("Malware:\n%s\n", string(JSON))
						*/
					}
					continue
					for malware := range session.ListMalwareFindings(layer.Malware) {
						name, value := StructCSV(malware)
						fmt.Println("Malware N: ", name)
						fmt.Println("Malware V: ", value)
					}
				}
				if layer.Vulnerabilities != "" {
					continue
					for vulnerability := range session.ListVulnerabilitiesFindings(layer.Vulnerabilities) {
						name, value := StructCSV(vulnerability)
						fmt.Println("Vulnerability: N", name)
						fmt.Println("Vulnerability: V", value)
					}

				}
				if layer.Contents != "" {
					continue
					for contents := range session.ListContentsFindings(layer.Contents) {
						name, value := StructCSV(contents)

						fmt.Println("Contents: N", name)
						fmt.Println("Contents: V", value)
					}
				}
			}

		}
	}

	//err = session.Delete()
	//if err != nil {
	//	panic(err)
	//}
	/*fmt.Println("Delete Session")
	err = session.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}*/
	//fmt.Println("Done")
}
