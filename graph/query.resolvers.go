package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"12-startups-one-month/graph/model"
	"context"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserService.GetUser(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	return r.UserService.GetUsers(ctx, limit, offset)
}

// GetViewsByUserID is the resolver for the getViewsByUserId field.
func (r *queryResolver) GetViewsByUserID(ctx context.Context, id string) ([]*model.Viewer, error) {
	return r.ViewerService.GetViewersByUserID(ctx, id)
}

// GetAvisByUserID is the resolver for the getAvisByUserId field.
func (r *queryResolver) GetAvisByUserID(ctx context.Context, id string) ([]*model.Avis, error) {
	return r.AvisService.GetAvisByUserID(ctx, id)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
