package repository

import (
	"context"
	"errors"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type gatheringRepository struct {
	db *gorm.DB
}

func NewGatheringRepository(db *gorm.DB) model.GatheringRepository {
	return &gatheringRepository{
		db: db,
	}
}

func (g *gatheringRepository) Create(ctx context.Context, gathering *model.Gathering) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"gathering": gathering,
	})

	tx := g.db.WithContext(ctx).Begin()
	err := tx.Create(gathering).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (g *gatheringRepository) FindByID(ctx context.Context, gatheringID int64) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"gatheringID": gatheringID,
	})

	var gathering model.Gathering
	err := g.db.WithContext(ctx).Take(&gathering, gatheringID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &gathering, err
}

func (g *gatheringRepository) UpdateByID(ctx context.Context, gathering *model.Gathering) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       ctx,
		"gathering": gathering,
	})

	oldGathering, err := g.FindByID(ctx, gathering.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if oldGathering == nil {
		return nil, nil
	}

	tx := g.db.WithContext(ctx).Begin()
	err = tx.Model(gathering).Omit(gathering.ImmutableColumns()...).Save(gathering).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return g.FindByID(ctx, gathering.ID)
}

func (g *gatheringRepository) DeleteByID(ctx context.Context, gatheringID int64) (*model.Gathering, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"gatheringID": gatheringID,
	})

	var gathering model.Gathering
	tx := g.db.WithContext(ctx).Begin()
	err := tx.Find(&gathering, gatheringID).Delete(&gathering).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = g.db.WithContext(ctx).Unscoped().Find(&gathering, gatheringID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &gathering, nil
}
