package color

import (
	"database/sql"
	"onlineshop/config"
)

// Color struct
type Color struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (c *Color) store() (int, error) {
	var lastInsertedID int
	sqlStatement := "INSERT INTO colors (name, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, c.Name).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (c *Color) update() error {
	_, err := config.DB.Exec("UPDATE colors SET name=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", c.Name, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Color) destroy() error {
	_, err := config.DB.Exec("DELETE FROM colors WHERE id=$1", c.ID)
	if err != nil {
		return err
	}
	return nil
}

func allColors() ([]Color, error) {
	rows, err := config.DB.Query("SELECT * FROM colors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colors := []Color{}
	for rows.Next() {
		color := Color{}
		err := rows.Scan(&color.ID, &color.Name, &color.CreatedAt, &color.UpdatedAt)
		if err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return colors, nil
}

func oneColor(id int) (Color, error) {
	color := Color{}
	err := config.DB.QueryRow("SELECT * FROM colors WHERE id=$1", id).Scan(&color.ID, &color.Name, &color.CreatedAt, &color.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return color, nil
		}
		return color, err
	}
	return color, nil
}
