package post

import (
	"database/sql"
	"net/http"
	"onlineshop/config"
	"onlineshop/helper"
	"strings"
	"time"
)

// Post struct
type Post struct {
	ID           int       `db:"id"`
	Title        string    `db:"title"`
	Body         string    `db:"body "`
	ImagePath    string    `db:"image_path"`
	CategoryID   int       `db:"category_id"`
	CategoryName string    `db:"category_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	ImageName    string    `db:"image_name"`
	Author       string    `db:"author"`
	Errors       map[string]string
	Tags         []string
}

func (p *Post) validate(r *http.Request) (bool, error) {
	p.Errors = make(map[string]string)
	title := strings.TrimSpace(p.Title)
	body := strings.TrimSpace(p.Body)
	author := strings.TrimSpace(p.Author)

	if title == "" || len(title) > 255 {
		p.Errors["Title"] = "The field Title must be a string with a maximum length of 255"
	}

	if body == "" {
		p.Errors["Body"] = "The field Body must not be empty"
	}

	if len(p.Tags) == 0 {
		p.Errors["Tags"] = "The field Post tags must not be empty"
	}

	if author == "" {
		p.Errors["Author"] = "The field Author must not be empty"
	}

	_, fh, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			if r.Method == http.MethodPost {
				p.Errors["ImagePath"] = "The image must not be empty"
			}
		} else {
			return false, err
		}
	} else {
		exts := []string{"png", "jpg", "jpeg"}
		ext := strings.Split(fh.Filename, ".")[1]
		if fh.Size > 2<<20 || helper.Contains(exts, ext) == false {
			p.Errors["ImagePath"] = "Your image must be in png, jpg, jpeg format and must not exceed 2MB"
		}
	}

	return (len(p.Errors) == 0), nil
}

func findAll() ([]Post, error) {
	stm := `SELECT p.*, pc.name AS category_name FROM posts AS p INNER JOIN post_categories AS pc ON p.category_id=pc.id ORDER BY id DESC`
	rows, err := config.DB.Query(stm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.ImagePath, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.ImageName, &post.Author, &post.CategoryName)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func FindWithLimit(load int) ([]Post, error) {
	perLoad := 3
	offset := (load - 1) * perLoad
	stm := `SELECT * FROM posts ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := config.DB.Query(stm, perLoad, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.ImagePath, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.ImageName, &post.Author)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func findOne(id int) (Post, error) {
	var post Post
	row := config.DB.QueryRow("SELECT * FROM posts WHERE id=$1", id)
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.ImagePath, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.ImageName, &post.Author)
	if err != nil && err != sql.ErrNoRows {
		return post, err
	}
	return post, nil
}

func (p *Post) store() (int, error) {
	var id int
	stm := "INSERT INTO posts (title, body, image_path, category_id, created_at, updated_at, image_name, author) VALUES ($1, $2, $3, $4, NOW()::timestamp(0), NOW()::timestamp(0), $5, $6) RETURNING id"
	err := config.DB.QueryRow(stm, p.Title, p.Body, p.ImagePath, p.CategoryID, p.ImageName, p.Author).Scan(&id)
	if err != nil {
		return id, err
	}

	for i := 0; i < len(p.Tags); i++ {
		stm := "INSERT INTO post_tag_items (post_id, tag_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0))"
		_, err := config.DB.Exec(stm, id, p.Tags[i])
		if err != nil {
			return id, err
		}
	}

	return id, nil
}

func (p *Post) update() error {
	stm := "UPDATE posts SET title=$1, body=$2, image_path=$3, category_id=$4, updated_at=NOW()::timestamp(0), image_name=$5, author=$6 WHERE id=$7"
	_, err := config.DB.Exec(stm, p.Title, p.Body, p.ImagePath, p.CategoryID, p.ImageName, p.Author, p.ID)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("DELETE FROM post_tag_items WHERE post_id=$1", p.ID)
	if err != nil {
		return err
	}

	for i := 0; i < len(p.Tags); i++ {
		stm := "INSERT INTO post_tag_items (post_id, tag_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0))"
		_, err := config.DB.Exec(stm, p.ID, p.Tags[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Post) destroy() error {
	_, err := config.DB.Exec("DELETE FROM posts WHERE id=$1", p.ID)
	if err != nil {
		return err
	}
	return nil
}
