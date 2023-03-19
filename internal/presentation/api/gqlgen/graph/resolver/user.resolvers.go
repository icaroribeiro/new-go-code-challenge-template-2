package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	presentableentity "github.com/icaroribeiro/go-code-challenge-template-2/internal/presentation/api/gqlgen/graph/presentity"
)

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*presentableentity.User, error) {
	domainUsers, err := r.UserService.WithDBTrx(nil).GetAll()
	if err != nil {
		return nil, err
	}

	users := presentableentity.Users{}
	users.FromDomain(domainUsers)

	allUsers := []*presentableentity.User{}
	for i := range users {
		allUsers = append(allUsers, &users[i])
	}

	return allUsers, nil
}
