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

func initializeMemberRepositoryWithMock(mockDB *gorm.DB) *memberRepository {
	return &memberRepository{
		db: mockDB,
	}
}

func TestCreateMemberRepo(t *testing.T) {
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
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `members`").
			WithArgs(member.FirstName,
				member.LastName, member.Email, member.CreatedAt, member.UpdatedAt, member.DeletedAt, member.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		ctx := context.TODO()
		err := repo.Create(ctx, member)
		assert.NoError(t, err)
	})

	t.Run("rollback on insert members error", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `members`").
			WithArgs(member.FirstName,
				member.LastName, member.Email, member.CreatedAt, member.UpdatedAt, member.ID).
			WillReturnError(errors.New("some error"))
		mockQuery.ExpectRollback()

		ctx := context.TODO()
		err := repo.Create(ctx, member)
		assert.Error(t, err)
	})

	t.Run("error commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("INSERT INTO `members`").
			WithArgs(member.FirstName,
				member.LastName, member.Email, member.CreatedAt, member.UpdatedAt, member.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error commit"))

		ctx := context.TODO()
		err := repo.Create(ctx, member)
		assert.Error(t, err)
	})
}

func TestFindMemberByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
	}

	dbMock, mockQuery := initializeMySQLMockConn()
	repo := initializeMemberRepositoryWithMock(dbMock)

	row := sqlmock.NewRows(
		[]string{"id", "first_name", "last_name",
			"email", "created_at", "updated_at"},
	).AddRow(member.ID, member.FirstName, member.LastName,
		member.Email, member.CreatedAt, member.UpdatedAt)

	t.Run("success", func(t *testing.T) {

		mockQuery.ExpectQuery("SELECT(.*)").
			WillReturnRows(row)

		ctx := context.TODO()
		member, err := repo.FindByID(ctx, member.ID)
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
		member, err := repo.FindByID(ctx, member.ID)
		assert.Error(t, err)
		assert.Nil(t, member)
	})
}

func TestUpdateMemberByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
	}

	t.Run("success", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "first_name", "last_name",
				"email", "updated_at"},
		).AddRow(member.ID, member.FirstName, member.LastName,
			member.Email, member.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				member.FirstName,
				member.LastName,
				member.Email,
				sqlmock.AnyArg(),
				member.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnRows(sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "email", "updated_at"},
		).AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.UpdatedAt))

		memberRes, err := repo.UpdateByID(context.TODO(), member)
		assert.NoError(t, err)
		assert.NotNil(t, memberRes)
		assert.Equal(t, member.ID, memberRes.ID)
	})

	t.Run("failed, find member by id", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnError(errors.New("error"))

		memberRes, err := repo.UpdateByID(context.TODO(), member)
		assert.Error(t, err)
		assert.Nil(t, memberRes)
	})

	t.Run("failed, old member nil", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnRows(sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "email", "updated_at"}))

		memberRes, err := repo.UpdateByID(context.TODO(), member)
		assert.NoError(t, err)
		assert.Nil(t, memberRes)
	})

	t.Run("failed, error on update", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "first_name", "last_name",
				"email", "updated_at"},
		).AddRow(member.ID, member.FirstName, member.LastName,
			member.Email, member.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				member.FirstName,
				member.LastName,
				member.Email,
				sqlmock.AnyArg(),
				member.ID,
			).
			WillReturnError(errors.New("error"))
		mockQuery.ExpectRollback()

		memberRes, err := repo.UpdateByID(context.TODO(), member)
		assert.Error(t, err)
		assert.Nil(t, memberRes)
	})

	t.Run("failed, error on commit", func(t *testing.T) {
		dbMock, mockQuery := initializeMySQLMockConn()
		repo := initializeMemberRepositoryWithMock(dbMock)

		row := sqlmock.NewRows(
			[]string{"id", "first_name", "last_name",
				"email", "updated_at"},
		).AddRow(member.ID, member.FirstName, member.LastName,
			member.Email, member.UpdatedAt)

		mockQuery.ExpectQuery("SELECT(.*)").WithArgs(member.ID).WillReturnRows(row)

		mockQuery.ExpectBegin()
		mockQuery.ExpectExec("UPDATE(.*)").
			WithArgs(
				member.FirstName,
				member.LastName,
				member.Email,
				sqlmock.AnyArg(),
				member.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error"))

		memberRes, err := repo.UpdateByID(context.TODO(), member)
		assert.Error(t, err)
		assert.Nil(t, memberRes)
	})
}

func TestDeleteMemberByIDRepo(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)

	dbMock, mockQuery := initializeMySQLMockConn()
	repo := initializeMemberRepositoryWithMock(dbMock)

	member := &model.Member{
		ID:        123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		CreatedAt: date,
		UpdatedAt: date,
	}

	memberRow := []string{"id", "first_name", "last_name", "email", "created_at", "updated_at"}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows(memberRow)
		rows.AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.CreatedAt, member.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), member.ID, member.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		memberResult, err := repo.DeleteByID(context.TODO(), member.ID)
		assert.NoError(t, err)
		assert.NotNil(t, memberResult)
	})

	t.Run("failed, error delete", func(t *testing.T) {
		rows := sqlmock.NewRows(memberRow)
		rows.AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.CreatedAt, member.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), member.ID, member.ID).
			WillReturnError(errors.New("error"))
		mockQuery.ExpectRollback()

		memberRes, err := repo.DeleteByID(context.TODO(), member.ID)
		assert.Nil(t, memberRes)
		assert.Error(t, err)
	})

	t.Run("failed, error select", func(t *testing.T) {
		rows := sqlmock.NewRows(memberRow)
		rows.AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.CreatedAt, member.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)
		mockQuery.ExpectRollback()

		memberResult, err := repo.DeleteByID(context.TODO(), member.ID)
		assert.Error(t, err)
		assert.Nil(t, memberResult)
	})

	t.Run("failed, error unscope", func(t *testing.T) {
		rows := sqlmock.NewRows(memberRow)
		rows.AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.CreatedAt, member.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), member.ID, member.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit()

		mockQuery.ExpectQuery("SELECT(.*)").WillReturnError(errors.New("error"))

		memberResult, err := repo.DeleteByID(context.TODO(), member.ID)
		assert.Error(t, err)
		assert.Nil(t, memberResult)
	})

	t.Run("failed, error commit", func(t *testing.T) {
		rows := sqlmock.NewRows(memberRow)
		rows.AddRow(member.ID, member.FirstName, member.LastName, member.Email, member.CreatedAt, member.UpdatedAt)

		mockQuery.ExpectBegin()
		mockQuery.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

		mockQuery.ExpectExec("UPDATE (.*)").
			WithArgs(sqlmock.AnyArg(), member.ID, member.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mockQuery.ExpectCommit().WillReturnError(errors.New("error"))

		memberResult, err := repo.DeleteByID(context.TODO(), member.ID)
		assert.Error(t, err)
		assert.Nil(t, memberResult)
	})
}
