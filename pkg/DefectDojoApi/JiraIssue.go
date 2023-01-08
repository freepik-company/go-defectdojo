package DefectDojoApi

import "time"

type JiraIssue struct {
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
}
