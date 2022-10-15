package service

import (
	"12-startups-one-month/config"
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"12-startups-one-month/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type IAvisService interface {
	CreateAvis(ctx context.Context, input *model.AvisInput) (*model.Avis, error)
	UpdateAvis(ctx context.Context, input *model.AvisInput) (*model.Avis, error)
	DeleteAvis(ctx context.Context, id string) (bool, error)
	GetAvisByUserID(ctx context.Context, id string) ([]*model.Avis, error)
}

type AvisService struct {
	server *config.Server
}

func NewAvisService(server *config.Server) *AvisService {
	return &AvisService{
		server: server,
	}
}

func sqlAvisToGraphAvis(sqlAvis *db.Avi) *model.Avis {
	if sqlAvis == nil {
		return nil
	}
	return &model.Avis{
		ID:           sqlAvis.ID.String(),
		CreatedAt:    sqlAvis.CreatedAt,
		UpdatedAt:    sqlAvis.UpdatedAt,
		DeletedAt:    &sqlAvis.DeletedAt.Time,
		Comment:      sqlAvis.Comment,
		Note:         int(sqlAvis.Note),
		UserIDTarget: sqlAvis.UserIDTarget.String(),
		UserIDWriter: sqlAvis.UserIDWriter.String(),
	}
}

func (s *AvisService) CreateAvis(ctx context.Context, input *model.AvisInput) (*model.Avis, error) {
	var res *model.Avis

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// create Avis
		tmp_avis, err := q.CreateAvis(ctx, db.CreateAvisParams{
			UserIDTarget: uuid.MustParse(input.UserIDTarget),
			UserIDWriter: uuid.MustParse(input.UserIDWriter),
			Comment:      input.Comment,
			Note:         int32(input.Note),
		})
		if err != nil {
			return err
		}

		// get Avis
		new_avis, err := q.GetAvisByID(ctx, tmp_avis.ID)
		if err != nil {
			return err
		}

		// convert to graphql model
		res = sqlAvisToGraphAvis(&new_avis)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_CREATE_AVIS", err)
	}
	return res, nil
}

func (s *AvisService) UpdateAvis(ctx context.Context, input *model.AvisInput) (*model.Avis, error) {
	var res *model.Avis

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// update Avis
		if err := q.UpdateAvis(ctx, db.UpdateAvisParams{
			// ID:      uuid.MustParse(input.),
			Comment: input.Comment,
			Note:    int32(input.Note),
		}); err != nil {
			return err
		}

		// // get updated_avis
		// updated_avis, err := q.GetAvisByID(ctx, uuid.MustParse(input.ID))
		// if err != nil {
		// 	return err
		// }

		// convert to graphql model
		// res = sqlAvisToGraphAvis(&updated_avis)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_AVIS", err)
	}
	return res, nil
}

func (s *AvisService) DeleteAvis(ctx context.Context, id string) (*bool, error) {
	var res bool = false

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if Avis exists
		isAvis, err := q.CheckAvisByID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}
		if !isAvis {
			return utils.ErrorResponse("AVIS_NOT_FOUND", errors.New("avis not found"))
		}
		if err := q.DeleteAvisByID(ctx, uuid.MustParse(id)); err != nil {
			res = false
			return err
		}
		res = true
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_DELETE_AVIS", err)
	}
	return &res, nil
}

func (s *AvisService) GetAvisByUserID(ctx context.Context, id string) ([]*model.Avis, error) {
	var res []*model.Avis

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// get Avis
		avis, err := q.GetAvisByUserID(ctx, uuid.MustParse(id))
		if err != nil {
			return err
		}

		// convert to graphql model
		for _, a := range avis {
			res = append(res, sqlAvisToGraphAvis(&a))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_AVIS", err)
	}
	return res, nil
}
