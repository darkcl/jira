package helpers

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	models "github.com/darkcl/jira/models"
)

// GetEpicFromIssue - Get Epic from issue
func GetEpicFromIssue(issueID string, client *jira.Client) *jira.Issue {
	// call editmeta to get epic id
	editMetaEndpoint := fmt.Sprintf("rest/api/2/issue/%s/editmeta", issueID)
	req, _ := client.NewRequest("GET", editMetaEndpoint, nil)
	meta := new(models.FieldResponse)
	_, err := client.Do(req, meta)
	if err != nil {
		panic(err)
	}
	var fieldsName string
	for k, v := range meta.Fields {
		if v.Schema.Custom == "com.pyxis.greenhopper.jira:gh-epic-link" {
			fieldsName = k
		}
	}

	issue, _, _ := client.Issue.Get(issueID, nil)

	epicNumber := fmt.Sprintf("%v", issue.Fields.Unknowns[fieldsName])
	epic, _, _ := client.Issue.Get(epicNumber, nil)

	return epic
}
