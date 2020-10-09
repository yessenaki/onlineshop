package file

import (
	"onlineshop/config"
	"time"
)

// File struct
type File struct {
	ID        int
	Name      string
	Path      string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f File) Store() (int, error) {
	var id int
	stm := `INSERT INTO files (name, path, product_id, created_at, updated_at) VALUES ($1, $2, $3, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id`
	err := config.DB.QueryRow(stm, f.Name, f.Path, f.ProductID).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func FindByProductID(id int) ([]File, error) {
	rows, err := config.DB.Query("SELECT * FROM files WHERE product_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := []File{}
	for rows.Next() {
		f := File{}
		err := rows.Scan(&f.ID, &f.Name, &f.Path, &f.ProductID, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func Destroy(id int) error {
	_, err := config.DB.Exec("DELETE FROM files WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
