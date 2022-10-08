package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"12-startups-one-month/graph/model"
	"12-startups-one-month/graph/mypkg"
	"context"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id mypkg.UUID) (*model.User, error) {
	return r.UserService.GetUser(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	return r.UserService.GetUsers(ctx, limit, offset)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
