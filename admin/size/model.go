package size

import (
	"database/sql"
	"onlineshop/config"
	"strings"
)

// Size struct
type Size struct {
	ID        int    `db:"id"`
	Size      string `db:"size"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	Type      int    `db:"type"`
	Errors    map[string]string
}

func (s *Size) validate() bool {
	s.Errors = make(map[string]string)
	size := strings.TrimSpace(s.Size)

	if size == "" || len(size) > 10 {
		s.Errors["Size"] = "The field Size must be a string with a maximum length of 10"
	}

	return len(s.Errors) == 0
}

func (s *Size) store() (int, error) {
	var lastInsertedID int
	sqlStatement := "INSERT INTO sizes (size, type, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, s.Size, s.Type).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (s *Size) update() error {
	_, err := config.DB.Exec("UPDATE sizes SET size=$1, type=$2, updated_at=NOW()::timestamp(0) WHERE id=$3", s.Size, s.Type, s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Size) destroy() error {
	_, err := config.DB.Exec("DELETE FROM sizes WHERE id=$1", s.ID)
	if err != nil {
		return err
	}
	return nil
}

func AllSizes() ([]Size, error) {
	rows, err := config.DB.Query("SELECT * FROM sizes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sizes := []Size{}
	for rows.Next() {
		size := Size{}
		err := rows.Scan(&size.ID, &size.Size, &size.CreatedAt, &size.UpdatedAt, &size.Type)
		if err != nil {
			return nil, err
		}
		sizes = append(sizes, size)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sizes, nil
}

func oneSize(id int) (Size, error) {
	size := Size{}
	err := config.DB.QueryRow("SELECT * FROM sizes WHERE id=$1", id).Scan(&size.ID, &size.Size, &size.CreatedAt, &size.UpdatedAt, &size.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return size, nil
		}
		return size, err
	}
	return size, nil
}
