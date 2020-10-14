package cart

import (
	"database/sql"
	"onlineshop/config"
	"time"
)

// Cart struct
type Cart struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Item struct
type Item struct {
	ID        int       `db:"id"`
	CartID    int       `db:"cart_id"`
	ProductID int       `db:"product_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
	Price     int       `db:"price"`
	ImagePath string    `db:"image_path"`
}

// UserCart struct
type UserCart struct {
	UserID    int `json:"user_id"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (uc UserCart) store() (bool, error) {
	var c Cart
	stm := "SELECT * FROM carts WHERE user_id=$1"
	err := config.DB.QueryRow(stm, uc.UserID).Scan(&c.ID, &c.UserID, &c.CreatedAt, &c.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	var i Item
	// if the cart is already exists
	if c.ID > 0 {
		stm := "SELECT * FROM cart_items WHERE cart_id=$1 AND product_id=$2"
		err := config.DB.QueryRow(stm, c.ID, uc.ProductID).Scan(&i.ID, &i.CartID, &i.ProductID, &i.Quantity, &i.CreatedAt, &i.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
			return false, err
		}

		// if the item is already in the cart
		if i.ID > 0 {
			return true, nil
		}
	} else {
		// if the cart doesn't exist, create a new one
		stm := "INSERT INTO carts (user_id, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
		err := config.DB.QueryRow(stm, uc.UserID).Scan(&c.ID)
		if err != nil {
			return false, err
		}
	}

	// create a new product
	stm = "INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err = config.DB.QueryRow(stm, c.ID, uc.ProductID, uc.Quantity).Scan(&i.ID)
	if err != nil {
		return false, err
	}

	return false, nil
}

func (uc *UserCart) changeQnt() ([]Item, error) {
	stm := "UPDATE cart_items SET quantity=$1, updated_at=NOW()::timestamp(0) WHERE cart_id=$2 AND product_id=$3"
	_, err := config.DB.Exec(stm, uc.Quantity, uc.CartID, uc.ProductID)
	if err != nil {
		return nil, err
	}

	stm = `SELECT ci.*, p.price AS price
		FROM cart_items AS ci
		INNER JOIN products AS p ON ci.product_id=p.id
		WHERE ci.cart_id=$1`

	rows, err := config.DB.Query(stm, uc.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt, &item.Price)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Cart) getItems() ([]Item, error) {
	var items []Item
	err := config.DB.QueryRow("SELECT * FROM carts WHERE user_id=$1", c.UserID).Scan(&c.ID, &c.UserID, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return items, nil
		}
		return nil, err
	}

	stm := `WITH images AS (SELECT DISTINCT ON (product_id) * FROM files ORDER BY product_id, id)
		SELECT ci.*, p.title AS title, p.price AS price, i.path AS image_path
		FROM cart_items AS ci
		INNER JOIN products AS p ON ci.product_id=p.id
		INNER JOIN images AS i ON p.id=i.product_id
		WHERE ci.cart_id=$1
		ORDER BY ci.id DESC`
	rows, err := config.DB.Query(stm, c.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt, &item.Title, &item.Price, &item.ImagePath)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemQuantity(userID int) (int, error) {
	var qnt int
	stm := "SELECT count(*) FROM cart_items WHERE cart_id=(SELECT id FROM carts WHERE user_id=$1 LIMIT 1)"
	err := config.DB.QueryRow(stm, userID).Scan(&qnt)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return qnt, nil
}

func (uc *UserCart) deleteItem() ([]Item, error) {
	_, err := config.DB.Exec("DELETE FROM cart_items WHERE cart_id=$1 AND product_id=$2", uc.CartID, uc.ProductID)
	if err != nil {
		return nil, err
	}

	var items []Item
	stm := `SELECT ci.id, ci.quantity, p.price AS price
		FROM cart_items AS ci
		INNER JOIN products AS p ON ci.product_id=p.id
		WHERE ci.cart_id=$1`
	rows, err := config.DB.Query(stm, uc.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
