package handler

import (
	"net/http"
	"sort"

	"github.com/ak9024/sawitpro/generated"
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

	id := uuid.New().String()

	newEstate := Estate{
		ID:     id,
		Width:  req.Width,
		Length: req.Length,
	}

	estates[id] = newEstate

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

	newTree := Tree{
		ID:     id,
		X:      req.X,
		Y:      req.Y,
		Height: req.Height,
	}

	trees[id] = append(trees[id], newTree)

	resp := generated.EstateTreeResponse{
		Id: id,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// handler for get stats of the estate
// GET /estate/<id>/stats
func (s *Server) GetEstateIdStats(ctx echo.Context, id string) error {
	var resp generated.EstateStatsResponse

	estateTrees, exists := trees[id]
	if exists || len(estateTrees) == 0 {
		resp.Count = 0
		resp.Max = 0
		resp.Min = 0
		resp.Median = 0
	}

	heights := []int{}
	for _, tree := range estateTrees {
		heights = append(heights, tree.Height)
	}

	sort.Ints(heights)
	count := len(estateTrees)
	resp.Count = count
	resp.Max = heights[count-1]
	resp.Min = heights[0]
	resp.Median = heights[count/2]

	return ctx.JSON(http.StatusOK, resp)
}

// handler for get drone plan
// GET /estate/<id>/dron-plan
func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id string) error {
	var resp generated.EstateDronePlanResponse
	var errResponse generated.ErrorResponse

	estate, exists := estates[id]
	if !exists {
		errResponse.Message = "Estate not found"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	estateTrees, exists := trees[id]
	if !exists {
		estateTrees = []Tree{}
	}

	// Formula zigzag pattern
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

	// total = horizontal + vertical
	resp.Distance = horizontalDistance + verticalDistance

	return ctx.JSON(http.StatusOK, resp)
}
