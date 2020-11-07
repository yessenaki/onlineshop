package color

import (
	"database/sql"
	"github.com/yesseneon/onlineshop/config"
	"strings"
	"time"
)

// Color struct
type Color struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Errors    map[string]string
	Checked   bool
}

func (c *Color) validate() bool {
	c.Errors = make(map[string]string)
	name := strings.TrimSpace(c.Name)

	if name == "" || len(name) > 30 {
		c.Errors["Name"] = "The field Name must be a string with a maximum length of 30"
	}

	return len(c.Errors) == 0
}

func (c *Color) store() (int, error) {
	var lastInsertedID int
	stm := "INSERT INTO colors (name, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(stm, c.Name).Scan(&lastInsertedID)
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

func AllColors() ([]Color, error) {
	rows, err := config.DB.Query("SELECT * FROM colors ORDER BY name")
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

func Arrange(ids []int) ([]Color, error) {
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

		for _, id := range ids {
			if id == color.ID {
				color.Checked = true
				break
			}
		}

		colors = append(colors, color)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return colors, nil
}
