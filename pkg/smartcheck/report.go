package main

import (
	"html/template"
	"os"
	"time"
)

type (
	SmartCheckReport struct {
		GeneratedOn time.Time
		Registry    []RegistryReport
	}
	RegistryReport struct {
		Name   string
		Images []ImageReport
	}
	ScanReport struct {
		Completed  time.Time
		Registries []RegistryReport
	}

	ImageReport struct {
		Name  string
		Layer []LayerReport
	}

	LayerReport struct {
		ID            string
		CreatedAt     time.Time
		CreatedBy     string
		Malware       []MalwareReport
		Vulnerability []VulnerabilityReport
		Contents      []ContentsReport
	}
	MalwareReport struct {
		Filename string
		Name     string
		URL      string
	}
	VulnerabilityReport struct {
	}
	ContentsReport struct {
	}
)

func main() {
	m := MalwareReport{
		"eicar.com",
		"EICAR_TEST_FILE",
		"http://enc/eicar.com",
	}
	l := LayerReport{
		ID:        "IDIDIDIDI",
		CreatedAt: time.Now(),
		CreatedBy: "Some command",
		Malware:   []MalwareReport{m},
	}
	tmpl, _ := template.ParseFiles("./report_template.html")
	f, err := os.Create("report.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, &l)
	if err != nil {
		panic(err)
	}

}
