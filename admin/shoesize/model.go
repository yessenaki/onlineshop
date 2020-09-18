package shoesize

import (
	"database/sql"
	"onlineshop/config"
)

// Shoesize struct
type Shoesize struct {
	ID        int    `db:"id"`
	Size      string `db:"size"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (sh *Shoesize) store() (int, error) {
	var lastInsertedID int
	sqlStatement := "INSERT INTO shoe_sizes (size, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, sh.Size).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (sh *Shoesize) update() error {
	_, err := config.DB.Exec("UPDATE shoe_sizes SET size=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", sh.Size, sh.ID)
	if err != nil {
		return err
	}
	return nil
}

func (sh *Shoesize) destroy() error {
	_, err := config.DB.Exec("DELETE FROM shoe_sizes WHERE id=$1", sh.ID)
	if err != nil {
		return err
	}
	return nil
}

func allShoesizes() ([]Shoesize, error) {
	rows, err := config.DB.Query("SELECT * FROM shoe_sizes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shs := []Shoesize{}
	for rows.Next() {
		sh := Shoesize{}
		err := rows.Scan(&sh.ID, &sh.Size, &sh.CreatedAt, &sh.UpdatedAt)
		if err != nil {
			return nil, err
		}
		shs = append(shs, sh)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return shs, nil
}

func oneShoesize(id int) (Shoesize, error) {
	sh := Shoesize{}
	err := config.DB.QueryRow("SELECT * FROM shoe_sizes WHERE id=$1", id).Scan(&sh.ID, &sh.Size, &sh.CreatedAt, &sh.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return sh, nil
		}
		return sh, err
	}
	return sh, nil
}
