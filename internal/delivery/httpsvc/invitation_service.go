package httpsvc

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fajarachmadyusup13/gathering-app/internal/delivery/httpsvc/middleware"
	httpsvcModel "github.com/fajarachmadyusup13/gathering-app/internal/delivery/httpsvc/model"
	"github.com/fajarachmadyusup13/gathering-app/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *HTTPService) InviteMemberToGathering(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.CreateInvitationRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	invitation := &model.Invitation{
		ID:          model.GenerateID(),
		MemberID:    body.MemberID,
		GatheringID: body.GatheringID,
		Status:      body.Status,
	}

	res, err := s.invitationUsecase.InviteMemberToGathering(ctx, invitation)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusCreated, res)
}

func (s *HTTPService) FindInvitationByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	invitation, err := s.invitationUsecase.FindInvitationByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, &invitation)
}

func (s *HTTPService) UpdateInvitation(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.UpdateInvitationRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	invitation := &model.Invitation{
		ID:          body.ID,
		MemberID:    body.MemberID,
		GatheringID: body.GatheringID,
		Status:      body.Status,
	}

	res, err := s.invitationUsecase.UpdateInvitationByID(ctx, invitation)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}

func (s *HTTPService) DeleteInvitationByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	res, err := s.invitationUsecase.DeleteInvitationByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}
