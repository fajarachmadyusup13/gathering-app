package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
)

type memberUsecase struct {
	memberRepo model.MemberRepository
}

func NewMemberUsecase(memberRepo model.MemberRepository) model.MemberUsecase {
	return &memberUsecase{
		memberRepo: memberRepo,
	}
}

func (mu *memberUsecase) Register(ctx context.Context, member *model.Member) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"member": member,
	})
	err := mu.memberRepo.Create(ctx, member)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (mu *memberUsecase) FindMemberByID(ctx context.Context, memberID int64) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"memberID": memberID,
	})
	res, err := mu.memberRepo.FindByID(ctx, memberID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		return nil, ErrRecordNotFound
	}

	return res, nil
}

func (mu *memberUsecase) UpdateMemberByID(ctx context.Context, member *model.Member) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"member": member,
	})

	oldMember, err := mu.memberRepo.FindByID(ctx, member.ID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case oldMember == nil:
		return nil, ErrRecordNotFound
	}

	res, err := mu.memberRepo.UpdateByID(ctx, member)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (mu *memberUsecase) DeleteMemberByID(ctx context.Context, memberID int64) (*model.Member, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":      ctx,
		"memberID": memberID,
	})

	member, err := mu.memberRepo.FindByID(ctx, memberID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case member == nil:
		return nil, ErrRecordNotFound
	}

	res, err := mu.memberRepo.DeleteByID(ctx, memberID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}
