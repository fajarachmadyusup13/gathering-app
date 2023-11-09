package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	Attendee struct {
		MemberID    int64
		GatheringID int64
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	}

	AttendeeRepository interface {
		Create(ctx context.Context, attendee *Attendee) error
		FindByMemberID(ctx context.Context, memberID int64) ([]*Attendee, error)
		FindByGatheringID(ctx context.Context, gatheringID int64) ([]*Attendee, error)
		DeleteByMemberIDAndGatheringID(ctx context.Context, memberID int64, gatheringID int64) (*Attendee, error)
	}
)
