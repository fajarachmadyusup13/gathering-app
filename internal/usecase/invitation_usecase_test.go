package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/fajarachmadyusup13/gathering-app/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInviteMemberToGatheringUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 444,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	member := &model.Member{
		ID:        invitation.MemberID,
		FirstName: "first",
		LastName:  "last",
		Email:     "first@last.com",
	}

	gathering := &model.Gathering{
		ID:          invitation.GatheringID,
		Creator:     111,
		Type:        model.WithExpirationForInvitations,
		ScheduledAt: nil,
		Name:        "gathering",
		Location:    "locc",
	}

	attendee := &model.Attendee{
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
	}

	t.Run("success", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockAttendeeRepo := mock.NewMockAttendeeRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().Create(ctx, invitation).Times(1).Return(nil)
		mockAttendeeRepo.EXPECT().Create(ctx, attendee).Times(1).Return(nil)
		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
			attendeeRepo:   mockAttendeeRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.NotNil(t, res)
		assert.NoError(t, err)

	})

	t.Run("failed, error find member by id", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)

	})

	t.Run("failed, error member not found", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())

	})

	t.Run("failed, error find gathering by id", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			memberRepo:    mockMemberRepo,
			gatheringRepo: mockGatheringRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)

	})

	t.Run("failed, error gathering not found", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			memberRepo:    mockMemberRepo,
			gatheringRepo: mockGatheringRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error create", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().Create(ctx, invitation).Times(1).Return(errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error create attendee", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockAttendeeRepo := mock.NewMockAttendeeRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().Create(ctx, invitation).Times(1).Return(nil)
		mockAttendeeRepo.EXPECT().Create(ctx, attendee).Times(1).Return(errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
			attendeeRepo:   mockAttendeeRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error return invitation", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockAttendeeRepo := mock.NewMockAttendeeRepository(ctrl)

		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().Create(ctx, invitation).Times(1).Return(nil)
		mockAttendeeRepo.EXPECT().Create(ctx, attendee).Times(1).Return(nil)
		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
			attendeeRepo:   mockAttendeeRepo,
		}

		res, err := invitationUsecase.InviteMemberToGathering(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)

	})
}

func TestFindInvitationByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 444,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.FindInvitationByID(ctx, invitation.ID)
		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("failed, error find invitation by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.FindInvitationByID(ctx, invitation.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error invitation not found", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.FindInvitationByID(ctx, invitation.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})
}

func TestUpdateInvitationByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 444,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	member := &model.Member{
		ID:        invitation.MemberID,
		FirstName: "first",
		LastName:  "last",
		Email:     "first@last.com",
	}

	gathering := &model.Gathering{
		ID:          invitation.GatheringID,
		Creator:     111,
		Type:        model.WithExpirationForInvitations,
		ScheduledAt: nil,
		Name:        "gathering",
		Location:    "locc",
	}

	t.Run("success", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().UpdateByID(ctx, invitation).Times(1).Return(invitation, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("failed, error find invitation by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error invitation not found", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error find member by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error member not found", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error find gathering by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, error gathering not found", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error update invitation by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockMemberRepo.EXPECT().FindByID(ctx, invitation.MemberID).Times(1).Return(member, nil)
		mockGatheringRepo.EXPECT().FindByID(ctx, invitation.GatheringID).Times(1).Return(gathering, nil)
		mockInvitationRepo.EXPECT().UpdateByID(ctx, invitation).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			memberRepo:     mockMemberRepo,
			gatheringRepo:  mockGatheringRepo,
		}

		res, err := invitationUsecase.UpdateInvitationByID(ctx, invitation)
		assert.Nil(t, res)
		assert.Error(t, err)
	})
}

func TestDeleteInvitationByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 444,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	attendee := &model.Attendee{
		MemberID:    invitation.MemberID,
		GatheringID: invitation.GatheringID,
	}

	t.Run("success", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)
		mockAttendeeRepo := mock.NewMockAttendeeRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockInvitationRepo.EXPECT().DeleteByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockAttendeeRepo.EXPECT().DeleteByMemberIDAndGatheringID(ctx, invitation.MemberID, invitation.GatheringID).Times(1).Return(attendee, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
			attendeeRepo:   mockAttendeeRepo,
		}

		res, err := invitationUsecase.DeleteInvitationByID(ctx, invitation.ID)
		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("failed, error find invitation by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.DeleteInvitationByID(ctx, invitation.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, invitation not found", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(nil, nil)

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.DeleteInvitationByID(ctx, invitation.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error delete invitation by id", func(t *testing.T) {
		mockInvitationRepo := mock.NewMockInvitationRepository(ctrl)

		mockInvitationRepo.EXPECT().FindByID(ctx, invitation.ID).Times(1).Return(invitation, nil)
		mockInvitationRepo.EXPECT().DeleteByID(ctx, invitation.ID).Times(1).Return(nil, errors.New("error"))

		invitationUsecase := invitationUsecase{
			invitationRepo: mockInvitationRepo,
		}

		res, err := invitationUsecase.DeleteInvitationByID(ctx, invitation.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

}
