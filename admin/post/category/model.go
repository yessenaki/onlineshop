package category

import (
	"database/sql"
	"onlineshop/config"
	"strings"
	"time"
)

// Category struct
type Category struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	PostQnt   int       `db:"post_qnt"`
	Errors    map[string]string
}

func (ctg *Category) validate() bool {
	ctg.Errors = make(map[string]string)
	name := strings.TrimSpace(ctg.Name)

	if name == "" || len(name) > 30 {
		ctg.Errors["Name"] = "The field Name must be a string with a maximum length of 30"
	}

	return len(ctg.Errors) == 0
}

func FindAll() ([]Category, error) {
	rows, err := config.DB.Query("SELECT * FROM post_categories ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ctgs []Category
	for rows.Next() {
		var ctg Category
		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.CreatedAt, &ctg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ctgs = append(ctgs, ctg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ctgs, nil
}

func findOne(id int) (Category, error) {
	var ctg Category
	row := config.DB.QueryRow("SELECT * FROM post_categories WHERE id=$1", id)
	err := row.Scan(&ctg.ID, &ctg.Name, &ctg.CreatedAt, &ctg.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return ctg, err
	}
	return ctg, nil
}

func (ctg *Category) store() (int, error) {
	var id int
	stm := "INSERT INTO post_categories (name, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(stm, ctg.Name).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (ctg *Category) update() error {
	_, err := config.DB.Exec("UPDATE post_categories SET name=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", ctg.Name, ctg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ctg *Category) destroy() error {
	_, err := config.DB.Exec("DELETE FROM post_categories WHERE id=$1", ctg.ID)
	if err != nil {
		return err
	}
	return nil
}
