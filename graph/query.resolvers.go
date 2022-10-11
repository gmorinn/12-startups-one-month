package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"12-startups-one-month/graph/model"
	"12-startups-one-month/graph/mypkg"
	"context"
	"fmt"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id mypkg.UUID) (*model.User, error) {
	return r.UserService.GetUser(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	return r.UserService.GetUsers(ctx, limit, offset)
}

// Views is the resolver for the views field.
func (r *queryResolver) Views(ctx context.Context, id mypkg.UUID) ([]*model.Viewer, error) {
	panic(fmt.Errorf("not implemented: Views - views"))
}

// Avis is the resolver for the avis field.
func (r *queryResolver) Avis(ctx context.Context, id mypkg.UUID) ([]*model.Avis, error) {
	panic(fmt.Errorf("not implemented: Avis - avis"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
