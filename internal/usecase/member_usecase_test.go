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

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().Create(ctx, member).Times(1).Return(nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		err := memberUsecase.Register(ctx, member)

		assert.NoError(t, err)
	})

	t.Run("error create member", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().Create(ctx, member).Times(1).Return(errors.New("error"))

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		err := memberUsecase.Register(ctx, member)

		assert.Error(t, err)
	})
}

func TestFindMemberByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(member, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.FindMemberByID(ctx, member.ID)
		assert.NotNil(t, res)
		assert.Equal(t, member.ID, res.ID)
		assert.NoError(t, err)
	})

	t.Run("failed, not found", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.FindMemberByID(ctx, member.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error from repo", func(t *testing.T) {
		errorRepo := errors.New("error from repo")
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, errorRepo)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.FindMemberByID(ctx, member.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, errorRepo.Error())
	})
}

func TestUpdateMemberByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		member.LastName = "xxx"

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(member, nil)

		mockMemberRepo.EXPECT().UpdateByID(ctx, member).Times(1).Return(member, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.UpdateMemberByID(ctx, member)
		assert.NotNil(t, res)
		assert.NoError(t, err)
		assert.Equal(t, "xxx", res.LastName)
	})

	t.Run("failed, error find by ID", func(t *testing.T) {
		member.LastName = "xxx"

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, errors.New("error"))

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.UpdateMemberByID(ctx, member)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, ID not found", func(t *testing.T) {
		member.LastName = "xxx"

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.UpdateMemberByID(ctx, member)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error update", func(t *testing.T) {
		member.LastName = "xxx"
		errorUpdate := errors.New("error update")

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(member, nil)

		mockMemberRepo.EXPECT().UpdateByID(ctx, member).Times(1).Return(nil, errorUpdate)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.UpdateMemberByID(ctx, member)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, errorUpdate.Error())
	})
}

func TestDeleteMemberByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()

	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
		DeletedAt: gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(member, nil)

		mockMemberRepo.EXPECT().DeleteByID(ctx, member.ID).Times(1).Return(member, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.DeleteMemberByID(ctx, member.ID)
		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("failed, error find by ID", func(t *testing.T) {
		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, errors.New("error"))

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.DeleteMemberByID(ctx, member.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("failed, ID not found", func(t *testing.T) {
		member.LastName = "xxx"

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(nil, nil)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.DeleteMemberByID(ctx, member.ID)
		assert.Nil(t, res)
		assert.EqualError(t, err, ErrRecordNotFound.Error())
	})

	t.Run("failed, error delete", func(t *testing.T) {
		errorDelete := errors.New("error delete")

		mockMemberRepo := mock.NewMockMemberRepository(ctrl)
		mockMemberRepo.EXPECT().FindByID(ctx, member.ID).Times(1).Return(member, nil)

		mockMemberRepo.EXPECT().DeleteByID(ctx, member.ID).Times(1).Return(nil, errorDelete)

		memberUsecase := memberUsecase{
			memberRepo: mockMemberRepo,
		}

		res, err := memberUsecase.DeleteMemberByID(ctx, member.ID)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, errorDelete.Error())
	})
}
