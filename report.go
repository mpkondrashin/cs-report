package main

import (
	"html/template"
	"io"
	//"io/ioutil" // change to os as for go 1.16
	//"os"
	"time"
)

type (
	ReportGenerator struct {
		Template string
		Report   SmartCheckReport
	}
	SmartCheckReport struct {
		GeneratedOn time.Time
		Completed   time.Time
		Difference  time.Duration
		Registries  []RegistryReport
	}
	RegistryReport struct {
		Name   string
		Images []ImageReport
	}
	ImageReport struct {
		Name   string
		Layers []LayerReport
	}

	LayerReport struct {
		ID        string
		CreatedAt time.Time
		CreatedBy string
		Malware   []MalwareReport
		Package   []PackageReport
		Contents  []ContentsReport
	}
	MalwareReport struct {
		Filename string
		Name     string
		URL      string
	}
	PackageReport struct {
		Name            string
		Version         string
		Vulnerabilities []VulnerabilityReport
	}
	VulnerabilityReport struct {
		Name     string
		Link     string
		Severity string
	}
	ContentsReport struct {
		Severity    string
		Rule        string
		Description string
		Filename    string
	}
)

func NewReportGenerator(template string) *ReportGenerator {
	return &ReportGenerator{
		Template: template,
	}
}

func (r *ReportGenerator) Generate(w io.Writer) error {
	tmpl, err := template.New("report").Parse(r.Template)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, &r.Report)
}

/*
func main() {
	template, err := ioutil.ReadFile("./report_template.html")
	if err != nil {
		panic(err)
	}
	r := NewReportGenerator(string(template))
	m := MalwareReport{
		"eicar.com",
		"EICAR_TEST_FILE",
		"http://enc/eicar.com",
	}
	vr := VulnerabilityReport{
		Name:     "CVE-2020-123",
		Link:     "http://mitre.org",
		Severity: "Hight",
	}

	pr := PackageReport{
		Name:            "nginx",
		Version:         "0.9",
		Vulnerabilities: []VulnerabilityReport{vr},
	}
	cr := ContentsReport{
		Severity:    "Low",
		Rule:        "pwd_file",
		Description: "very ver bad",
		Filename:    "/etc/some_file",
	}
	l := LayerReport{
		ID:        "IDIDIDIDI",
		CreatedAt: time.Now(),
		CreatedBy: "Some command",
		Malware:   []MalwareReport{m},
		Package:   []PackageReport{pr},
		Contents:  []ContentsReport{cr},
	}
	i := ImageReport{
		Name:   "image name",
		Layers: []LayerReport{l},
	}
	rr := RegistryReport{
		Name:   "Registry name",
		Images: []ImageReport{i},
	}
	scr := SmartCheckReport{
		GeneratedOn: time.Now(),
		Completed:   time.Now(),
		Registries:  []RegistryReport{rr},
	}
	r.Report = scr
	filename := "report.html"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = r.Generate(f)
	if err != nil {
		panic(err)
	}
}
*/
