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

func initializeAttendeeRepositoryWithMock(mockDB *gorm.DB) *attendeeRepository {
	return &attendeeRepository{
		db: mockDB,
	}
}

func TestCreateAttendeeRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	attendee := &model.Attendee{
		MemberID:    321,
		GatheringID: 222,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}
	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `attendees`").
			WithArgs(attendee.MemberID,
				attendee.GatheringID, attendee.CreatedAt, attendee.UpdatedAt, attendee.DeletedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		ctx := context.TODO()
		err := repo.Create(ctx, attendee)
		assert.NoError(t, err)
	})

	t.Run("rollback on insert attendee error", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `attendees`").
			WithArgs(attendee.MemberID,
				attendee.GatheringID, attendee.CreatedAt, attendee.UpdatedAt, attendee.DeletedAt).
			WillReturnError(errors.New("error"))
		mockQuery.ExpectRollback()

		ctx := context.TODO()
		err := repo.Create(ctx, attendee)
		assert.Error(t, err)
	})

	t.Run("error commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `attendees`").
			WithArgs(attendee.MemberID,
				attendee.GatheringID, attendee.CreatedAt, attendee.UpdatedAt, attendee.DeletedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error"))

		ctx := context.TODO()
		err := repo.Create(ctx, attendee)
		assert.Error(t, err)
	})
}

func TestFindAttendeeByMemberIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	attendee := &model.Attendee{
		MemberID:    321,
		GatheringID: 222,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"member_id", "gathering_id",
				"created_at", "updated_at"},
		).AddRow(attendee.MemberID, attendee.GatheringID,
			attendee.CreatedAt, attendee.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(row)

		res, err := repo.FindByMemberID(context.TODO(), attendee.MemberID)
		assert.NotNil(t, res)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, int64(attendee.MemberID), res[0].MemberID)
		assert.NoError(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := repo.FindByMemberID(context.TODO(), attendee.MemberID)
		assert.Nil(t, res)
		assert.NoError(t, err)
	})

	t.Run("error from DB", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(errors.New("error"))

		res, err := repo.FindByMemberID(context.TODO(), attendee.MemberID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})
}

func TestFindAttendeeByGatheringIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	attendee := &model.Attendee{
		MemberID:    321,
		GatheringID: 222,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"member_id", "gathering_id",
				"created_at", "updated_at"},
		).AddRow(attendee.MemberID, attendee.GatheringID,
			attendee.CreatedAt, attendee.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(row)

		res, err := repo.FindByGatheringID(context.TODO(), attendee.GatheringID)
		assert.NotNil(t, res)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, int64(attendee.GatheringID), res[0].GatheringID)
		assert.NoError(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(gorm.ErrRecordNotFound)

		res, err := repo.FindByGatheringID(context.TODO(), attendee.GatheringID)
		assert.Nil(t, res)
		assert.NoError(t, err)
	})

	t.Run("error from DB", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(errors.New("error"))

		res, err := repo.FindByGatheringID(context.TODO(), attendee.GatheringID)
		assert.Nil(t, res)
		assert.Error(t, err)
	})
}

func TestDeleteAttendeeByMemberIDAndGatheringIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	attendee := &model.Attendee{
		MemberID:    321,
		GatheringID: 222,
		CreatedAt:   date,
		UpdatedAt:   date,
		DeletedAt:   gorm.DeletedAt{},
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"member_id", "gathering_id",
				"created_at", "updated_at"},
		).AddRow(attendee.MemberID, attendee.GatheringID,
			attendee.CreatedAt, attendee.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), attendee.MemberID, attendee.GatheringID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(rows)

		attendeeRes, err := repo.DeleteByMemberIDAndGatheringID(context.TODO(), attendee.MemberID, attendee.GatheringID)
		assert.NoError(t, err)
		assert.NotNil(t, attendeeRes)
	})

	t.Run("failed, error delete", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		rows := sqlmock.NewRows(
			[]string{"member_id", "gathering_id",
				"created_at", "updated_at"},
		).AddRow(attendee.MemberID, attendee.GatheringID,
			attendee.CreatedAt, attendee.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), attendee.MemberID, attendee.GatheringID).
			WillReturnError(errors.New("error"))
		mockQuery.ExpectRollback()

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(rows)

		attendeeRes, err := repo.DeleteByMemberIDAndGatheringID(context.TODO(), attendee.MemberID, attendee.GatheringID)
		assert.Error(t, err)
		assert.Nil(t, attendeeRes)
	})

	t.Run("failed, error select", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeAttendeeRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), attendee.MemberID, attendee.GatheringID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnError(errors.New("error"))

		attendeeRes, err := repo.DeleteByMemberIDAndGatheringID(context.TODO(), attendee.MemberID, attendee.GatheringID)
		assert.Error(t, err)
		assert.Nil(t, attendeeRes)
	})
}
