package DefectDojoApi

import "time"

type Test struct {
	Id                   int            `json:"id"`
	Tags                 []string       `json:"tags"`
	TestTypeName         string         `json:"test_type_name"`
	FindingGroups        []FindingGroup `json:"finding_groups"`
	ScanType             string         `json:"scan_type"`
	Title                string         `json:"title"`
	Description          string         `json:"description"`
	TargetStart          time.Time      `json:"target_start"`
	TargetEnd            time.Time      `json:"target_end"`
	EstimatedTime        string         `json:"estimated_time"`
	ActualTime           string         `json:"actual_time"`
	PercentComplete      int            `json:"percent_complete"`
	Updated              time.Time      `json:"updated"`
	Created              time.Time      `json:"created"`
	Version              string         `json:"version"`
	BuildId              string         `json:"build_id"`
	CommitHash           string         `json:"commit_hash"`
	BranchTag            string         `json:"branch_tag"`
	Engagement           int            `json:"engagement"`
	Lead                 int            `json:"lead"`
	TestType             int            `json:"test_type"`
	Environment          int            `json:"environment"`
	ApiScanConfiguration int            `json:"api_scan_configuration"`
	Notes                []Note         `json:"notes"`
	Files                []File         `json:"files"`
}
