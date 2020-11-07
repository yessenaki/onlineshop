package brand

import (
	"database/sql"
	"github.com/yesseneon/onlineshop/config"
	"strings"
	"time"
)

// Brand struct
type Brand struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Errors    map[string]string
	Checked   bool
}

func (b *Brand) validate() bool {
	b.Errors = make(map[string]string)
	name := strings.TrimSpace(b.Name)

	if name == "" || len(name) > 30 {
		b.Errors["Name"] = "The field Name must be a string with a maximum length of 30"
	}

	return len(b.Errors) == 0
}

func (b *Brand) store() (int, error) {
	var lastInsertedID int
	stm := "INSERT INTO brands (name, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(stm, b.Name).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (b *Brand) update() error {
	_, err := config.DB.Exec("UPDATE brands SET name=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", b.Name, b.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b *Brand) destroy() error {
	_, err := config.DB.Exec("DELETE FROM brands WHERE id=$1", b.ID)
	if err != nil {
		return err
	}
	return nil
}

func AllBrands() ([]Brand, error) {
	rows, err := config.DB.Query("SELECT * FROM brands ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := []Brand{}
	for rows.Next() {
		brand := Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}

func oneBrand(id int) (Brand, error) {
	brand := Brand{}
	err := config.DB.QueryRow("SELECT * FROM brands WHERE id=$1", id).Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return brand, nil
		}
		return brand, err
	}
	return brand, nil
}

func Arrange(ids []int) ([]Brand, error) {
	rows, err := config.DB.Query("SELECT * FROM brands")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := []Brand{}
	for rows.Next() {
		brand := Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			return nil, err
		}

		for _, id := range ids {
			if id == brand.ID {
				brand.Checked = true
				break
			}
		}

		brands = append(brands, brand)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}
