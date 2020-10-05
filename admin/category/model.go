package category

import (
	"database/sql"
	"onlineshop/config"
	"strings"
	"time"
)

// Category struct
type Category struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	ParentID   int       `db:"parent_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Gender     int       `db:"gender"`
	IsKids     int       `db:"is_kids"`
	ParentName string    `db:"parent_name"`
	Errors     map[string]string
	Active     bool
}

// Parent struct
type Parent struct {
	ID     int
	Name   string
	Active bool
	Childs []Category
}

func ByParams(gender int, isKids int, ctgID int) ([]Parent, error) {
	sql := "SELECT * FROM categories WHERE gender IN (2, $1) AND is_kids IN (2, $2) ORDER BY id"

	rows, err := config.DB.Query(sql, gender, isKids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ctgs := []Category{}
	for rows.Next() {
		ctg := Category{}

		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt, &ctg.Gender, &ctg.IsKids)
		if err != nil {
			return nil, err
		}

		ctgs = append(ctgs, ctg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	parents := []Parent{}
	for _, ctg := range ctgs {
		if ctg.ParentID == 0 {
			parent := Parent{}
			parent.ID = ctg.ID
			parent.Name = ctg.Name

			childs := []Category{}
			for _, c := range ctgs {
				if c.ParentID == ctg.ID {
					if c.ID == ctgID {
						parent.Active = true
						c.Active = true
					}
					childs = append(childs, c)
				}
			}

			parent.Childs = childs
			parents = append(parents, parent)
		}
	}

	return parents, nil
}

func (ctg *Category) validate() bool {
	ctg.Errors = make(map[string]string)
	name := strings.TrimSpace(ctg.Name)

	if name == "" || len(name) > 30 {
		ctg.Errors["Name"] = "The field Name must be a string with a maximum length of 30"
	}

	return len(ctg.Errors) == 0
}

func (ctg *Category) store() (int, error) {
	var lastInsertedID int
	sqlStatement := `INSERT INTO categories (name, parent_id, created_at, updated_at, gender, is_kids)
		VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0), $3, $4) RETURNING id`
	err := config.DB.QueryRow(sqlStatement, ctg.Name, ctg.ParentID, ctg.Gender, ctg.IsKids).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (ctg *Category) update() error {
	sql := `UPDATE categories SET name=$1, parent_id=$2, updated_at=NOW()::timestamp(0), gender=$3, is_kids=$4 WHERE id=$5`
	_, err := config.DB.Exec(sql, ctg.Name, ctg.ParentID, ctg.Gender, ctg.IsKids, ctg.ID)
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
	sqlStatement := `SELECT c1.*, c2.name as parent_name
		FROM categories as c1
		LEFT OUTER JOIN categories as c2
		ON c1.parent_id=c2.id
		ORDER BY c1.id`

	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ctgs := []Category{}
	for rows.Next() {
		ctg := Category{}
		var parentName sql.NullString

		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt, &ctg.Gender, &ctg.IsKids, &parentName)
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
	err := row.Scan(&ctg.ID, &ctg.Name, &ctg.ParentID, &ctg.CreatedAt, &ctg.UpdatedAt, &ctg.Gender, &ctg.IsKids)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctg, nil
		}
		return ctg, err
	}
	return ctg, nil
}
