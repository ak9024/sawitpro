// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"database/sql"
)

type RepositoryInterface interface {
	CreateEstate(ctx context.Context, input Estate) (output Estate, err error)
	CreateEstateTree(ctx context.Context, input EstateTree) (output EstateTree, err error)
	GetStats(ctx context.Context, estateId string) (treeCount, maxHeight, minHeight int, medianHeight float64, err error)
	GetEstateById(ctx context.Context, id string) (width, length int, err error)
	GetTreesById(ctx context.Context, estateId string) (rows *sql.Rows, err error)
}
