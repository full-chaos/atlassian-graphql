package mappers

import (
	"errors"
	"strings"

	"atlassian-graphql/graphql/canonical"
	"atlassian-graphql/graphql/gen"
)

func ProjectWithOpsgenieTeams(cloudID string, project gen.JiraProjectNode, teams []gen.OpsgenieTeamNode) (canonical.CanonicalProjectWithOpsgenieTeams, error) {
	cloud := strings.TrimSpace(cloudID)
	if cloud == "" {
		return canonical.CanonicalProjectWithOpsgenieTeams{}, errors.New("cloudID is required")
	}

	key := strings.TrimSpace(project.Key)
	if key == "" {
		return canonical.CanonicalProjectWithOpsgenieTeams{}, errors.New("project.key is required")
	}
	name := strings.TrimSpace(project.Name)
	if name == "" {
		return canonical.CanonicalProjectWithOpsgenieTeams{}, errors.New("project.name is required")
	}

	seen := map[string]struct{}{}
	outTeams := make([]canonical.OpsgenieTeamRef, 0, len(teams))
	for _, t := range teams {
		id := strings.TrimSpace(t.ID)
		if id == "" {
			return canonical.CanonicalProjectWithOpsgenieTeams{}, errors.New("opsgenie_team.id is required")
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}

		teamName := strings.TrimSpace(t.Name)
		if teamName == "" {
			return canonical.CanonicalProjectWithOpsgenieTeams{}, errors.New("opsgenie_team.name is required")
		}
		outTeams = append(outTeams, canonical.OpsgenieTeamRef{ID: id, Name: teamName})
	}

	return canonical.CanonicalProjectWithOpsgenieTeams{
		Project: canonical.JiraProject{
			CloudID: cloud,
			Key:     key,
			Name:    name,
		},
		OpsgenieTeams: outTeams,
	}, nil
}

