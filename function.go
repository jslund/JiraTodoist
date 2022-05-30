package function

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("IssueUpdated", issueUpdated)
}

func issueUpdated(w http.ResponseWriter, r *http.Request) {
	var issue JiraIssue

	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		http.Error(w, "Error parsing requests", http.StatusBadRequest)
	}

	summary := issue.Issue.Fields.Summary
	priority, _ := strconv.Atoi(issue.Issue.Fields.Priority.ID)

	if err := AddItem(summary, priority); err != nil {
		http.Error(w, "Could not add item", http.StatusInternalServerError)
		log.Fatalf("Could not add item %v", summary)
	}

	log.Printf("Title %s", issue.Issue.Fields.Summary)

}
