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

func TestCreateGatheringUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	gathering := &model.Gathering{
		ID:        123,
		Creator:   321,
		Type:      model.WithFixedNumberOfAttendees,
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().Create(ctx, gathering).Times(1).Return(nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		err := gatheringUsecase.CreateGathering(ctx, gathering)

		assert.NoError(t, err)
	})

	t.Run("error create member", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().Create(ctx, gathering).Times(1).Return(errors.New("error"))

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		err := gatheringUsecase.CreateGathering(ctx, gathering)

		assert.Error(t, err)
	})
}

func TestFindGatheringByIDUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	gathering := &model.Gathering{
		ID:        123,
		Creator:   321,
		Type:      model.WithFixedNumberOfAttendees,
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.FindGatheringByID(ctx, gathering.ID)
		assert.NotNil(t, res)
		assert.Equal(t, gathering.ID, res.ID)
		assert.NoError(t, err)
	})

	t.Run("failed, not found", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.FindGatheringByID(ctx, gathering.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error from repo", func(t *testing.T) {
		errorRepo := errors.New("error from repo")
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, errorRepo)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.FindGatheringByID(ctx, gathering.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, errorRepo.Error())
	})
}

func TestUpdateGatheringByIDUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	gathering := &model.Gathering{
		ID:        123,
		Creator:   321,
		Type:      model.WithFixedNumberOfAttendees,
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
		Name:      "aaa",
	}

	t.Run("success", func(t *testing.T) {
		gathering.Name = "xxx"

		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		mockGatheringRepo.EXPECT().UpdateByID(ctx, gathering).Times(1).Return(gathering, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.UpdateGatheringByID(ctx, gathering)
		assert.NotNil(t, res)
		assert.NoError(t, err)
		assert.Equal(t, "xxx", res.Name)
	})

	t.Run("failed, error find by ID", func(t *testing.T) {
		gathering.Name = "xxx"

		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, errors.New("error"))

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.UpdateGatheringByID(ctx, gathering)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, ID not found", func(t *testing.T) {
		gathering.Name = "xxx"

		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.UpdateGatheringByID(ctx, gathering)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error update", func(t *testing.T) {
		gathering.Name = "xxx"
		errorUpdate := errors.New("error update")

		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		mockGatheringRepo.EXPECT().UpdateByID(ctx, gathering).Times(1).Return(nil, errorUpdate)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.UpdateGatheringByID(ctx, gathering)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, errorUpdate.Error())
	})
}

func TestDeleteGatheringByIDUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	gathering := &model.Gathering{
		ID:        123,
		Creator:   321,
		Type:      model.WithFixedNumberOfAttendees,
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
		Name:      "aaa",
	}

	t.Run("success", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		mockGatheringRepo.EXPECT().DeleteByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.DeleteGatheringByID(ctx, gathering.ID)
		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("failed, error find by ID", func(t *testing.T) {
		mockGatherignRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatherignRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, errors.New("error"))

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatherignRepo,
		}

		res, err := gatheringUsecase.DeleteGatheringByID(ctx, gathering.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, ID not found", func(t *testing.T) {
		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(nil, nil)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.DeleteGatheringByID(ctx, gathering.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error delete", func(t *testing.T) {
		errorDelete := errors.New("error delete")

		mockGatheringRepo := mock.NewMockGatheringRepository(ctrl)
		mockGatheringRepo.EXPECT().FindByID(ctx, gathering.ID).Times(1).Return(gathering, nil)

		mockGatheringRepo.EXPECT().DeleteByID(ctx, gathering.ID).Times(1).Return(nil, errorDelete)

		gatheringUsecase := gatheringUsecase{
			gatheringRepo: mockGatheringRepo,
		}

		res, err := gatheringUsecase.DeleteGatheringByID(ctx, gathering.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, errorDelete.Error())
	})
}
