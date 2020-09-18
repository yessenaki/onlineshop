package category

import (
	"database/sql"
	"onlineshop/config"
)

// Category struct
type Category struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	ParentID    int    `db:"parent_id"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	ParentTitle string `db:"parent_title"`
}

func (ctg *Category) store() (int, error) {
	var lastInsertID int
	sqlStatement := "INSERT INTO categories (title, parent_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, ctg.Title, ctg.ParentID).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func (ctg *Category) update() error {
	_, err := config.DB.Exec("UPDATE categories SET title=$1, parent_id=$2, updated_at=NOW()::timestamp(0) WHERE id=$3", ctg.Title, ctg.ParentID, ctg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ctg *Category) destroy() error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id=$1", ctg.ID)
	if err != nil {
		return err
	}
	return nil
}

func allCategories() ([]Category, error) {
	sqlStatement := `SELECT c1.id, c1.title, c1.parent_id, c1.created_at, c1.updated_at, c2.title as parent_title
		FROM categories as c1
		LEFT OUTER JOIN categories as c2
		ON c1.parent_id = c2.id`
	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ctgs := []Category{}
	for rows.Next() {
		ctg := Category{}
		var parentTitle sql.NullString
		err := rows.Scan(&ctg.ID, &ctg.Title, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt, &parentTitle)
		if err != nil {
			return nil, err
		}
		if parentTitle.Valid {
			ctg.ParentTitle = parentTitle.String
		}
		ctgs = append(ctgs, ctg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ctgs, nil
}

func oneCategory(id int) (Category, error) {
	ctg := Category{}
	row := config.DB.QueryRow("SELECT * FROM categories WHERE id = $1", id)
	err := row.Scan(&ctg.ID, &ctg.Title, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctg, nil
		}
		return ctg, err
	}
	return ctg, nil
}
