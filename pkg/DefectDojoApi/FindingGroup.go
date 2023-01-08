package DefectDojoApi

type FindingGroup struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Test      int       `json:"test"`
	JiraIssue JiraIssue `json:"jira_issue"`
}
