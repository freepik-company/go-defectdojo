package DefectDojoApi

import (
	jsonPkg "encoding/json"
	"log"
	"time"
)

type Test struct {
	ActualTime           string         `json:"actual_time,omitempty"`
	ApiScanConfiguration int            `json:"api_scan_configuration,omitempty"`
	BranchTag            string         `json:"branch_tag,omitempty"`
	BuildId              string         `json:"build_id,omitempty"`
	CommitHash           string         `json:"commit_hash,omitempty"`
	Created              time.Time      `json:"created,omitempty"`
	Description          string         `json:"description,omitempty"`
	Engagement           int            `json:"engagement"`
	Environment          int            `json:"environment,omitempty"`
	EstimatedTime        string         `json:"estimated_time,omitempty"`
	Files                []File         `json:"files,omitempty"`
	FindingGroups        []FindingGroup `json:"finding_groups,omitempty"`
	Id                   int            `json:"id,omitempty"`
	Lead                 int            `json:"lead,omitempty"`
	Notes                []Note         `json:"notes,omitempty"`
	PercentComplete      int            `json:"percent_complete,omitempty"`
	ScanType             string         `json:"scan_type,omitempty"`
	Tags                 []string       `json:"tags,omitempty"`
	TargetEnd            time.Time      `json:"target_end"`
	TargetStart          time.Time      `json:"target_start"`
	TestType             int            `json:"test_type"`
	TestTypeName         string         `json:"test_type_name,omitempty"`
	Title                string         `json:"title,omitempty"`
	Updated              time.Time      `json:"updated,omitempty"`
	Version              string         `json:"version,omitempty"`
}

func NewTest(
	engagement int,
	targetEnd time.Time,
	targetStart time.Time,
	testType int,
) *Test {
	return &Test{
		Engagement:  engagement,
		TargetEnd:   targetEnd,
		TargetStart: targetStart,
		TestType:    testType,
	}
}

func newTestFromJson(json string) *Test {
	test := new(Test)
	err := jsonPkg.Unmarshal([]byte(json), &test)

	if err != nil {
		log.Fatal(err)
	}

	return test
}

func (this Test) String() string {
	json, err := jsonPkg.Marshal(this)

	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}
