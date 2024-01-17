package access

import (
	"context"

	"github.com/alexandrhub/auth/internal/client/db"
	"github.com/alexandrhub/auth/internal/repository"
	"github.com/alexandrhub/auth/internal/repository/access/model"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.AccessRepository {
	return &repo{db: dbClient}
}

func (r *repo) Roles(_ context.Context) (map[string]string, error) {
	accessibleRoles := make(map[string]string)
	accessibleRoles[model.ExamplePath] = "admin"
	return accessibleRoles, nil
}
