package httpsvc

import (
	"github.com/fajarachmadyusup13/gathering-app/internal/delivery/httpsvc/middleware"
	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/gin-gonic/gin"
)

type HTTPService struct {
	memberUsecase     model.MemberUsecase
	gatheringUsecase  model.GatheringUsecase
	invitationUsecase model.InvitationUsecase
}

func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

func (s *HTTPService) InitRoutes(route *gin.Engine) {
	route.Use(middleware.CustomErrorMiddleware)

	member := route.Group("/member")
	member.POST("register", s.RegisterMember)
	member.POST("update", s.UpdateMember)
	member.GET("findByID", s.FindMemberByID)
	member.POST("deleteByID", s.DeleteMemberByID)

	gathering := route.Group("/gathering")
	gathering.POST("create", s.CreateGathering)
	gathering.POST("update", s.UpdateGathering)
	gathering.GET("findByID", s.FindGatheringByID)
	gathering.POST("deleteByID", s.DeleteGatheringByID)

	invitation := route.Group("/invitation")
	invitation.POST("invite", s.InviteMemberToGathering)
	invitation.GET("findByID", s.FindInvitationByID)
	invitation.POST("update", s.UpdateInvitation)
	invitation.POST("deleteByID", s.DeleteInvitationByID)

}

func (s *HTTPService) RegisterMemberUsecase(m model.MemberUsecase) {
	s.memberUsecase = m
}

func (s *HTTPService) RegisterGatheringUsecase(g model.GatheringUsecase) {
	s.gatheringUsecase = g
}

func (s *HTTPService) RegisterInvitationUsecase(i model.InvitationUsecase) {
	s.invitationUsecase = i
}
