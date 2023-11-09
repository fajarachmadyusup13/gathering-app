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

func (s *HTTPService) RegisterMember(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.RegisterMemberRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	member := &model.Member{
		ID:        model.GenerateID(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	err = s.memberUsecase.Register(ctx, member)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusCreated, &member)
}

func (s *HTTPService) FindMemberByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	member, err := s.memberUsecase.FindMemberByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, &member)
}

func (s *HTTPService) UpdateMember(c *gin.Context) {
	ctx := c.Request.Context()
	body := httpsvcModel.UpdateMemberRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	member := &model.Member{
		ID:        body.ID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	res, err := s.memberUsecase.UpdateMemberByID(ctx, member)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}

func (s *HTTPService) DeleteMemberByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Query("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		err = middleware.NewHTTPError(http.StatusBadRequest, err.Error())
		c.Error(err)
	}

	res, err := s.memberUsecase.DeleteMemberByID(ctx, int64(intID))
	if err != nil {
		err = middleware.NewHTTPError(http.StatusInternalServerError, err.Error())
		c.Error(err)
	}

	c.JSON(http.StatusOK, res)
}
