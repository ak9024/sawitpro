package repository

import (
	"context"
	"database/sql"
)

// Create new estate
func (r *Repository) CreateEstate(ctx context.Context, input Estate) (output Estate, err error) {
	var id string
	err = r.Db.QueryRowContext(ctx, InsertEstatesQuery,
		input.Id,
		input.Width,
		input.Length,
	).Scan(&id)
	if err != nil {
		return output, err
	}

	output.Id = id
	return
}

// Create new estate tree
func (r *Repository) CreateEstateTree(ctx context.Context, input EstateTree) (output EstateTree, err error) {
	var id string
	err = r.Db.QueryRowContext(ctx, InsertEstateTreeQuery,
		input.Id,
		input.EstateID,
		input.X,
		input.Y,
		input.Height,
	).Scan(&id)
	if err != nil {
		return output, err
	}

	output.Id = id
	output.EstateID = input.EstateID
	output.X = input.X
	output.Y = input.Y
	output.Height = input.Height

	return
}

// Get stats from trees
func (c *Repository) GetStats(ctx context.Context, estateId string) (treeCount, maxHeight, minHeight int, medianHeight float64, err error) {
	err = c.Db.QueryRowContext(ctx, GetStatsQuery, estateId).Scan(
		&treeCount,
		&maxHeight,
		&minHeight,
		&medianHeight,
	)
	return
}

// Get estate by id
func (c *Repository) GetEstateById(ctx context.Context, id string) (width, length int, err error) {
	err = c.Db.QueryRowContext(ctx, GetEstateByIdQuery, id).Scan(&width, &length)
	return
}

// Get trees by id
func (c *Repository) GetTreesById(ctx context.Context, estateId string) (rows *sql.Rows, err error) {
	rows, err = c.Db.QueryContext(ctx, GetTreesByIdQuery, estateId)
	return
}
