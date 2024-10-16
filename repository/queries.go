package repository

const (
	InsertEstatesQuery = `
    INSERT INTO estates (id, width, length)
    VALUES ($1, $2, $3)
    returning id`

	InsertEstateTreeQuery = `
    INSERT INTO trees (id, estate_id, x, y, height)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

	GetEstateTreeByIdQuery = `SELECT * FROM trees WHERE id = $1`

	GetEstateByIdQuery = `SELECT * FROM estates WHERE id = $1`
)
