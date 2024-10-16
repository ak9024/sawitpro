package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// handler for post new estate
// POST /estate
func (s *Server) PostEstate(ctx echo.Context) error {
	var req generated.EstateRequest
	var errResponse generated.ErrorResponse

	if err := ctx.Bind(&req); err != nil {
		errResponse.Message = "Invalid request body"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	if req.Width <= 0 || req.Width > 50000 {
		errResponse.Message = "Invalid width"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	if req.Length <= 0 || req.Length > 50000 {
		errResponse.Message = "Invalid length"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// @TODO
	// - [ ] Add repository to insert estate to databases
	// - [x] Add uuid string as a estate reponnse
	id := uuid.New().String()
	resp := generated.EstateResponse{
		Id: id,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// handler for store estate tree
// POST /estate/<id>/tree
func (s *Server) PostEstateIdTree(ctx echo.Context, id string) error {
	var req generated.EstateTreeRequest
	var errResponse generated.ErrorResponse

	if err := ctx.Bind(&req); err != nil {
		errResponse.Message = "Invalid request body"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	resp := generated.EstateTreeResponse{
		Id: id,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// handler for get stats of the estate
// GET /estate/<id>/stats
func (s *Server) GetEstateIdStats(ctx echo.Context, id string) error {
	resp := generated.EstateStatsResponse{
		Count:  3,
		Max:    10,
		Min:    10,
		Median: 10,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id string) error {
	resp := generated.EstateDronePlanResponse{
		Distance: 300,
	}

	return ctx.JSON(http.StatusOK, resp)
}
