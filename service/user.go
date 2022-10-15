package service

import (
	config "12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type IUserService interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]*model.User, error)
	UpdateUser(ctx context.Context, input *model.UpdateUserProfileInput) (*model.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
	UpdateRole(ctx context.Context, role []model.UserType, id string) (*model.User, error)
}

type UserService struct {
	server *config.Server
}

func NewUserService(server *config.Server) *UserService {
	return &UserService{
		server: server,
	}
}

func sqlUserToGraphUser(sqlUser *db.User) *model.User {
	if sqlUser == nil {
		return nil
	}
	sexe := model.SexeType(strings.ToUpper(string(sqlUser.Sexe)))
	formule := model.FormuleType(sqlUser.Formule)
	return &model.User{
		ID:        sqlUser.ID.String(),
		CreatedAt: sqlUser.CreatedAt,
		UpdatedAt: sqlUser.UpdatedAt,
		DeletedAt: &sqlUser.DeletedAt.Time,
		Email:     sqlUser.Email,
		Firstname: utils.NullSToPointeurS(sqlUser.Firstname),
		Lastname:  utils.NullSToPointeurS(sqlUser.Lastname),
		Role:      utils.ConvertRoleToUserType(sqlUser.Role),
		Age:       utils.NullI32ToPointeurI(sqlUser.Age),
		Sexe:      &sexe,
		// Goals: sqlUser.Goals,
		IdealPartner:   utils.NullSToPointeurS(sqlUser.IdealPartners),
		ProfilePicture: utils.NullSToPointeurS(sqlUser.ProfilePicture),
		City:           utils.NullSToPointeurS(sqlUser.City),
		Ask:            int(sqlUser.Ask),
		Badge:          sqlUser.Badge,
		Formule:        &formule,
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	var res *model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}
		res = sqlUserToGraphUser(&sqlUser)
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
			res = append(res, sqlUserToGraphUser(&sqlUser))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_USERS", err)
	}
	return res, err
}

func (s *UserService) UpdateUser(ctx context.Context, input *model.UpdateUserProfileInput) (*model.User, error) {
	var res *model.User = nil

	// get user from context
	user := s.server.GetUserContext(ctx)
	if user.ID != uuid.MustParse(input.ID) {
		return nil, utils.ErrorResponse("UPDATE_USER", errors.New("user is not allowed to update this user"))
	}

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var sexe db.Sexe = db.SexeNone
		if input.Sexe != nil {
			sexe = db.Sexe(strings.ToLower(string(*input.Sexe)))
		}

		var age sql.NullInt32 = sql.NullInt32{}
		if input.Age != nil {
			age = sql.NullInt32{
				Int32: int32(*input.Age),
				Valid: true,
			}
		}

		if err := q.UpdateUser(ctx, db.UpdateUserParams{
			ID:        uuid.MustParse(input.ID),
			Email:     input.Email,
			Firstname: utils.PointeurS(input.Firstname),
			Lastname:  utils.PointeurS(input.Lastname),
			Age:       age,
			Sexe:      sexe,
			// Goals: sqlUser.Goals,
			IdealPartners:  utils.PointeurS(input.IdealPartner),
			ProfilePicture: utils.PointeurS(input.ProfilePicture),
			City:           utils.PointeurS(input.City),
		}); err != nil {
			return err
		}

		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(input.ID))
		if err != nil {
			return err
		}
		res = sqlUserToGraphUser(&sqlUser)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_USER", err)
	}
	return res, err
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (bool, error) {
	var res bool = false

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if user exists
		isUser, err := q.CheckUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}
		if !isUser {
			return utils.ErrorResponse("USER_NOT_FOUND", errors.New("user not found"))
		}
		if err := q.DeleteUserByID(ctx, uuid.MustParse(id)); err != nil {
			return err
		}
		res = true
		return nil
	})

	if err != nil {
		return res, utils.ErrorResponse("TX_DELETE_USER", err)
	}
	return res, err
}

func (s *UserService) UpdateRole(ctx context.Context, role []model.UserType, id string) (*model.User, error) {
	var res *model.User = nil

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		newRoles := []db.Role{}
		for _, v := range role {
			newRoles = append(newRoles, db.Role(strings.ToLower(string(v))))
		}
		if err := q.UpdateRole(ctx, db.UpdateRoleParams{
			ID:   uuid.MustParse(id),
			Role: newRoles,
		}); err != nil {
			return err
		}

		sqlUser, err := q.GetUserByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}
		res = sqlUserToGraphUser(&sqlUser)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_USER_ROLE", err)
	}
	return res, err
}
