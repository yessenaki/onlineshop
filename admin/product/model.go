package product

import (
	"database/sql"
	"fmt"
	"net/http"
	"onlineshop/config"
	"onlineshop/helper"
	"strconv"
	"strings"
	"time"
)

// Product struct
type Product struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Price       int       `db:"price"`
	OldPrice    int       `db:"old_price"`
	Gender      int       `db:"gender"`
	IsKids      int       `db:"is_kids"`
	IsNew       int       `db:"is_new"`
	IsDiscount  int       `db:"is_discount"`
	DscPercent  int       `db:"dsc_percent"`
	BrandID     int       `db:"brand_id"`
	ColorID     int       `db:"color_id"`
	CategoryID  int       `db:"category_id"`
	SizeID      int       `db:"size_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_id"`
	Image       string    `db:"image"`
	ImageName   string    `db:"image_name"`
	Description string    `db:"description"`
	BrandName   string    `db:"brand_name"`
	ColorName   string    `db:"color_name"`
	CtgName     string    `db:"ctg_name"`
	SizeName    string    `db:"size_name"`
	Errors      map[string]string
}

func (p *Product) validate(r *http.Request) bool {
	p.Errors = make(map[string]string)
	title := strings.TrimSpace(p.Title)
	descr := strings.TrimSpace(p.Description)

	if title == "" || len(title) > 50 {
		p.Errors["Title"] = "The field Title must be a string with a maximum length of 50"
	}

	if p.Price == 0 || p.Price >= 1000000 {
		p.Errors["Price"] = "Price must be more than 0.00 and less than 10000.00"
	}

	if p.IsDiscount == 1 {
		if p.OldPrice == 0 || p.OldPrice >= 1000000 {
			p.Errors["OldPrice"] = "Old price must be more than 0.00 and less than 10000.00"
		}

		if p.DscPercent == 0 || p.DscPercent > 100 {
			p.Errors["DscPercent"] = "The discount percent must be more than 0% and less than 101%"
		}
	}

	if descr == "" {
		p.Errors["Description"] = "The field Description cannot be empty"
	}

	_, fileHeader, err := r.FormFile("image")
	exts := []string{"png", "jpg", "jpeg"}
	if err == http.ErrMissingFile {
		if r.Method == http.MethodPost {
			p.Errors["Image"] = "Please choose an image file"
		}
	} else {
		ext := strings.Split(fileHeader.Filename, ".")[1]
		if fileHeader.Size > 2<<20 || helper.Contains(exts, ext) == false {
			p.Errors["Image"] = "Your image must be in png, jpg, jpeg format and must not exceed 2MB"
		}
	}

	return len(p.Errors) == 0
}

