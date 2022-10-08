package service

import (
	config "12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	"12-startups-one-month/graph/mypkg"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type IUserService interface {
	GetUser(ctx context.Context, id mypkg.UUID) (*model.User, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]*model.User, error)
	UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error)
	DeleteUser(ctx context.Context, id mypkg.UUID) (*bool, error)
	UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.User, error)
}

type UserService struct {
	server *config.Server
}

func NewUserService(server *config.Server) *UserService {
	return &UserService{
		server: server,
	}
}

func SqlUserToGraphUser(sqlUser *db.User) *model.User {
	if sqlUser == nil {
		return nil
	}
	return &model.User{
		ID:        mypkg.UUID(sqlUser.ID.String()),
		Firstname: sqlUser.Firstname.String,
		Lastname:  sqlUser.Lastname.String,
		Email:     mypkg.Email(sqlUser.Email),
		Role:      model.UserType(sqlUser.Role),
		CreatedAt: sqlUser.CreatedAt,
		UpdatedAt: sqlUser.UpdatedAt,
		DeletedAt: &sqlUser.DeletedAt.Time,
	}
}

func (s *UserService) GetUser(ctx context.Context, id mypkg.UUID) (*model.User, error) {
	var res *model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		res = SqlUserToGraphUser(&sqlUser)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_USER", err)
	}
	return res, err
}

func (s *UserService) GetUsers(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	var res []*model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if limit is valid
		if limit < 0 {
			return utils.ErrorResponse("INVALID_LIMIT", errors.New("limit must be greater than 0"))
		}
		// check if offset is valid
		if offset < 0 {
			return utils.ErrorResponse("INVALID_OFFSET", errors.New("offset must be greater than 0"))
		}
		sqlUsers, err := q.GetAllUser(ctx, db.GetAllUserParams{
			Limit:        int32(limit),
			Offset:       int32(offset),
			FirstnameAsc: true,
		})
		if err != nil {
			return err
		}
		for _, sqlUser := range sqlUsers {
			res = append(res, SqlUserToGraphUser(&sqlUser))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_USERS", err)
	}
	return res, err
}

func (s *UserService) UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error) {
	var res *model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.UpdateUser(ctx, db.UpdateUserParams{
			ID:        uuid.MustParse(string(input.ID)),
			Firstname: utils.NullS(input.Firstname),
			Lastname:  utils.NullS(input.Lastname),
			Email:     string(input.Email),
		}); err != nil {
			return err
		}

		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(string(input.ID)))
		if err != nil {
			return err
		}
		res = SqlUserToGraphUser(&sqlUser)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_USER", err)
	}
	return res, err
}

func (s *UserService) DeleteUser(ctx context.Context, id mypkg.UUID) (*bool, error) {
	var res bool = false

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if user exists
		isUser, err := q.CheckUserByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		if !isUser {
			return utils.ErrorResponse("USER_NOT_FOUND", errors.New("user not found"))
		}
		if err := q.DeleteUserByID(ctx, uuid.MustParse(string(id))); err != nil {
			return err
		}
		res = true
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_DELETE_USER", err)
	}
	return &res, err
}

func (s *UserService) UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.User, error) {
	var res *model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.UpdateRole(ctx, db.UpdateRoleParams{
			ID:   uuid.MustParse(string(id)),
			Role: db.Role(strings.ToLower(string(role))),
		}); err != nil {
			return err
		}

		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		res = SqlUserToGraphUser(&sqlUser)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_USER_ROLE", err)
	}
	return res, err
}
