package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	Member struct {
		ID        int64          `json:"id" gorm:"primary_key"`
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"email"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`
	}

	MemberRepository interface {
		Create(ctx context.Context, member *Member) error
		FindByID(ctx context.Context, memberID int64) (*Member, error)
		FindAll(ctx context.Context) ([]int64, error)
		UpdateByID(ctx context.Context, member *Member) (*Member, error)
		DeleteByID(ctx context.Context, memberID int64) (*Member, error)
	}

	MemberUsecase interface {
		Register(ctx context.Context, member *Member) error
		FindMemberByID(ctx context.Context, memberID int64) (*Member, error)
		UpdateMemberByID(ctx context.Context, member *Member) (*Member, error)
		DeleteMemberByID(ctx context.Context, memberID int64) (*Member, error)
	}
)

func (m *Member) ImmutableColumns() []string {
	return []string{"created_at", "deleted_at"}
}
