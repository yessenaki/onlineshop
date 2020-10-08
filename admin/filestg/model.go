package filestg

import (
	"onlineshop/config"
	"time"
)

// Filestg = file storage
type Filestg struct {
	ID        int
	Name      string
	Path      string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (fs Filestg) Store() (int, error) {
	var id int
	stm := `INSERT INTO files (name, path, product_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id`
	err := config.DB.QueryRow(stm, fs.Name, fs.Path, fs.ProductID).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
