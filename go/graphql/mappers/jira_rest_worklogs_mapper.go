package mappers

import (
	"errors"
	"strings"

	"atlassian-graphql/graphql/canonical"
	"atlassian-graphql/graphql/gen"
)

func JiraWorklogFromREST(issueKey string, worklog gen.Worklog) (canonical.JiraWorklog, error) {
	issue := strings.TrimSpace(issueKey)
	if issue == "" {
		return canonical.JiraWorklog{}, errors.New("issueKey is required")
	}
	if worklog.ID == nil || strings.TrimSpace(*worklog.ID) == "" {
		return canonical.JiraWorklog{}, errors.New("worklog.id is required")
	}
	if worklog.Started == nil || strings.TrimSpace(*worklog.Started) == "" {
		return canonical.JiraWorklog{}, errors.New("worklog.started is required")
	}
	if worklog.TimeSpentSeconds == nil || *worklog.TimeSpentSeconds < 0 {
		return canonical.JiraWorklog{}, errors.New("worklog.timeSpentSeconds is required and must be >= 0")
	}
	if worklog.Created == nil || strings.TrimSpace(*worklog.Created) == "" {
		return canonical.JiraWorklog{}, errors.New("worklog.created is required")
	}
	if worklog.Updated == nil || strings.TrimSpace(*worklog.Updated) == "" {
		return canonical.JiraWorklog{}, errors.New("worklog.updated is required")
	}

	var author *canonical.JiraUser
	if worklog.Author != nil {
		if worklog.Author.AccountID == nil || strings.TrimSpace(*worklog.Author.AccountID) == "" {
			return canonical.JiraWorklog{}, errors.New("worklog.author.accountId is required")
		}
		if worklog.Author.DisplayName == nil || strings.TrimSpace(*worklog.Author.DisplayName) == "" {
			return canonical.JiraWorklog{}, errors.New("worklog.author.displayName is required")
		}
		u := canonical.JiraUser{
			AccountID:   strings.TrimSpace(*worklog.Author.AccountID),
			DisplayName: strings.TrimSpace(*worklog.Author.DisplayName),
		}
		if worklog.Author.EmailAddress != nil && strings.TrimSpace(*worklog.Author.EmailAddress) != "" {
			v := strings.TrimSpace(*worklog.Author.EmailAddress)
			u.Email = &v
		}
		author = &u
	}

	return canonical.JiraWorklog{
		IssueKey:         issue,
		WorklogID:        strings.TrimSpace(*worklog.ID),
		Author:           author,
		StartedAt:        strings.TrimSpace(*worklog.Started),
		TimeSpentSeconds: *worklog.TimeSpentSeconds,
		CreatedAt:        strings.TrimSpace(*worklog.Created),
		UpdatedAt:        strings.TrimSpace(*worklog.Updated),
	}, nil
}
