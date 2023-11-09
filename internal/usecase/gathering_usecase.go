package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
)

type gatheringUsecase struct {
	gatheringRepo model.GatheringRepository
}

func NewGatheringUsecase(gatheringRepo model.GatheringRepository) model.GatheringUsecase {
	return &gatheringUsecase{
		gatheringRepo: gatheringRepo,
	}
}

func (gu *gatheringUsecase) CreateGathering(ctx context.Context, gathering *model.Gathering) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"gathering": gathering,
	})
	err := gu.gatheringRepo.Create(ctx, gathering)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (gu *gatheringUsecase) FindGatheringByID(ctx context.Context, gatheringID int64) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"gatheringID": gatheringID,
	})
	res, err := gu.gatheringRepo.FindByID(ctx, gatheringID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		return nil, ErrRecordNotFound
	}

	return res, nil
}

func (gu *gatheringUsecase) UpdateGatheringByID(ctx context.Context, gathering *model.Gathering) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"gathering": gathering,
	})

	oldGathering, err := gu.gatheringRepo.FindByID(ctx, gathering.ID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case oldGathering == nil:
		return nil, ErrRecordNotFound
	}

	res, err := gu.gatheringRepo.UpdateByID(ctx, gathering)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (gu *gatheringUsecase) DeleteGatheringByID(ctx context.Context, gatheringID int64) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"gatheringID": gatheringID,
	})

	gathering, err := gu.gatheringRepo.FindByID(ctx, gatheringID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case gathering == nil:
		return nil, ErrRecordNotFound
	}

	res, err := gu.gatheringRepo.DeleteByID(ctx, gatheringID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}
