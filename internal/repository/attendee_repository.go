package repository

import (
	"context"
	"errors"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type attendeeRepository struct {
	db *gorm.DB
}

func NewAttendeeRepository(db *gorm.DB) model.AttendeeRepository {
	return &attendeeRepository{
		db: db,
	}
}

func (a *attendeeRepository) Create(ctx context.Context, attendee *model.Attendee) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"attendee": attendee,
	})

	tx := a.db.WithContext(ctx).Begin()
	err := tx.Create(attendee).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (a *attendeeRepository) FindByMemberID(ctx context.Context, memberID int64) ([]*model.Attendee, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"memberID": memberID,
	})

	var attendees []*model.Attendee

	err := a.db.Where(&model.Attendee{MemberID: memberID}).Find(&attendees).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}

	return attendees, nil
}

func (a *attendeeRepository) FindByGatheringID(ctx context.Context, gatheringID int64) ([]*model.Attendee, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"gatheringID": gatheringID,
	})

	var attendees []*model.Attendee

	err := a.db.Where(&model.Attendee{GatheringID: gatheringID}).Find(&attendees).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}

	return attendees, nil
}

func (a *attendeeRepository) DeleteByMemberIDAndGatheringID(ctx context.Context, memberID int64, gatheringID int64) (*model.Attendee, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         ctx,
		"memberID":    memberID,
		"gatheringID": gatheringID,
	})

	var attendee model.Attendee

	tx := a.db.WithContext(ctx).Begin()
	err := tx.Where(&model.Attendee{MemberID: memberID, GatheringID: gatheringID}).
		Delete(&model.Attendee{}).Error
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

	err = a.db.WithContext(ctx).Unscoped().
		Where(&model.Attendee{MemberID: memberID, GatheringID: gatheringID}).
		Find(&attendee).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &attendee, nil
}
