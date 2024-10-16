package handler

import (
	"net/http"
	"sort"

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
		errResponse.Message = err.Error()
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

	var resp generated.EstateStatsResponse
	var errResponse generated.ErrorResponse

	estateTrees, exists, err := s.Repository.GetEstateTreeById(ctx, id)
	if err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	if exists || len(estateTrees) == 0 {
		resp.Count = 0
		resp.Max = 0
		resp.Min = 0
		resp.Median = 0
		return c.JSON(http.StatusOK, resp)
	}

	heights := []int{}
	for _, tree := range estateTrees {
		heights = append(heights, tree.Height)
	}

	sort.Ints(heights)
	count := len(heights)
	resp.Count = count

	// if count is exists or more than 0
	// execute to calculate
	if count > 0 {
		resp.Max = heights[count-1]
		resp.Min = heights[0]
		// if count % 2 != 0, get median from heights[count/2]
		if count%2 == 0 {
			resp.Median = (heights[count/2-1] + heights[count/2]) / 2
		} else {
			resp.Median = heights[count/2]
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// handler for get drone plan
// GET /estate/<id>/dron-plan
func (s *Server) GetEstateIdDronePlan(c echo.Context, id string) error {
	ctx := c.Request().Context()

	var errResponse generated.ErrorResponse

	estate, exists, err := s.Repository.GetEstateById(ctx, id)
	if err != nil {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	if !exists {
		errResponse.Message = "Estate not found"
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	// get trees base on estate_id from database
	estateTrees, exists, err := s.Repository.GetEstateTreeById(ctx, id)
	if err != nil || len(estateTrees) == 0 {
		errResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	// formula zigzag pattern
	// horizontal = (width  - 1) x length + (length - 1) x width
	horizontalDistance := (estate.Width-1)*estate.Length + (estate.Length-1)*estate.Width

	// set vertical distance of drone 0 meter above the ground
	verticalDistance := 0

	for _, tree := range estateTrees {
		// for each tree the drone: tree height +1 meter.
		verticalDistance += tree.Height
	}

	// final distance, back to the ground
	verticalDistance += 1

	return c.JSON(http.StatusOK, generated.EstateDronePlanResponse{
		Distance: horizontalDistance + verticalDistance,
	})
}
