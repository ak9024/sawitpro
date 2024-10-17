package handler

import (
	"net/http"

	"github.com/ak9024/sawitpro/generated"
	"github.com/ak9024/sawitpro/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Estate struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Length int    `json:"length"`
}

type Tree struct {
	ID     string `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Height int    `json:"height"`
}

var estates = map[string]Estate{}
var trees = map[string][]Tree{}

// handler for post new estate
// POST /estate
func (s *Server) PostEstate(c echo.Context) error {
	ctx := c.Request().Context()

	var req generated.EstateRequest
	var errResponse generated.ErrorResponse

	if err := c.Bind(&req); err != nil {
		errResponse.Message = "Invalid request body"
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	// initialize new uuid
	id := uuid.New().String()

	// insert estate to database
	if _, err := s.Repository.CreateEstate(ctx, repository.Estate{
		Id:     id,
		Width:  req.Width,
		Length: req.Length,
	}); err != nil {
		errResponse.Message = "Error to create new estate"
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	return c.JSON(http.StatusOK, generated.EstateResponse{
		Id: id,
	})
}

// handler for store estate tree
// POST /estate/<id>/tree
func (s *Server) PostEstateIdTree(c echo.Context, id string) error {
	ctx := c.Request().Context()

	var req generated.EstateTreeRequest
	var errResponse generated.ErrorResponse

	if err := c.Bind(&req); err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	treeID := uuid.New().String()

	if _, err := s.Repository.CreateEstateTree(ctx, repository.EstateTree{
		Id:       treeID,
		EstateID: id,
		X:        req.X,
		Y:        req.Y,
		Height:   req.Height,
	}); err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	return c.JSON(http.StatusOK, generated.EstateTreeResponse{
		Id: id,
	})
}

// handler for get stats of the estate
// GET /estate/<id>/stats
func (s *Server) GetEstateIdStats(c echo.Context, id string) error {
	ctx := c.Request().Context()

	var errResponse generated.ErrorResponse

	count, min, max, median, err := s.Repository.GetStats(ctx, id)
	if err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	return c.JSON(http.StatusOK, generated.EstateStatsResponse{
		Count:  count,
		Min:    min,
		Max:    max,
		Median: int(median),
	})
}

// handler for get drone plan
// GET /estate/<id>/dron-plan
func (s *Server) GetEstateIdDronePlan(c echo.Context, id string) error {
	ctx := c.Request().Context()

	var errResponse generated.ErrorResponse

	width, length, err := s.Repository.GetEstateById(ctx, id)
	if err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	// get trees base on estate_id from database
	rows, err := s.Repository.GetTreesById(ctx, id)
	if err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	defer rows.Close()

	// horizontal = (width  - 1) x length + (length - 1) x width
	horizontalDistance := (width-1)*length + (length-1)*width

	// set vertical distance of drone 0 meter above the ground
	verticalDistance := 0

	for rows.Next() {
		var height int

		if err := rows.Scan(&height); err != nil {
			errResponse.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, errResponse)
		}

		// for each tree the drone vertical distance increase +1 meter.
		verticalDistance += height
	}

	// final distance, back to the ground
	verticalDistance += 1

	return c.JSON(http.StatusOK, generated.EstateDronePlanResponse{
		Distance: horizontalDistance + verticalDistance,
	})
}
