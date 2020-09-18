package clothessize

import (
	"database/sql"
	"onlineshop/config"
)

// Clothessize struct
type Clothessize struct {
	ID        int    `db:"id"`
	Size      string `db:"size"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (cth *Clothessize) store() (int, error) {
	var lastInsertedID int
	sqlStatement := "INSERT INTO clothes_sizes (size, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, cth.Size).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (cth *Clothessize) update() error {
	_, err := config.DB.Exec("UPDATE clothes_sizes SET size=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", cth.Size, cth.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cth *Clothessize) destroy() error {
	_, err := config.DB.Exec("DELETE FROM clothes_sizes WHERE id=$1", cth.ID)
	if err != nil {
		return err
	}
	return nil
}

func allClothessizes() ([]Clothessize, error) {
	rows, err := config.DB.Query("SELECT * FROM clothes_sizes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cths := []Clothessize{}
	for rows.Next() {
		cth := Clothessize{}
		err := rows.Scan(&cth.ID, &cth.Size, &cth.CreatedAt, &cth.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cths = append(cths, cth)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cths, nil
}

func oneClothessize(id int) (Clothessize, error) {
	cth := Clothessize{}
	err := config.DB.QueryRow("SELECT * FROM clothes_sizes WHERE id=$1", id).Scan(&cth.ID, &cth.Size, &cth.CreatedAt, &cth.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return cth, nil
		}
		return cth, err
	}
	return cth, nil
}
