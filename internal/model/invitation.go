package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	InvitationStatus int

	Invitation struct {
		ID          int64
		MemberID    int64
		GatheringID int64
		Status      InvitationStatus
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	}

	InvitationRepository interface {
		Create(ctx context.Context, invitation *Invitation) error
		FindByID(ctx context.Context, invitationID int64) (*Invitation, error)
		UpdateByID(ctx context.Context, invitation *Invitation) (*Invitation, error)
		DeleteByID(ctx context.Context, invitationID int64) (*Invitation, error)
	}

	InvitationUsecase interface {
		InviteMemberToGathering(ctx context.Context, invitation *Invitation) (*Invitation, error)
		FindInvitationByID(ctx context.Context, invitationID int64) (*Invitation, error)
		UpdateInvitationByID(ctx context.Context, invitation *Invitation) (*Invitation, error)
		DeleteInvitationByID(ctx context.Context, invitationID int64) (*Invitation, error)
	}
)

const (
	Pending = InvitationStatus(1)
	Active  = InvitationStatus(2)
	Expired = InvitationStatus(3)
)

func (i *Invitation) ImmutableColumns() []string {
	return []string{"created_at", "deleted_at"}
}
