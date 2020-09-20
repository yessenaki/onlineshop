package category

import (
	"database/sql"
	"onlineshop/config"
)

// Category struct
type Category struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	ParentID   int    `db:"parent_id"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	ParentName string `db:"parent_name"`
}

func (ctg *Category) store() (int, error) {
	var lastInsertedID int
	sqlStatement := "INSERT INTO categories (name, parent_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, ctg.Name, ctg.ParentID).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (ctg *Category) update() error {
	_, err := config.DB.Exec("UPDATE categories SET name=$1, parent_id=$2, updated_at=NOW()::timestamp(0) WHERE id=$3", ctg.Name, ctg.ParentID, ctg.ID)
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

func AllCategories() ([]Category, error) {
	sqlStatement := `SELECT c1.id, c1.name, c1.parent_id, c1.created_at, c1.updated_at, c2.name as parent_name
		FROM categories as c1
		LEFT OUTER JOIN categories as c2
		ON c1.parent_id=c2.id`
	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ctgs := []Category{}
	for rows.Next() {
		ctg := Category{}
		var parentName sql.NullString
		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt, &parentName)
		if err != nil {
			return nil, err
		}
		if parentName.Valid {
			ctg.ParentName = parentName.String
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
	row := config.DB.QueryRow("SELECT * FROM categories WHERE id=$1", id)
	err := row.Scan(&ctg.ID, &ctg.Name, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctg, nil
		}
		return ctg, err
	}
	return ctg, nil
}
