package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initializeInvitationRepositoryWithMock(mockDB *gorm.DB) *invitationRepository {
	return &invitationRepository{
		db: mockDB,
	}
}

func TestCreateInvitationRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 222,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `invitations`").
			WithArgs(invitation.MemberID,
				invitation.GatheringID, invitation.Status, invitation.CreatedAt, invitation.UpdatedAt, invitation.DeletedAt, invitation.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		ctx := context.TODO()
		err := repo.Create(ctx, invitation)
		assert.NoError(t, err)
	})

	t.Run("rollback on insert members error", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `invitations`").
			WithArgs(invitation.MemberID,
				invitation.GatheringID, invitation.Status, invitation.CreatedAt, invitation.UpdatedAt, invitation.DeletedAt, invitation.ID).
			WillReturnError(errors.New("some error"))
		mockQuery.ExpectRollback()

		ctx := context.TODO()
		err := repo.Create(ctx, invitation)
		assert.Error(t, err)
	})

	t.Run("error commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `members`").
			WithArgs(invitation.MemberID,
				invitation.GatheringID, invitation.Status, invitation.CreatedAt, invitation.UpdatedAt, invitation.DeletedAt, invitation.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error commit"))

		ctx := context.TODO()
		err := repo.Create(ctx, invitation)
		assert.Error(t, err)
	})
}

func TestFindInvitationByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 222,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	dbMock, mockQuery := initializeMySQLMockConn()
	repo := initializeInvitationRepositoryWithMock(dbMock)

	row := sqlmock.NewRows(
		[]string{"id", "member_id", "gathering_id",
			"status", "created_at", "updated_at"},
	).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
		invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

	t.Run("success", func(t *testing.T) {

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(row)

		ctx := context.TODO()
		member, err := repo.FindByID(ctx, invitation.ID)
		assert.NoError(t, err)
		assert.NotNil(t, member)
	})

	t.Run("error not found", func(t *testing.T) {
		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(gorm.ErrRecordNotFound)

		ctx := context.TODO()
		member, err := repo.FindByID(ctx, 0)
		assert.Nil(t, err)
		assert.Nil(t, member)
	})

	t.Run("error from DB", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(errors.New("text"))

		ctx := context.TODO()
		member, err := repo.FindByID(ctx, invitation.ID)
		assert.Error(t, err)
		assert.Nil(t, member)
	})
}

func TestUpdateInvitationByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 222,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				invitation.MemberID,
				invitation.GatheringID,
				invitation.Status,
				sqlmock.AnyArg(),
				invitation.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id", "status", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID, invitation.Status, invitation.UpdatedAt))

		invitationRes, err := repo.UpdateByID(context.TODO(), invitation)
		assert.NoError(t, err)
		assert.NotNil(t, invitationRes)
		assert.Equal(t, invitation.ID, invitationRes.ID)
	})

	t.Run("failed, find member by id", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnError(errors.New("error"))

		invitationRes, err := repo.UpdateByID(context.TODO(), invitation)
		assert.Error(t, err)
		assert.Nil(t, invitationRes)
	})

	t.Run("failed, old member nil", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id", "status", "updated_at"}))

		invitationRes, err := repo.UpdateByID(context.TODO(), invitation)
		assert.NoError(t, err)
		assert.Nil(t, invitationRes)
	})

	t.Run("failed, error on updatex", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				invitation.MemberID,
				invitation.GatheringID,
				invitation.Status,
				sqlmock.AnyArg(),
				invitation.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id", "status", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID, invitation.Status, invitation.UpdatedAt))

		invitationRes, err := repo.UpdateByID(context.TODO(), invitation)
		assert.NoError(t, err)
		assert.NotNil(t, invitationRes)
		assert.Equal(t, invitation.ID, invitationRes.ID)
	})

	t.Run("failed, error on commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(invitation.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				invitation.MemberID,
				invitation.GatheringID,
				invitation.Status,
				sqlmock.AnyArg(),
				invitation.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error"))

		invitaionRes, err := repo.UpdateByID(context.TODO(), invitation)
		assert.Error(t, err)
		assert.Nil(t, invitaionRes)
	})
}

func TestDeleteInvitationByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	invitation := &model.Invitation{
		ID:          123,
		MemberID:    321,
		GatheringID: 222,
		Status:      model.Active,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), invitation.ID, invitation.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		invitationRes, err := repo.DeleteByID(context.TODO(), invitation.ID)
		assert.NoError(t, err)
		assert.NotNil(t, invitationRes)
	})

	t.Run("failed, error delete", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), invitation.ID, invitation.ID).
			WillReturnError(errors.New("error"))
		mockQuery.ExpectRollback()

		invitationRes, err := repo.DeleteByID(context.TODO(), invitation.ID)
		assert.Nil(t, invitationRes)
		assert.Error(t, err)
	})

	t.Run("failed, error select", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)
		mockQuery.ExpectRollback()

		invitationRes, err := repo.DeleteByID(context.TODO(), invitation.ID)
		assert.Error(t, err)
		assert.Nil(t, invitationRes)
	})

	t.Run("failed, error unscope", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), invitation.ID, invitation.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WillReturnError(errors.New("error"))

		invitationRes, err := repo.DeleteByID(context.TODO(), invitation.ID)
		assert.Error(t, err)
		assert.Nil(t, invitationRes)
	})

	t.Run("failed, error commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeInvitationRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"id", "member_id", "gathering_id",
				"status", "created_at", "updated_at"},
		).AddRow(invitation.ID, invitation.MemberID, invitation.GatheringID,
			invitation.Status, invitation.CreatedAt, invitation.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), invitation.ID, invitation.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error"))

		invitationRes, err := repo.DeleteByID(context.TODO(), invitation.ID)
		assert.Error(t, err)
		assert.Nil(t, invitationRes)
	})
}
