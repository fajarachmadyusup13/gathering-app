package repository

import (
	"context"
	"errors"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) model.MemberRepository {
	return &memberRepository{
		db: db,
	}
}

func (m *memberRepository) Create(ctx context.Context, member *model.Member) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"member": member,
	})

	tx := m.db.WithContext(ctx).Begin()
	err := tx.Create(member).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *memberRepository) FindByID(ctx context.Context, memberID int64) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"memberID": memberID,
	})

	var member model.Member
	err := m.db.WithContext(ctx).Take(&member, memberID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &member, err
}

func (m *memberRepository) FindAll(ctx context.Context) ([]int64, error) {
	return nil, nil
}

func (m *memberRepository) UpdateByID(ctx context.Context, member *model.Member) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"member": member,
	})

	oldMember, err := m.FindByID(ctx, member.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if oldMember == nil {
		return nil, nil
	}

	tx := m.db.WithContext(ctx).Begin()
	err = tx.Model(member).Omit(member.ImmutableColumns()...).Save(member).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return m.FindByID(ctx, member.ID)
}

func (m *memberRepository) DeleteByID(ctx context.Context, memberID int64) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"memberID": memberID,
	})

	var member model.Member
	tx := m.db.WithContext(ctx).Begin()
	err := tx.Find(&member, memberID).Delete(&member).Error
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

	err = m.db.WithContext(ctx).Unscoped().Find(&member, memberID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &member, nil
}
