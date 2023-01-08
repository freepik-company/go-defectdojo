package DefectDojoApi

import (
	jsonPkg "encoding/json"
	"log"
	"time"
)

type Engagement struct {
	Id                         int       `json:"id,omitempty"`
	Tags                       []string  `json:"tags,omitempty"`
	Name                       string    `json:"name,omitempty"`
	Description                string    `json:"description,omitempty"`
	Version                    string    `json:"version,omitempty"`
	FirstContacted             string    `json:"first_contacted,omitempty"`
	TargetStart                string    `json:"target_start,omitempty"`
	TargetEnd                  string    `json:"target_end,omitempty"`
	Reason                     string    `json:"reason,omitempty"`
	Updated                    time.Time `json:"updated,omitempty"`
	Created                    time.Time `json:"created,omitempty"`
	Active                     bool      `json:"active,omitempty"`
	Tracker                    string    `json:"tracker,omitempty"`
	TestStrategy               string    `json:"test_strategy,omitempty"`
	ThreatModel                bool      `json:"threat_model,omitempty"`
	ApiTest                    bool      `json:"api_test,omitempty"`
	PenTest                    bool      `json:"pen_test,omitempty"`
	CheckList                  bool      `json:"check_list,omitempty"`
	Status                     string    `json:"status,omitempty"`
	Progress                   string    `json:"progress,omitempty"`
	TmodelPath                 string    `json:"tmodel_path,omitempty"`
	DoneTesting                bool      `json:"done_testing,omitempty"`
	EngagementType             string    `json:"engagement_type,omitempty"`
	BuildId                    string    `json:"build_id,omitempty"`
	CommitHash                 string    `json:"commit_hash,omitempty"`
	BranchTag                  string    `json:"branch_tag,omitempty"`
	SourceCodeManagementUri    string    `json:"source_code_management_uri,omitempty"`
	DeduplicationOnEngagement  bool      `json:"deduplication_on_engagement,omitempty"`
	Lead                       int       `json:"lead,omitempty"`
	Requester                  int       `json:"requester,omitempty"`
	Preset                     int       `json:"preset,omitempty"`
	ReportType                 int       `json:"report_type,omitempty"`
	Product                    int       `json:"product,omitempty"`
	BuildServer                int       `json:"build_server,omitempty"`
	SourceCodeManagementServer int       `json:"source_code_management_server,omitempty"`
	OrchestrationEngine        int       `json:"orchestration_engine,omitempty"`
	Notes                      []Note    `json:"notes,omitempty"`
	Files                      []File    `json:"files,omitempty"`
	RiskAcceptance             []int     `json:"risk_acceptance,omitempty"`
}

func (this Engagement) String() string {
	json, err := jsonPkg.Marshal(this)

	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}

func NewEngagement(targetStart string, targetEnd string, product int) *Engagement {
	return &Engagement{TargetStart: targetStart, TargetEnd: targetEnd, Product: product}
}

func newEngagementFromJson(json string) *Engagement {
	this := new(Engagement)
	err := jsonPkg.Unmarshal([]byte(json), &this)

	if err != nil {
		log.Fatal(err)
	}

	return this
}
