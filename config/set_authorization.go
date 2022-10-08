package config

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	sqlc "12-startups-one-month/internal"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/google/uuid"
)

type JwtContent struct {
	ID   uuid.UUID `json:"id"`
	Role db.Role   `json:"role"`
}

// storeRefresh store refres_token into database
func (server *Server) StoreRefresh(ctx context.Context, token string, exp time.Time, userID uuid.UUID) error {
	return server.Store.CreateRefreshToken(ctx, sqlc.CreateRefreshTokenParams{
		Token:     token,
		ExpirOn:   exp,
		UserID:    userID,
		Ip:        "A temporary IP",
		UserAgent: "A temporary UserAgent",
	})
}

// generate access token, refresh token and expiry time for user based on the id and role
func (server *Server) GenerateJwtToken(ID uuid.UUID, role string) (string, string, time.Time, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   ID.String(),
		"role": role,
		"exp":  time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(server.Config.Security.AccessTokenDuration))).Unix(),
	})
	t, err := accessToken.SignedString([]byte(server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_ACCESS_JWT %v", err)
	}
	expt := time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(server.Config.Security.RefreshTokenDuration)))
	exp := expt.Unix()

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   ID.String(),
		"role": role,
		"exp":  exp,
	})
	r, err := refreshToken.SignedString([]byte(server.Config.Security.Secret))
	if err != nil {
		return "", "", time.Now(), fmt.Errorf("ERROR_GENERATE_REFRESH_JWT %v", err)
	}
	return t, r, expt, nil
}

// hasJWT middleware validate jwt from headers `jwtToken` header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		extractor := request.HeaderExtractor{"jwtToken"}
		filter := func(t string) (string, error) {
			if len(t) > 6 && strings.ToUpper(t[0:7]) == "BEARER " {
				return t[7:], nil
			}
			return t, nil
		}
		token, err := request.ParseFromRequest(c.Request, &request.PostExtractionFilter{Extractor: extractor, Filter: filter}, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(os.Getenv("API_SECRET")))
			return b, nil
		})

		if err != nil {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "jwt", token)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (server *Server) GetUserContext(ctx context.Context) *JwtContent {
	var res *JwtContent = nil
	raw := ctx.Value("jwt")
	if raw == nil {
		return nil
	}
	u, ok := raw.(*jwt.Token)
	if !ok {
		return nil
	}
	claims := u.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(string)
	if !ok {
		return nil
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil
	}
	res = &JwtContent{
		ID:   uuid.MustParse(id),
		Role: db.Role(role),
	}
	return res
}

func (server *Server) JwtAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userCtx := server.GetUserContext(ctx)

	if userCtx == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}
	return next(ctx)
}

func (server *Server) HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.UserType) (interface{}, error) {
	userCtx := server.GetUserContext(ctx)

	if userCtx == nil || string(userCtx.Role) != strings.ToLower(role.String()) {
		return nil, &gqlerror.Error{
			Message: "Can't access this resource",
		}
	}

	return next(ctx)
}
