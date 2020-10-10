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
}

// UserCart struct
type UserCart struct {
	UserID    int `json:"user_id"`
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
	if c.ID > 0 {
		stm = "SELECT * FROM cart_items WHERE cart_id=$1 AND product_id=$2"
		err = config.DB.QueryRow(stm, c.ID, uc.ProductID).Scan(&i.ID, &i.CartID, &i.ProductID, &i.Quantity, &i.CreatedAt, &i.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
			return false, err
		}

		// if the item is already in the cart
		if i.ID > 0 {
			return true, nil
		}
	} else {
		stm = "INSERT INTO carts (user_id, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
		err = config.DB.QueryRow(stm, uc.UserID).Scan(&c.ID)
		if err != nil {
			return false, err
		}
	}

	stm = "INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err = config.DB.QueryRow(stm, c.ID, uc.ProductID, uc.Quantity).Scan(&i.ID)
	if err != nil {
		return false, err
	}

	return false, nil
}
