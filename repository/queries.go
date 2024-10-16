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
	// GetStatsQuery retrieves statistical information about trees for a given estate.
	// It returns the count of trees, maximum height, minimum height, and median height.
	GetStatsQuery = `SELECT COUNT(*), MAX(height), MIN(height), 
    PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY height)
    FROM trees WHERE estate_id = $1`
	GetEstateByIdQuery = `SELECT width, length FROM estates WHERE id = $1`
	GetTreesByIdQuery  = `SELECT height FROM trees WHERE estate_id = $1`
)
