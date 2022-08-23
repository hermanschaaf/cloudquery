package client

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"
)

type HerokuServices struct {
	Teams TeamsService
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_teams.go . TeamsService
type TeamsService interface {
	TeamList(ctx context.Context, lr *heroku.ListRange) (heroku.TeamListResult, error)
}
