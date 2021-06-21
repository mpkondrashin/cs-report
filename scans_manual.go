package main

type (
	ResponseListScansScanSource struct {
		Type               string
		Registry           string
		Repository         string
		Tag                string
		Digest             string
		insecureSkipVerify bool
		RootCAs            string
	}

	ResponseListScansScanContext struct {
		//
	}

	ResponseListScansScanDelailsResult struct {
		Id              string
		CreatedAt       string
		CreatedBy       string
		Malware         string
		Vulnerabilities string
		Contents        string
		Findings        ResponseListScansScanDelailsResultFindings
	}
	ResponseListScansScanDelailsResultFindings struct {
		Scanners ResponseListScansScanDelailsResultFindingsScanners
		Malware  int
	}
	ResponseListScansScanDelailsResultFindingsScanners struct {
		Malware               ResponseListScansScanDelailsResultFindingsScannersMalware
		Vulnerabilities       ResponseListScansScanDelailsResultFindingsScannersVulnerabilities
		VulnerabilityPatterns ResponseListScansScanDelailsResultFindingsScannersVulnerabilityPatterns
	}
	ResponseListScansScanDelailsResultFindingsScannersMalware struct {
		Status   string
		Updated  string
		Versions ResponseListScansScanDelailsResultFindingsScannersMalwareVersions
	}
	ResponseListScansScanDelailsResultFindingsScannersMalwareVersions struct {
		Icrc      string
		TrendX    string
		Denylist  string
		Blacklist string
	}

	ResponseListScansScanDelailsResultFindingsScannersVulnerabilities struct {
		Status  string
		Date    string
		Updated string
	}
	ResponseListScansScanDelailsResultFindingsScannersVulnerabilityPatterns struct {
		Status  string
		Date    string
		Updated string
	}

	ResponseListScansScanDelailsResultFindingsContents struct {
		Total      ResponseListScansScanDelailsResultFindingsContentsTotal
		Unresolved ResponseListScansScanDelailsResultFindingsContentsTotalUnresolved
	}
	ResponseListScansScanDelailsResultFindingsContentsTotal struct {
		Defcon1    int
		Critical   int
		High       int
		Medium     int
		Low        int
		Negligible int
		Unknown    int
	}
	ResponseListScansScanDelailsResultFindingsContentsTotalUnresolved struct {
	}
	ResponseListScansScanDelailsResultFindingsVulnerabilities struct {
	}
	ResponseListScansScanDelailsResultFindingsChecklists struct {
	}

	/*},
	  "malware": 0,
	  "contents": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  },
	  "vulnerabilities": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "fixAvailable": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  },
	  "checklists": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  }
	  }
	  }
	  ],
	  "checklists": "/api/scans/60e53669-c8ef-4d0f-a8ff-3dbbb098d8ff/checklists"
	  },
	  "findings": {
	  "scanners": {
	  "malware": {
	  "status": "ok",
	  "updated": "2018-05-01T00:00:00Z",
	  "versions": {
	  "icrc": "string",
	  "trendX": "string",
	  "denylist": "string",
	  "blacklist": "string"
	  }
	  },
	  "vulnerabilities": {
	  "status": "ok",
	  "updated": "2018-05-01T00:00:00Z"
	  },
	  "vulnerabilityPatterns": {
	  "status": "ok",
	  "date": "2018-05-01T00:00:00Z",
	  "updated": "2018-05-01T00:00:00Z"
	  }
	  },
	  "malware": 0,
	  "contents": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  },
	  "vulnerabilities": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "fixAvailable": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  },
	  "checklists": {
	  "total": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  },
	  "unresolved": {
	  "defcon1": 0,
	  "critical": 0,
	  "high": 0,
	  "medium": 0,
	  "low": 0,
	  "negligible": 0,
	  "unknown": 0
	  }
	  }
	  }
	  }
	*/

	ResponseListScansScanDelails struct {
		Detail       string
		Requested    string
		Started      string
		Updated      string
		Completed    string
		Digest       string
		Os           string
		Architecture string
		// Labels ...
		Results []ResponseListScansScanDelailsResult
	}

	ResponseListScansScan struct {
		Id     string
		Href   string
		Name   string
		Source ResponseListScansScanSource
		// Context ResponseListScansScanContext
		Status string // "completed-with-findings"/ failed
		//	Details ""results": [
	}
)
