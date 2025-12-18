package mappers

import (
	"errors"
	"fmt"
	"strings"

	"atlassian-graphql/graphql/canonical"
	"atlassian-graphql/graphql/gen"
)

func JiraChangelogEventFromREST(issueKey string, changelog gen.Changelog) (canonical.JiraChangelogEvent, error) {
	issue := strings.TrimSpace(issueKey)
	if issue == "" {
		return canonical.JiraChangelogEvent{}, errors.New("issueKey is required")
	}
	if changelog.ID == nil || strings.TrimSpace(*changelog.ID) == "" {
		return canonical.JiraChangelogEvent{}, errors.New("changelog.id is required")
	}
	if changelog.Created == nil || strings.TrimSpace(*changelog.Created) == "" {
		return canonical.JiraChangelogEvent{}, errors.New("changelog.created is required")
	}

	eventID := strings.TrimSpace(*changelog.ID)
	createdAt := strings.TrimSpace(*changelog.Created)

	items := make([]canonical.JiraChangelogItem, 0, len(changelog.Items))
	for idx, it := range changelog.Items {
		if it.Field == nil || strings.TrimSpace(*it.Field) == "" {
			return canonical.JiraChangelogEvent{}, fmt.Errorf("changelog.items[%d].field is required", idx)
		}
		item := canonical.JiraChangelogItem{
			Field: strings.TrimSpace(*it.Field),
		}
		if it.From != nil && strings.TrimSpace(*it.From) != "" {
			v := strings.TrimSpace(*it.From)
			item.From = &v
		}
		if it.To != nil && strings.TrimSpace(*it.To) != "" {
			v := strings.TrimSpace(*it.To)
			item.To = &v
		}
		if it.FromString != nil && strings.TrimSpace(*it.FromString) != "" {
			v := strings.TrimSpace(*it.FromString)
			item.FromString = &v
		}
		if it.ToString != nil && strings.TrimSpace(*it.ToString) != "" {
			v := strings.TrimSpace(*it.ToString)
			item.ToString = &v
		}
		items = append(items, item)
	}
	if len(items) == 0 {
		return canonical.JiraChangelogEvent{}, errors.New("changelog.items is required")
	}

	var author *canonical.JiraUser
	if changelog.Author != nil {
		if changelog.Author.AccountID == nil || strings.TrimSpace(*changelog.Author.AccountID) == "" {
			return canonical.JiraChangelogEvent{}, errors.New("changelog.author.accountId is required")
		}
		if changelog.Author.DisplayName == nil || strings.TrimSpace(*changelog.Author.DisplayName) == "" {
			return canonical.JiraChangelogEvent{}, errors.New("changelog.author.displayName is required")
		}
		u := canonical.JiraUser{
			AccountID:   strings.TrimSpace(*changelog.Author.AccountID),
			DisplayName: strings.TrimSpace(*changelog.Author.DisplayName),
		}
		if changelog.Author.EmailAddress != nil && strings.TrimSpace(*changelog.Author.EmailAddress) != "" {
			v := strings.TrimSpace(*changelog.Author.EmailAddress)
			u.Email = &v
		}
		author = &u
	}

	return canonical.JiraChangelogEvent{
		IssueKey:  issue,
		EventID:   eventID,
		Author:    author,
		CreatedAt: createdAt,
		Items:     items,
	}, nil
}
