package main

import "time"

type (
	ResponseListScans struct {
		Scans []ResponseScan `json:"scans"`
		Next  string         `json:"next"`
	}

	ResponseScan struct {
		ID     string `json:"id"`
		Href   string `json:"href"`
		Name   string `json:"name"`
		Source struct {
			Type               string `json:"type"`
			Registry           string `json:"registry"`
			Repository         string `json:"repository"`
			Tag                string `json:"tag"`
			Digest             string `json:"digest"`
			InsecureSkipVerify bool   `json:"insecureSkipVerify"`
			RootCAs            string `json:"rootCAs"`
		} `json:"source"`
		Context struct {
			Key string `json:"key"`
		} `json:"context"`
		Status  string `json:"status"`
		Details struct {
			Detail       string    `json:"detail"`
			Requested    time.Time `json:"requested"`
			Started      time.Time `json:"started"`
			Updated      time.Time `json:"updated"`
			Completed    time.Time `json:"completed"`
			Digest       string    `json:"digest"`
			Os           string    `json:"os"`
			Architecture string    `json:"architecture"`
			Labels       struct {
				Key string `json:"key"`
			} `json:"labels"`
			Results []struct {
				ID              string    `json:"id"`
				CreatedAt       time.Time `json:"createdAt"`
				CreatedBy       string    `json:"createdBy"`
				Malware         string    `json:"malware"`
				Vulnerabilities string    `json:"vulnerabilities"`
				Contents        string    `json:"contents"`
				Findings        struct {
					Scanners struct {
						Malware struct {
							Status   string    `json:"status"`
							Updated  time.Time `json:"updated"`
							Versions struct {
								Icrc      string `json:"icrc"`
								TrendX    string `json:"trendX"`
								Denylist  string `json:"denylist"`
								Blacklist string `json:"blacklist"`
							} `json:"versions"`
						} `json:"malware"`
						Vulnerabilities struct {
							Status  string    `json:"status"`
							Updated time.Time `json:"updated"`
						} `json:"vulnerabilities"`
						VulnerabilityPatterns struct {
							Status  string    `json:"status"`
							Date    time.Time `json:"date"`
							Updated time.Time `json:"updated"`
						} `json:"vulnerabilityPatterns"`
					} `json:"scanners"`
					Malware  int `json:"malware"`
					Contents struct {
						Total struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"total"`
						Unresolved struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"unresolved"`
					} `json:"contents"`
					Vulnerabilities struct {
						Total struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"total"`
						Unresolved struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"unresolved"`
						FixAvailable struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"fixAvailable"`
					} `json:"vulnerabilities"`
					Checklists struct {
						Total struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"total"`
						Unresolved struct {
							Defcon1    int `json:"defcon1"`
							Critical   int `json:"critical"`
							High       int `json:"high"`
							Medium     int `json:"medium"`
							Low        int `json:"low"`
							Negligible int `json:"negligible"`
							Unknown    int `json:"unknown"`
						} `json:"unresolved"`
					} `json:"checklists"`
				} `json:"findings"`
			} `json:"results"`
			Checklists string `json:"checklists"`
		} `json:"details"`
		Findings struct {
			Scanners struct {
				Malware struct {
					Status   string    `json:"status"`
					Updated  time.Time `json:"updated"`
					Versions struct {
						Icrc      string `json:"icrc"`
						TrendX    string `json:"trendX"`
						Denylist  string `json:"denylist"`
						Blacklist string `json:"blacklist"`
					} `json:"versions"`
				} `json:"malware"`
				Vulnerabilities struct {
					Status  string    `json:"status"`
					Updated time.Time `json:"updated"`
				} `json:"vulnerabilities"`
				VulnerabilityPatterns struct {
					Status  string    `json:"status"`
					Date    time.Time `json:"date"`
					Updated time.Time `json:"updated"`
				} `json:"vulnerabilityPatterns"`
			} `json:"scanners"`
			Malware  int `json:"malware"`
			Contents struct {
				Total struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"total"`
				Unresolved struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"unresolved"`
			} `json:"contents"`
			Vulnerabilities struct {
				Total struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"total"`
				Unresolved struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"unresolved"`
				FixAvailable struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"fixAvailable"`
			} `json:"vulnerabilities"`
			Checklists struct {
				Total struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"total"`
				Unresolved struct {
					Defcon1    int `json:"defcon1"`
					Critical   int `json:"critical"`
					High       int `json:"high"`
					Medium     int `json:"medium"`
					Low        int `json:"low"`
					Negligible int `json:"negligible"`
					Unknown    int `json:"unknown"`
				} `json:"unresolved"`
			} `json:"checklists"`
		} `json:"findings"`
	}

	ResponseLayerMalware struct {
		Filename string `json:"filename"`
		Icrc     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"icrc"`
		Trendx struct {
			Found struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"found"`
			Confidence int `json:"confidence"`
			Related    []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"related"`
		} `json:"trendx"`
	}

	ResponseLayerVulnerabilities struct {
		Name            string `json:"name"`
		NamespaceName   string `json:"namespaceName"`
		Version         string `json:"version"`
		VersionFormat   string `json:"versionFormat"`
		Vulnerabilities []struct {
			Description string `json:"description"`
			FixedBy     string `json:"fixedBy"`
			Fixed       struct {
				Name          string `json:"name"`
				NamespaceName string `json:"namespaceName"`
				Version       string `json:"version"`
				VersionFormat string `json:"versionFormat"`
				Layer         string `json:"layer"`
			} `json:"fixed"`
			Override struct {
				ID         string    `json:"id"`
				Href       string    `json:"href"`
				Name       string    `json:"name"`
				Package    string    `json:"package"`
				Version    string    `json:"version"`
				Registry   string    `json:"registry"`
				Repository string    `json:"repository"`
				Tag        string    `json:"tag"`
				Created    time.Time `json:"created"`
				Updated    time.Time `json:"updated"`
				Expires    time.Time `json:"expires"`
				Comment    string    `json:"comment"`
			} `json:"override"`
			Link          string `json:"link"`
			Name          string `json:"name"`
			NamespaceName string `json:"namespaceName"`
			Severity      string `json:"severity"`
			Metadata      struct {
			} `json:"metadata"`
		} `json:"vulnerabilities"`
		Metrics struct {
			Total struct {
				Defcon1    int `json:"defcon1"`
				Critical   int `json:"critical"`
				High       int `json:"high"`
				Medium     int `json:"medium"`
				Low        int `json:"low"`
				Negligible int `json:"negligible"`
				Unknown    int `json:"unknown"`
			} `json:"total"`
			Unresolved struct {
				Defcon1    int `json:"defcon1"`
				Critical   int `json:"critical"`
				High       int `json:"high"`
				Medium     int `json:"medium"`
				Low        int `json:"low"`
				Negligible int `json:"negligible"`
				Unknown    int `json:"unknown"`
			} `json:"unresolved"`
			FixAvailable struct {
				Defcon1    int `json:"defcon1"`
				Critical   int `json:"critical"`
				High       int `json:"high"`
				Medium     int `json:"medium"`
				Low        int `json:"low"`
				Negligible int `json:"negligible"`
				Unknown    int `json:"unknown"`
			} `json:"fixAvailable"`
		} `json:"metrics"`
	}
)
