package model

import (
	"time"

	"github.com/fajarachmadyusup13/gathering-app/internal/model"
)

type RegisterMemberRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email,max=320"`
}

type UpdateMemberRequest struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email,max=320"`
}

type FindByIDRequest struct {
	ID int64 `json:"id"`
}

type CreateGatheringRequest struct {
	Creator     int64               `json:"creator"`
	Name        string              `json:"name"`
	Location    string              `json:"location"`
	ScheduledAt *time.Time          `json:"scheduled_at"`
	Type        model.GatheringType `json:"type"`
}

type UpdateGatheringRequest struct {
	ID          int64               `json:"id"`
	Creator     int64               `json:"creator"`
	Name        string              `json:"name"`
	Location    string              `json:"location"`
	ScheduledAt *time.Time          `json:"scheduled_at"`
	Type        model.GatheringType `json:"type"`
}

type CreateInvitationRequest struct {
	MemberID    int64                  `json:"member_id"`
	GatheringID int64                  `json:"gathering_id"`
	Status      model.InvitationStatus `json:"status"`
}

type UpdateInvitationRequest struct {
	ID          int64                  `json:"id"`
	MemberID    int64                  `json:"member_id"`
	GatheringID int64                  `json:"gathering_id"`
	Status      model.InvitationStatus `json:"status"`
}
