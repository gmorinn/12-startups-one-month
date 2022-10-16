package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"12-startups-one-month/graph/model"
	"12-startups-one-month/graph/mypkg"
	"context"
)

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, input model.SigninInput) (*model.JWTResponse, error) {
	return r.Resolver.AuthService.Signin(ctx, &input)
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (*model.JWTResponse, error) {
	return r.AuthService.Signup(ctx, &input)
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, refreshToken mypkg.JWT) (*model.JWTResponse, error) {
	return r.AuthService.RefreshToken(ctx, &refreshToken)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserProfileInput) (*model.User, error) {
	return r.UserService.UpdateUser(ctx, &input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return r.UserService.DeleteUser(ctx, id)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, role []model.UserType, id string) (*model.User, error) {
	return r.UserService.UpdateRole(ctx, role, id)
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file model.UploadInput) (*model.UploadResponse, error) {
	return r.FileService.UploadSingleFile(ctx, &file)
}

// AddViewer is the resolver for the addViewer field.
func (r *mutationResolver) AddViewer(ctx context.Context, userViewed string) (*model.Viewer, error) {
	return r.ViewerService.AddViewer(ctx, userViewed)
}

// CreateAvis is the resolver for the createAvis field.
func (r *mutationResolver) CreateAvis(ctx context.Context, input model.AvisCreateInput) (*model.Avis, error) {
	return r.AvisService.CreateAvis(ctx, &input)
}

// UpdateAvis is the resolver for the updateAvis field.
func (r *mutationResolver) UpdateAvis(ctx context.Context, input model.AvisUpdateInput) (*model.Avis, error) {
	return r.AvisService.UpdateAvis(ctx, &input)
}

// DeleteAvis is the resolver for the deleteAvis field.
func (r *mutationResolver) DeleteAvis(ctx context.Context, id string) (bool, error) {
	return r.AvisService.DeleteAvis(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
