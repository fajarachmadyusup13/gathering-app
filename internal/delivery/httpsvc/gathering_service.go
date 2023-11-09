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

func (s *HTTPService) CreateGathering(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.CreateGatheringRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	gathering := &model.Gathering{
		ID:          model.GenerateID(),
		Creator:     body.Creator,
		ScheduledAt: body.ScheduledAt,
		Type:        body.Type,
		Name:        body.Name,
		Location:    body.Location,
	}

	err = s.gatheringUsecase.CreateGathering(ctx, gathering)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusCreated, gathering)
}

func (s *HTTPService) FindGatheringByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	gathering, err := s.gatheringUsecase.FindGatheringByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, &gathering)
}

func (s *HTTPService) UpdateGathering(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.UpdateGatheringRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	gathering := &model.Gathering{
		ID:          body.ID,
		Creator:     body.Creator,
		ScheduledAt: body.ScheduledAt,
		Type:        body.Type,
		Name:        body.Name,
		Location:    body.Location,
	}

	res, err := s.gatheringUsecase.UpdateGatheringByID(ctx, gathering)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}

func (s *HTTPService) DeleteGatheringByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	res, err := s.gatheringUsecase.DeleteGatheringByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}
