package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"time"
)

//go:embed report_template.html
var f embed.FS

var Conf struct {
	URL             string
	UserID          string
	Password        string
	IgnoreTLSErrors bool
}

func Configure() {
	flag.StringVar(&Conf.URL, "url", "", "Smart Check/Container Security URL")
	flag.StringVar(&Conf.UserID, "user", "administrator", "User name")
	flag.StringVar(&Conf.Password, "password", "", "Password")
	flag.BoolVar(&Conf.IgnoreTLSErrors, "ignore_tls", false, "Ignore TLS Errors")
	flag.Parse()
	if Conf.URL == "" {
		fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", "url")
		os.Exit(2)
	}
	if Conf.Password == "" {
		fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", "password")
		os.Exit(2)
	}
	fmt.Println(Conf)
}

func main() {
	Configure()
	//URL := "https://192.168.184.18:31616"
	sc := NewSmartCheck(Conf.URL, Conf.IgnoreTLSErrors)
	request := RequestCreateSessionUser{
		User: RequestCreateSessionUserCredentials{
			UserID:   Conf.UserID,
			Password: Conf.Password,
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
	template, err := f.ReadFile("report_template.html")
	if err != nil {
		panic(err)
	}
	rep := NewReportGenerator(string(template))
	rep.Report.GeneratedOn = time.Now()
	rep.Report.Registries = make([]RegistryReport, 0)
	for r := range session.ListRegistries() {
		registry := RegistryReport{
			Name:   r.Name,
			Images: make([]ImageReport, 0),
		}
		//fmt.Println("Registry:", r.ID)
		for im := range session.ListRegistryImages(r.ID) {
			//fmt.Println("Image:", im.ID) //, im.Tag, im.Registry, im.Repository, im.Status)
			scan := session.ImageLastScan(im)
			image := ImageReport{
				Name:   scan.Name,
				Layers: make([]LayerReport, 0),
			}
			rep.Report.Completed = scan.Details.Completed
			for _, layer := range scan.Details.Results {
				layerReport := LayerReport{
					ID:        layer.ID,
					CreatedAt: layer.CreatedAt,
					CreatedBy: layer.CreatedBy,
				}

				//fmt.Println("Result:")
				if layer.Malware+layer.Vulnerabilities+layer.Contents == "" {
					continue
				}
				fmt.Println("==========================")
				fmt.Println("Scan competed:", scan.Details.Completed)
				fmt.Println("Image:", scan.Name)
				fmt.Println("Layer ID:", layer.ID)
				fmt.Println("Created at:", layer.CreatedAt)
				fmt.Println("Create with:", layer.CreatedBy)
				fmt.Println("Findings:")

				if layer.Malware != "" {
					layerReport.Malware = make([]MalwareReport, 0)
					/*
						JSON, err := json.MarshalIndent(layer, "", "  ")
						if err != nil {
							panic(err)
						}
						fmt.Printf("Layer:\n%s\n", string(JSON))
					*/
					for malware := range session.ListMalwareFindings(layer.Malware) {
						name := malware.Icrc.Name
						url := malware.Icrc.URL
						if malware.Trendx.Found.Name != "" {
							name = fmt.Sprintf("%s (Confidence %s%%)",
								malware.Trendx.Found.Name,
								malware.Trendx.Confidence)
							url = malware.Trendx.Found.URL
						}
						malwareReport := MalwareReport{
							Filename: malware.Filename,
							Name:     name,
							URL:      url,
						}
						fmt.Printf("Malware: %s %s (%s)\n", malware.Filename, name, url)
						// Fixed in ohter layres?!
						layerReport.Malware = append(layerReport.Malware, malwareReport)
					}
				}
				if layer.Vulnerabilities != "" {
					layerReport.Package = make([]PackageReport, 0)
					for v := range session.ListVulnerabilitiesFindings(layer.Vulnerabilities) {
						packageReport := PackageReport{
							Name:            v.Name,
							Version:         v.Version,
							Vulnerabilities: make([]VulnerabilityReport, 0),
						}
						fmt.Println("Module/Package:", v.Name)
						fmt.Println("Version:", v.Version)
						for _, cve := range v.Vulnerabilities {
							vulnerabilityReport := VulnerabilityReport{
								Name:     cve.Name,
								Link:     cve.Link,
								Severity: cve.Severity,
							}

							fmt.Println("CVE:", cve.Name)
							fmt.Println("URL:", cve.Link)
							fmt.Println("Severity:", cve.Severity)
							// cve.Description
							// Overriden!!!
							packageReport.Vulnerabilities = append(packageReport.Vulnerabilities, vulnerabilityReport)
						}
						layerReport.Package = append(layerReport.Package, packageReport)
					}
				}
				if layer.Contents != "" {
					layerReport.Contents = make([]ContentsReport, 0)

					for contents := range session.ListContentsFindings(layer.Contents) {
						contentsReport := ContentsReport{
							Severity:    contents.Severity,
							Rule:        contents.Rule,
							Description: contents.Description,
							Filename:    contents.Filename,
						}
						/*JSON, err := json.MarshalIndent(contents, "", "  ")
						if err != nil {
							panic(err)
						}
						fmt.Printf("Content:\n%s\n", string(JSON))
						*/
						fmt.Println("Severity: ", contents.Severity)
						fmt.Println("Rule: ", contents.Rule)
						fmt.Println("Description: ", contents.Description)
						fmt.Println("File: ", contents.Filename)
						layerReport.Contents = append(layerReport.Contents, contentsReport)
					}
				}
				image.Layers = append(image.Layers, layerReport)
			}
			registry.Images = append(registry.Images, image)
		}
		rep.Report.Registries = append(rep.Report.Registries, registry)

	}
	rep.Report.Difference = rep.Report.GeneratedOn.Sub(rep.Report.Completed)
	filename := rep.Report.GeneratedOn.Format("report_20060102.html")
	//filename := "report.html"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = rep.Generate(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output writted to: %s\n", filename)
}
