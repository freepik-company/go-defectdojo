package DefectDojoApi

import (
	"time"
)

type Finding struct {
	AcceptedRisks           []interface{}   `json:"accepted_risks"`
	Active                  bool            `json:"active"`
	Age                     int             `json:"age"`
	ComponentName           interface{}     `json:"component_name"`
	ComponentVersion        interface{}     `json:"component_version"`
	Created                 time.Time       `json:"created"`
	Cvssv3                  string          `json:"cvssv3,omitempty"`
	Cvssv3Score             string          `json:"cvssv3_score,omitempty"`
	Cwe                     interface{}     `json:"cwe"`
	Date                    string          `json:"date"`
	DefectReviewRequestedBy *int            `json:"defect_review_requested_by"`
	Description             string          `json:"description"`
	DisplayStatus           string          `json:"display_status"`
	Duplicate               bool            `json:"duplicate"`
	DuplicateFinding        interface{}     `json:"duplicate_finding"`
	DynamicFinding          bool            `json:"dynamic_finding"`
	Endpoints               []int           `json:"endpoints"`
	FalseP                  bool            `json:"false_p"`
	FilePath                interface{}     `json:"file_path"`
	Files                   []interface{}   `json:"files"`
	FindingGroups           []interface{}   `json:"finding_groups"`
	FindingMeta             []interface{}   `json:"finding_meta"`
	FoundBy                 []int           `json:"found_by"`
	HashCode                string          `json:"hash_code"`
	Id                      int             `json:"id"`
	Impact                  string          `json:"impact"`
	IsMitigated             bool            `json:"is_mitigated"`
	JiraChange              time.Time       `json:"jira_change"`
	JiraCreation            time.Time       `json:"jira_creation"`
	LastReviewed            *time.Time      `json:"last_reviewed"`
	LastReviewedBy          *int            `json:"last_reviewed_by"`
	LastStatusUpdate        time.Time       `json:"last_status_update"`
	Line                    interface{}     `json:"line"`
	Mitigated               *time.Time      `json:"mitigated"`
	MitigatedBy             *int            `json:"mitigated_by"`
	Mitigation              string          `json:"mitigation"`
	NbOccurences            interface{}     `json:"nb_occurences"`
	Notes                   []Note          `json:"notes"`
	NumericalSeverity       string          `json:"numerical_severity"`
	OutOfScope              bool            `json:"out_of_scope"`
	Param                   interface{}     `json:"param"`
	Payload                 interface{}     `json:"payload"`
	PlannedRemediationDate  interface{}     `json:"planned_remediation_date"`
	PublishDate             interface{}     `json:"publish_date"`
	PushToJira              bool            `json:"push_to_jira"`
	References              string          `json:"references"`
	RelatedFields           interface{}     `json:"related_fields"`
	Reporter                int             `json:"reporter,omitempty"`
	RequestResponse         RequestResponse `json:"request_response"`
	ReviewRequestedBy       interface{}     `json:"review_requested_by"`
	Reviewers               []int           `json:"reviewers,omitempty"`
	RiskAccepted            bool            `json:"risk_accepted"`
	SastSinkObject          string          `json:"sast_sink_object"`
	SastSourceFilePath      string          `json:"sast_source_file_path"`
	SastSourceLine          int             `json:"sast_source_line"`
	SastSourceObject        string          `json:"sast_source_object"`
	ScannerConfidence       interface{}     `json:"scanner_confidence"`
	Service                 interface{}     `json:"service"`
	Severity                string          `json:"severity"`
	SeverityJustification   string          `json:"severity_justification"`
	SlaDaysRemaining        int             `json:"sla_days_remaining"`
	SlaStartDate            interface{}     `json:"sla_start_date"`
	SonarqubeIssue          interface{}     `json:"sonarqube_issue"`
	StaticFinding           bool            `json:"static_finding"`
	StepsToReproduce        string          `json:"steps_to_reproduce"`
	Tags                    []string        `json:"tags,omitempty"`
	Test                    int             `json:"test"`
	ThreadId                int             `json:"thread_id"`
	Title                   string          `json:"title"`
	UnderDefectReview       bool            `json:"under_defect_review"`
	UnderReview             bool            `json:"under_review"`
	UniqueIdFromTool        interface{}     `json:"unique_id_from_tool"`
	Url                     interface{}     `json:"url"`
	Verified                bool            `json:"verified"`
	VulnIdFromTool          interface{}     `json:"vuln_id_from_tool"`
	VulnerabilityIds        []int           `json:"vulnerability_ids,omitempty"`
}

func NewFinding(
	title string,
	severity string,
	description string,
	active bool,
	verified bool,
	numericalSeverity string,
	test int,
	foundBy []int,
) *Finding {
	return &Finding{
		Title:             title,
		Severity:          severity,
		Description:       description,
		Active:            active,
		Verified:          verified,
		NumericalSeverity: numericalSeverity,
		Test:              test,
		FoundBy:           foundBy,
		Date:              time.Now().Format(YYYYMMDD),
	}
}
