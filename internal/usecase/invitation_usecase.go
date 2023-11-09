package usecase

import (
	"context"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/sirupsen/logrus"
)

type invitationUsecase struct {
	invitationRepo model.InvitationRepository
	memberRepo     model.MemberRepository
	gatheringRepo  model.GatheringRepository
	attendeeRepo   model.AttendeeRepository
}

func NewInvitationUsecase(invitationRepo model.InvitationRepository,
	memberRepo model.MemberRepository,
	gatheringRepo model.GatheringRepository,
	attendeeRepo model.AttendeeRepository) model.InvitationUsecase {
	return &invitationUsecase{
		invitationRepo: invitationRepo,
		memberRepo:     memberRepo,
		gatheringRepo:  gatheringRepo,
		attendeeRepo:   attendeeRepo,
	}
}

func (iu *invitationUsecase) InviteMemberToGathering(ctx context.Context, invitation *model.Invitation) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"invitation": invitation,
	})

	memberRes, err := iu.memberRepo.FindByID(ctx, invitation.MemberID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case memberRes == nil:
		return nil, ErrRecordNotFound
	}

	gatheringRes, err := iu.gatheringRepo.FindByID(ctx, invitation.GatheringID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case gatheringRes == nil:
		return nil, ErrRecordNotFound
	}

	err = iu.invitationRepo.Create(ctx, invitation)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	attendee := &model.Attendee{
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
	}

	err = iu.attendeeRepo.Create(ctx, attendee)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return iu.invitationRepo.FindByID(ctx, invitation.ID)
}

func (iu *invitationUsecase) FindInvitationByID(ctx context.Context, invitationID int64) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"invitationID": invitationID,
	})

	res, err := iu.invitationRepo.FindByID(ctx, invitationID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		return nil, ErrRecordNotFound
	}

	return res, nil
}

func (iu *invitationUsecase) UpdateInvitationByID(ctx context.Context, invitation *model.Invitation) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":        ctx,
		"invitation": invitation,
	})

	oldInvitation, err := iu.invitationRepo.FindByID(ctx, invitation.ID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case oldInvitation == nil:
		return nil, ErrRecordNotFound
	}

	memberRes, err := iu.memberRepo.FindByID(ctx, invitation.MemberID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case memberRes == nil:
		return nil, ErrRecordNotFound
	}

	gatheringRes, err := iu.gatheringRepo.FindByID(ctx, invitation.GatheringID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case gatheringRes == nil:
		return nil, ErrRecordNotFound
	}

	res, err := iu.invitationRepo.UpdateByID(ctx, invitation)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (iu *invitationUsecase) DeleteInvitationByID(ctx context.Context, invitationID int64) (*model.Invitation, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":          ctx,
		"invitationID": invitationID,
	})

	invitation, err := iu.invitationRepo.FindByID(ctx, invitationID)
	switch {
	case err != nil:
		logger.Error(err)
		return nil, err
	case invitation == nil:
		return nil, ErrRecordNotFound
	}

	res, err := iu.invitationRepo.DeleteByID(ctx, invitationID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = iu.attendeeRepo.DeleteByMemberIDAndGatheringID(ctx, invitation.MemberID, invitation.GatheringID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}
