package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	GatheringType int

	Gathering struct {
		ID          int64          `json:"id"`
		Creator     int64          `json:"creator"`
		Type        GatheringType  `json:"type"`
		ScheduledAt *time.Time     `json:"scheduled_at"`
		Name        string         `json:"name"`
		Location    string         `json:"location"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	}

	GatheringRepository interface {
		Create(ctx context.Context, gathering *Gathering) error
		FindByID(ctx context.Context, gatheringID int64) (*Gathering, error)
		UpdateByID(ctx context.Context, gathering *Gathering) (*Gathering, error)
		DeleteByID(ctx context.Context, gatheringID int64) (*Gathering, error)
	}

	GatheringUsecase interface {
		CreateGathering(ctx context.Context, gathering *Gathering) error
		FindGatheringByID(ctx context.Context, gatheringID int64) (*Gathering, error)
		UpdateGatheringByID(ctx context.Context, gathering *Gathering) (*Gathering, error)
		DeleteGatheringByID(ctx context.Context, gatheringID int64) (*Gathering, error)
	}
)

const (
	WithFixedNumberOfAttendees   = GatheringType(1)
	WithExpirationForInvitations = GatheringType(2)
)

func (g *Gathering) ImmutableColumns() []string {
	return []string{"created_at"}
}
