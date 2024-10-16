package repository

import (
	"context"
	"database/sql"
)

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

func (c *Repository) GetEstateTreeById(ctx context.Context, id string) (output []EstateTree, exists bool, err error) {
	rows, err := c.Db.QueryContext(ctx, GetEstateTreeByIdQuery, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	for rows.Next() {
		var tree EstateTree
		err := rows.Scan(&tree.Id, &tree.EstateID, &tree.X, &tree.Y, &tree.Height)
		if err != nil {
			return nil, false, err
		}
		output = append(output, tree)
	}

	if err = rows.Err(); err != nil {
		return nil, false, err
	}

	if len(output) == 0 {
		return nil, false, nil
	}

	return output, true, nil
}

func (c *Repository) GetEstateById(ctx context.Context, id string) (output Estate, exists bool, err error) {
	row := c.Db.QueryRowContext(ctx, GetEstateByIdQuery, id)

	err = row.Scan(&output.Id, &output.Width, &output.Length)
	if err != nil {
		if err == sql.ErrNoRows {
			return Estate{}, false, nil
		}
		return Estate{}, false, err
	}

	return output, true, nil
}
