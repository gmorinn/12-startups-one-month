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
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return r.UserService.UpdateUser(ctx, &input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id mypkg.UUID) (*bool, error) {
	return r.UserService.DeleteUser(ctx, id)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.User, error) {
	return r.UserService.UpdateRole(ctx, role, id)
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file model.UploadInput) (*model.UploadResponse, error) {
	return r.FileService.UploadSingleFile(ctx, &file)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
