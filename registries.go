package main

import "time"

type (
	ResponseRegistry struct {
		ID                 string `json:"id"`
		Href               string `json:"href"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Host               string `json:"host"`
		InsecureSkipVerify bool   `json:"insecureSkipVerify"`
		RootCAs            string `json:"rootCAs"`
		Filter             struct {
			Include []string `json:"include"`
			Exclude []string `json:"exclude"`
		} `json:"filter"`
		Metrics struct {
			Content struct {
				Repositories int `json:"repositories"`
				Images       int `json:"images"`
			} `json:"content"`
			Scans struct {
				Pending    int `json:"pending"`
				InProgress int `json:"inProgress"`
				Ok         int `json:"ok"`
				Warnings   int `json:"warnings"`
				Errors     int `json:"errors"`
				Cancelled  int `json:"cancelled"`
			} `json:"scans"`
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
			History struct {
				OneD struct {
					Scans int `json:"scans"`
				} `json:"1d"`
				OneW struct {
					Scans int `json:"scans"`
				} `json:"1w"`
			} `json:"history"`
		} `json:"metrics"`
		Schedule     bool      `json:"schedule"`
		Status       string    `json:"status"`
		StatusDetail string    `json:"statusDetail"`
		Created      time.Time `json:"created"`
		Updated      time.Time `json:"updated"`
	}

	ResponseImage struct {
		ID         string `json:"id"`
		Href       string `json:"href"`
		Registry   string `json:"registry"`
		Repository string `json:"repository"`
		Tag        string `json:"tag"`
		Digest     string `json:"digest"`
		Status     string `json:"status"`
		Findings   struct {
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
		Updated time.Time `json:"updated"`
	}
)
