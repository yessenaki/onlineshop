package tag

import (
	"database/sql"
	"github.com/yesseneon/onlineshop/config"
	"strings"
	"time"
)

// Tag struct
type Tag struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Errors    map[string]string
	Selected  bool
}

// PostTagItem struct
type PostTagItem struct {
	ID        int       `db:"id"`
	PostID    int       `db:"post_id"`
	TagID     int       `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	TagName   string    `db:"tag_name"`
}

func (t *Tag) validate() bool {
	t.Errors = make(map[string]string)
	name := strings.TrimSpace(t.Name)

	if name == "" || len(name) > 30 {
		t.Errors["Name"] = "The field Name must be a string with a maximum length of 30"
	}

	return len(t.Errors) == 0
}

func FindAll() ([]Tag, error) {
	rows, err := config.DB.Query("SELECT * FROM post_tags ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tags, nil
}

func FindWithSelected(postID int) ([]Tag, error) {
	items, err := FindPostTagItems(postID)
	if err != nil {
		return nil, err
	}

	rows, err := config.DB.Query("SELECT * FROM post_tags ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}

		for _, item := range items {
			if tag.ID == item.TagID {
				tag.Selected = true
				break
			}
		}

		tags = append(tags, tag)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tags, nil
}

func FindPostTagItems(postID int) ([]PostTagItem, error) {
	stm := `SELECT pti.*, pt.name AS tag_name FROM post_tag_items AS pti INNER JOIN post_tags AS pt ON pti.tag_id=pt.id WHERE pti.post_id=$1`
	rows, err := config.DB.Query(stm, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PostTagItem
	for rows.Next() {
		var item PostTagItem
		err := rows.Scan(&item.ID, &item.PostID, &item.TagID, &item.CreatedAt, &item.UpdatedAt, &item.TagName)
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

func findOne(id int) (Tag, error) {
	var tag Tag
	row := config.DB.QueryRow("SELECT * FROM post_tags WHERE id=$1", id)
	err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return tag, err
	}
	return tag, nil
}

func (t *Tag) store() (int, error) {
	var id int
	stm := "INSERT INTO post_tags (name, created_at, updated_at) VALUES ($1, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(stm, t.Name).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (t *Tag) update() error {
	_, err := config.DB.Exec("UPDATE post_tags SET name=$1, updated_at=NOW()::timestamp(0) WHERE id=$2", t.Name, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) destroy() error {
	_, err := config.DB.Exec("DELETE FROM post_tags WHERE id=$1", t.ID)
	if err != nil {
		return err
	}
	return nil
}
