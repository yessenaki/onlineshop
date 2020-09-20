package product

import (
	"database/sql"
	"onlineshop/config"
)

// Product struct
type Product struct {
	ID         int     `db:"id"`
	Title      string  `db:"title"`
	Price      float64 `db:"price"`
	OldPrice   float64 `db:"old_price"`
	Gender     int     `db:"gender"`
	IsKids     int     `db:"is_kids"`
	IsNew      int     `db:"is_new"`
	IsDiscount int     `db:"is_discount"`
	DscPercent int     `db:"dsc_percent"`
	BrandID    int     `db:"brand_id"`
	ColorID    int     `db:"color_id"`
	CategoryID int     `db:"category_id"`
	SizeID     int     `db:"size_id"`
	CreatedAt  string  `db:"created_at"`
	UpdatedAt  string  `db:"updated_id"`
}

func (p *Product) store() (int, error) {
	var lastInsertedID int
	sqlStatement := `INSERT INTO products (title, price, old_price, gender, is_kids, is_new, is_discount,
		dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id`
	err := config.DB.QueryRow(sqlStatement, p.Title, p.Price, p.OldPrice, p.Gender, p.IsKids, p.IsNew, p.IsDiscount,
		p.DscPercent, p.BrandID, p.ColorID, p.CategoryID, p.SizeID).Scan(&lastInsertedID)
	if err != nil {
		return lastInsertedID, err
	}
	return lastInsertedID, nil
}

func (p *Product) update() error {
	sqlStatement := `UPDATE products SET title=$1, price=$2, old_price=$3, gender=$4, is_kids=$5, is_new=$6, is_discount=$7,
		dsc_percent=$8, brand_id=$9, color_id=$10, category_id=$11, size_id=$12, updated_at=NOW()::timestamp(0) WHERE id=$13`
	_, err := config.DB.Exec(sqlStatement, p.Title, p.Price, p.OldPrice, p.Gender, p.IsKids, p.IsNew, p.IsDiscount,
		p.DscPercent, p.BrandID, p.ColorID, p.CategoryID, p.SizeID, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) destroy() error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id=$1", p.ID)
	if err != nil {
		return err
	}
	return nil
}

func allProducts() ([]Product, error) {
	rows, err := config.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prods := []Product{}
	for rows.Next() {
		prod := Product{}
		err := rows.Scan(
			&prod.ID, &prod.Title, &prod.Price, &prod.OldPrice, &prod.Gender, &prod.IsKids, &prod.IsNew, &prod.IsDiscount,
			&prod.DscPercent, &prod.BrandID, &prod.ColorID, &prod.CategoryID, &prod.SizeID, &prod.CreatedAt, &prod.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		prods = append(prods, prod)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prods, nil
}

func oneProduct(id int) (Product, error) {
	prod := Product{}
	err := config.DB.QueryRow("SELECT * FROM products WHERE id=$1", id).Scan(
		&prod.ID, &prod.Title, &prod.Price, &prod.OldPrice, &prod.Gender, &prod.IsKids, &prod.IsNew, &prod.IsDiscount,
		&prod.DscPercent, &prod.BrandID, &prod.ColorID, &prod.CategoryID, &prod.SizeID, &prod.CreatedAt, &prod.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return prod, nil
		}
		return prod, err
	}
	return prod, nil
}
