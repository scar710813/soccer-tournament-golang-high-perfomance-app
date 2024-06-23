package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

/*
Summary: Create a new team based on the provided data and store it in the database
Method: Create
Params: context.Context, *domain.Team
Return: error
*/

func (s Service) Create(ctx context.Context, team *domain.Team) (err error) {
	err = s.Repo.Insert(ctx, team)
	if err != nil {
		return domain.ManageError(err, "Error creating team")
	}
	return nil
}
