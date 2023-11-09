package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type invitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) model.InvitationRepository {
	return &invitationRepository{
		db: db,
	}
}

func (i *invitationRepository) Create(ctx context.Context, invitation *model.Invitation) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"invitation": invitation,
	})

	tx := i.db.WithContext(ctx).Begin()
	err := tx.Create(invitation).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (i *invitationRepository) FindByID(ctx context.Context, invitationID int64) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"invitationID": invitationID,
	})

	var invitation model.Invitation
	err := i.db.WithContext(ctx).Take(&invitation, invitationID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &invitation, err
}

func (i *invitationRepository) UpdateByID(ctx context.Context, invitation *model.Invitation) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"invitation": invitation,
	})

	oldInvitation, err := i.FindByID(ctx, invitation.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if oldInvitation == nil {
		return nil, nil
	}

	tx := i.db.WithContext(ctx).Begin()
	err = tx.Model(invitation).Omit(invitation.ImmutableColumns()...).Save(invitation).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		fmt.Println("ASASASAS")
		logger.Error(err)
		return nil, err
	}

	return i.FindByID(ctx, invitation.ID)
}

func (i *invitationRepository) DeleteByID(ctx context.Context, invitationID int64) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"invitationID": invitationID,
	})

	var invitation model.Invitation
	tx := i.db.WithContext(ctx).Begin()
	err := tx.Find(&invitation, invitationID).Delete(&invitation).Error
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

	err = i.db.WithContext(ctx).Unscoped().Find(&invitation, invitationID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &invitation, nil
}
