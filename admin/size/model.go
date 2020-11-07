package size

import (
	"database/sql"
	"github.com/yesseneon/onlineshop/admin/category"
	"github.com/yesseneon/onlineshop/config"
	"strings"
	"time"
)

// Size struct
type Size struct {
	ID        int       `db:"id"`
	Size      string    `db:"size"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Type      int       `db:"type"`
	Errors    map[string]string
	Checked   bool
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
	stm := "INSERT INTO sizes (size, type, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(stm, s.Size, s.Type).Scan(&lastInsertedID)
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
	rows, err := config.DB.Query("SELECT * FROM sizes ORDER BY id")
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

func Arrange(ids []int, ctgID int) ([]Size, error) {
	ctg, err := category.FindOne(ctgID)
	if err != nil {
		return nil, err
	}

	stm := "SELECT * FROM sizes"
	if ctg.ID > 0 {
		if ctg.ParentID == 3 || ctg.ParentID == 4 {
			return []Size{}, nil
		} else if ctg.ParentID == 1 {
			stm = stm + " WHERE type=0"
		} else if ctg.ParentID == 2 {
			stm = stm + " WHERE type=1"
		}
	}
	stm = stm + " ORDER BY id"

	rows, err := config.DB.Query(stm)
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

		for _, id := range ids {
			if id == size.ID {
				size.Checked = true
				break
			}
		}

		sizes = append(sizes, size)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sizes, nil
}
