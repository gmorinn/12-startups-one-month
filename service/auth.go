package service

import (
	config "12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	"12-startups-one-month/graph/mypkg"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type IAuthService interface {
	Signin(ctx context.Context, input *model.SigninInput) (*model.JWTResponse, error)
	Signup(ctx context.Context, input *model.SignupInput) (*model.JWTResponse, error)
	RefreshToken(ctx context.Context, refreshToken *mypkg.JWT) (*model.JWTResponse, error)
}

type AuthService struct {
	server *config.Server
}

func NewAuthService(server *config.Server) *AuthService {
	return &AuthService{
		server: server,
	}
}

func (s *AuthService) Signin(ctx context.Context, input *model.SigninInput) (*model.JWTResponse, error) {
	var response model.JWTResponse

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.LoginUserParams{
			Email: string(input.Email),
			Crypt: input.Password,
		}
		user, err := s.server.Store.LoginUser(ctx, arg)
		if err != nil {
			return utils.ErrorResponse("ERROR_LOGIN_USER", err)
		}
		t, r, expt, err := s.server.GenerateJwtToken(user.ID, utils.ConvertRoleToString(user.Role))
		if err != nil {
			return utils.ErrorResponse("ERROR_TOKEN", err)
		}
		if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
			return utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
		}
		response = model.JWTResponse{
			AccessToken:  mypkg.JWT(t),
			RefreshToken: mypkg.JWT(r),
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_USERS", err)
	}
	return &response, nil
}

func (s *AuthService) Signup(ctx context.Context, input *model.SignupInput) (*model.JWTResponse, error) {
	var user db.User

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if input.Password != input.ConfirmPassword {
			return fmt.Errorf("password and confirm password not match")
		}
		isExist, err := q.CheckEmailExist(ctx, string(input.Email))
		if err != nil {
			return utils.ErrorResponse("ERROR_GET_MAIL", err)
		}
		if isExist {
			return utils.ErrorResponse("EMAIL_ALREADY_EXIST", fmt.Errorf("email already exist"))
		}
		arg := db.SignupParams{
			Email: string(input.Email),
			Crypt: input.Password,
		}
		user, err = q.Signup(ctx, arg)
		if err != nil {
			return utils.ErrorResponse("ERROR_CREATE_USER", err)
		}
		return nil
	})
	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_USERS", err)
	}
	t, r, expt, err := s.server.GenerateJwtToken(user.ID, utils.ConvertRoleToString(user.Role))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_TOKEN", err)
	}
	if err := s.server.StoreRefresh(ctx, r, expt, user.ID); err != nil {
		return nil, utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
	}
	return &model.JWTResponse{
		AccessToken:  mypkg.JWT(t),
		RefreshToken: mypkg.JWT(r),
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken *mypkg.JWT) (*model.JWTResponse, error) {
	token, err := jwt.Parse(string(*refreshToken), func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return nil, utils.ErrorResponse("TOKEN_ERROR", err)
	}
	if !token.Valid {
		return nil, utils.ErrorResponse("TOKEN_IS_NOT_VALID", fmt.Errorf("invalid format token"))
	}
	refresh, err := s.server.Store.GetRefreshToken(ctx, string(*refreshToken))
	if err != nil {
		return nil, utils.ErrorResponse("FIND_REFRESH_TOKEN", err)
	}
	t, r, expt, err := s.server.GenerateJwtToken(refresh.UserID, utils.ConvertRoleToString(refresh.UserRole))
	if err != nil {
		return nil, utils.ErrorResponse("ERROR_TOKEN", err)
	}
	if err := s.server.StoreRefresh(ctx, r, expt, refresh.UserID); err != nil {
		return nil, utils.ErrorResponse("ERROR_REFRESH_TOKEN", err)
	}
	response := model.JWTResponse{
		AccessToken:  mypkg.JWT(t),
		RefreshToken: mypkg.JWT(r),
	}
	return &response, nil
}
