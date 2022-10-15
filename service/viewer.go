package service

import (
	config "12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type IViewerService interface {
	AddViewer(ctx context.Context, user_viewed string) (*model.Viewer, error)
	GetViewersByUserID(ctx context.Context, user_id string) ([]*model.Viewer, error)
}

type ViewerService struct {
	server *config.Server
}

func NewViewerService(server *config.Server) *ViewerService {
	return &ViewerService{
		server: server,
	}
}

func SqlViewerToGraphViewer(sqlViewer *db.Viewer) *model.Viewer {
	if sqlViewer == nil {
		return nil
	}
	return &model.Viewer{
		ID:              sqlViewer.ID.String(),
		CreatedAt:       sqlViewer.CreatedAt,
		UpdatedAt:       sqlViewer.UpdatedAt,
		DeletedAt:       &sqlViewer.DeletedAt.Time,
		DateViewed:      sqlViewer.DateViewed,
		ProfileIDViewed: sqlViewer.ProfilIDViewed.String(),
		UserIDViewer:    sqlViewer.UserIDViewer.String(),
	}
}

func (s *ViewerService) AddViewer(ctx context.Context, user_viewed string) (*model.Viewer, error) {
	var res *model.Viewer

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		user_ctx := s.server.GetUserContext(ctx)
		if user_ctx == nil {
			return fmt.Errorf("user context is nil")
		}
		if user_ctx.ID.String() == user_viewed {
			return fmt.Errorf("user_id and user_viewed are the same")
		}
		newView, err := q.CreateView(ctx, db.CreateViewParams{
			ProfilIDViewed: user_ctx.ID,
			UserIDViewer:   uuid.MustParse(user_viewed),
		})
		if err != nil {
			return err
		}

		// get Viewer
		viewer, err := q.GetViewByID(ctx, newView.ID)
		if err != nil {
			return err
		}

		// convert to graphql model
		res = SqlViewerToGraphViewer(&viewer)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_CREATE_VIEWER", err)
	}
	return res, nil
}

func (s *ViewerService) GetViewersByUserID(ctx context.Context, user_id string) ([]*model.Viewer, error) {
	var res []*model.Viewer

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		viewers, err := q.GetViewsByUserID(ctx, uuid.MustParse(user_id))
		if err != nil {
			return err
		}

		// remove all duplicate UserIDViewer in viewers
		viewers = utils.RemoveDuplicateViewer(viewers)

		for _, viewer := range viewers {
			res = append(res, SqlViewerToGraphViewer(&viewer))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_VIEWERS", err)
	}
	return res, nil
}