func (p *Product) store() (int, error) {
	var id int
	stm := `INSERT INTO products (title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent,
		brand_id,color_id,category_id, size_id, created_at, updated_at, image, image_name, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW()::timestamp(0), NOW()::timestamp(0), $13, $14, $15) RETURNING id`
	err := config.DB.QueryRow(stm,
		p.Title, p.Price, p.OldPrice, p.Gender, p.IsKids, p.IsNew, p.IsDiscount, p.DscPercent, p.BrandID, p.ColorID, p.CategoryID, p.SizeID, p.Image, p.ImageName, p.Description,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (p *Product) update() error {
	stm := `UPDATE products SET title=$1, price=$2, old_price=$3, gender=$4, is_kids=$5, is_new=$6, is_discount=$7, dsc_percent=$8, brand_id=$9,
		color_id=$10, category_id=$11, size_id=$12, updated_at=NOW()::timestamp(0), image=$13, image_name=$14, description=$15 WHERE id=$16`
	_, err := config.DB.Exec(stm,
		p.Title, p.Price, p.OldPrice, p.Gender, p.IsKids, p.IsNew, p.IsDiscount, p.DscPercent, p.BrandID, p.ColorID, p.CategoryID, p.SizeID, p.Image, p.ImageName, p.Description, p.ID,
	)
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

func FindAll() ([]Product, error) {
	stm := `SELECT p.*, b.name as brand_name, c.name as color_name, ctg.name as ctg_name, COALESCE(s.size, 'None', s.size) as size_name
		FROM products as p
		INNER JOIN brands as b
		ON p.brand_id=b.id
		INNER JOIN colors as c
		ON p.color_id=c.id
		INNER JOIN categories as ctg
		ON p.category_id=ctg.id
		LEFT JOIN sizes as s
		ON p.size_id=s.id`

	rows, err := config.DB.Query(stm)
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
			&prod.Image, &prod.ImageName, &prod.Description, &prod.BrandName, &prod.ColorName, &prod.CtgName, &prod.SizeName,
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

func FindOne(id int) (Product, error) {
	stm := `SELECT p.*, b.name as brand_name, c.name as color_name, ctg.name as ctg_name, COALESCE(s.size, 'None', s.size) as size_name
		FROM products as p
		INNER JOIN brands as b
		ON p.brand_id=b.id
		INNER JOIN colors as c
		ON p.color_id=c.id
		INNER JOIN categories as ctg
		ON p.category_id=ctg.id
		LEFT JOIN sizes as s
		ON p.size_id=s.id
		WHERE p.id=$1`
	prod := Product{}
	err := config.DB.QueryRow(stm, id).Scan(
		&prod.ID, &prod.Title, &prod.Price, &prod.OldPrice, &prod.Gender, &prod.IsKids, &prod.IsNew, &prod.IsDiscount,
		&prod.DscPercent, &prod.BrandID, &prod.ColorID, &prod.CategoryID, &prod.SizeID, &prod.CreatedAt, &prod.UpdatedAt,
		&prod.Image, &prod.ImageName, &prod.Description, &prod.BrandName, &prod.ColorName, &prod.CtgName, &prod.SizeName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return prod, nil
		}
		return prod, err
	}
	return prod, nil
}

func FindByParams(params map[string]interface{}) ([]Product, int, error) {
	var stm string
	if params["ctgID"] != 0 {
		stm = fmt.Sprintf(stm+" AND category_id=%d", params["ctgID"])
	}

	if params["brands"] != "" {
		ids := arrangeList(params["brands"].(string))
		stm = fmt.Sprintf(stm+" AND brand_id IN (%s)", ids)
	}

	if params["sizes"] != "" {
		ids := arrangeList(params["sizes"].(string))
		stm = fmt.Sprintf(stm+" AND size_id IN (%s)", ids)
	}

	if params["colors"] != "" {
		ids := arrangeList(params["colors"].(string))
		stm = fmt.Sprintf(stm+" AND color_id IN (%s)", ids)
	}

	var quantity int
	err := config.DB.QueryRow("SELECT count(*) FROM products WHERE gender=$1 AND is_kids=$2"+stm, params["gender"], params["isKids"]).Scan(&quantity)
	if err != nil {
		return nil, 0, err
	}

	stm = stm + " ORDER BY id DESC"

	resultsPerPage := 9
	offset := (params["page"].(int) - 1) * resultsPerPage
	stm = fmt.Sprintf(stm+" LIMIT %d OFFSET %d", resultsPerPage, offset)

	rows, err := config.DB.Query("SELECT * FROM products WHERE gender=$1 AND is_kids=$2"+stm, params["gender"], params["isKids"])
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	prods := []Product{}
	for rows.Next() {
		prod := Product{}
		err := rows.Scan(
			&prod.ID, &prod.Title, &prod.Price, &prod.OldPrice, &prod.Gender, &prod.IsKids, &prod.IsNew, &prod.IsDiscount, &prod.DscPercent,
			&prod.BrandID, &prod.ColorID, &prod.CategoryID, &prod.SizeID, &prod.CreatedAt, &prod.UpdatedAt, &prod.Image, &prod.ImageName, &prod.Description,
		)
		if err != nil {
			return nil, 0, err
		}
		prods = append(prods, prod)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return prods, quantity, nil
}

func arrangeList(l string) string {
	ids := helper.ListToSlice(l)
	var nl string
	for i := 0; i < len(ids); i++ {
		nl = nl + strconv.Itoa(ids[i])
		if i != len(ids)-1 {
			nl = nl + ","
		}
	}
	return nl
}
