package main

import "time"

type ResponseListScans struct {
	Scans []struct {
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
	} `json:"scans"`
	Next string `json:"next"`
}
