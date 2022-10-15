package graph

import (
	"12-startups-one-month/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService   service.IUserService
	AuthService   service.IAuthService
	FileService   service.IFileService
	ViewerService service.IViewerService
	AvisService   service.IAvisService
}
