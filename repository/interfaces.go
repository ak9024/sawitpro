// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	CreateEstate(ctx context.Context, input Estate) (output Estate, err error)
	CreateEstateTree(ctx context.Context, input EstateTree) (output EstateTree, err error)
	GetEstateTreeById(ctx context.Context, id string) (output []EstateTree, exists bool, err error)
	GetEstateById(ctx context.Context, id string) (output Estate, exists bool, err error)
}
