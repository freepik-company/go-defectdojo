package DefectDojoApi

import "time"

const (
	YYYYMMDD = "2006-01-02"
)

type Author struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Note struct {
	Id       int           `json:"id"`
	Author   Author        `json:"author"`
	Editor   interface{}   `json:"editor"`
	History  []interface{} `json:"history"`
	Entry    string        `json:"entry"`
	Date     time.Time     `json:"date"`
	Private  bool          `json:"private"`
	Edited   bool          `json:"edited"`
	EditTime time.Time     `json:"edit_time"`
	NoteType interface{}   `json:"note_type"`
}

type ReqResp struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type RequestResponse struct {
	ReqResp []ReqResp `json:"req_resp"`
}

type FindingResults struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Finding `json:"results"`
	Prefetch struct {
	} `json:"prefetch"`
}

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

type Test struct {
	Id            int      `json:"id"`
	Tags          []string `json:"tags"`
	TestTypeName  string   `json:"test_type_name"`
	FindingGroups []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Test      int    `json:"test"`
		JiraIssue struct {
			Id           int       `json:"id"`
			Url          string    `json:"url"`
			JiraId       string    `json:"jira_id"`
			JiraKey      string    `json:"jira_key"`
			JiraCreation time.Time `json:"jira_creation"`
			JiraChange   time.Time `json:"jira_change"`
			JiraProject  int       `json:"jira_project"`
			Finding      int       `json:"finding"`
			Engagement   int       `json:"engagement"`
			FindingGroup int       `json:"finding_group"`
		} `json:"jira_issue"`
	} `json:"finding_groups"`
	ScanType             string    `json:"scan_type"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	TargetStart          time.Time `json:"target_start"`
	TargetEnd            time.Time `json:"target_end"`
	EstimatedTime        string    `json:"estimated_time"`
	ActualTime           string    `json:"actual_time"`
	PercentComplete      int       `json:"percent_complete"`
	Updated              time.Time `json:"updated"`
	Created              time.Time `json:"created"`
	Version              string    `json:"version"`
	BuildId              string    `json:"build_id"`
	CommitHash           string    `json:"commit_hash"`
	BranchTag            string    `json:"branch_tag"`
	Engagement           int       `json:"engagement"`
	Lead                 int       `json:"lead"`
	TestType             int       `json:"test_type"`
	Environment          int       `json:"environment"`
	ApiScanConfiguration int       `json:"api_scan_configuration"`
	Notes                []struct {
		Id     int `json:"id"`
		Author struct {
			Id        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"author"`
		Editor struct {
			Id        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"editor"`
		History []struct {
			Id            int `json:"id"`
			CurrentEditor struct {
				Id        int    `json:"id"`
				Username  string `json:"username"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"current_editor"`
			Data     string    `json:"data"`
			Time     time.Time `json:"time"`
			NoteType int       `json:"note_type"`
		} `json:"history"`
		Entry    string    `json:"entry"`
		Date     time.Time `json:"date"`
		Private  bool      `json:"private"`
		Edited   bool      `json:"edited"`
		EditTime time.Time `json:"edit_time"`
		NoteType int       `json:"note_type"`
	} `json:"notes"`
	Files []struct {
		Id    int    `json:"id"`
		File  string `json:"file"`
		Title string `json:"title"`
	} `json:"files"`
}

type TestResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Test `json:"results"`
}
