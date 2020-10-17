package post

import (
	"database/sql"
	"fmt"
	"net/http"
	"onlineshop/admin/post/category"
	"onlineshop/admin/post/tag"
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
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	ImageName    string    `db:"image_name"`
	Author       string    `db:"author"`
	CategoryName string    `db:"category_name"`
	Tags         []string
	Errors       map[string]string
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

func FindAll() ([]Post, error) {
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

func FindWithLimit(load int, ctgID int, tagID int) ([]Post, error) {
	stm := "SELECT t2.* FROM post_tag_items AS t1 INNER JOIN posts AS t2 ON t1.post_id=t2.id"

	if ctgID > 0 {
		stm = fmt.Sprintf(stm+" WHERE t2.category_id=%d", ctgID)
	}

	if tagID > 0 {
		if ctgID > 0 {
			stm = fmt.Sprintf(stm+" AND t1.tag_id=%d", tagID)
		} else {
			stm = fmt.Sprintf(stm+" WHERE t1.tag_id=%d", tagID)
		}
	}

	stm = stm + " GROUP BY t2.id ORDER BY t2.id DESC LIMIT $1 OFFSET $2"

	perLoad := 1
	offset := (load - 1) * perLoad
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

func FindOne(id int) (Post, error) {
	var post Post
	stm := `SELECT t1.*, t2.name as category_name FROM posts AS t1
		INNER JOIN post_categories AS t2 ON t1.category_id=t2.id
		WHERE t1.id=$1`
	row := config.DB.QueryRow(stm, id)
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.ImagePath, &post.CategoryID, &post.CreatedAt, &post.UpdatedAt, &post.ImageName, &post.Author, &post.CategoryName)
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

func FindTags(postID int) ([]tag.Tag, error) {
	stm := `SELECT t2.* FROM post_tag_items AS t1
		INNER JOIN post_tags AS t2 ON t1.tag_id=t2.id
		WHERE t1.post_id=$1`
	rows, err := config.DB.Query(stm, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []tag.Tag
	for rows.Next() {
		var tag tag.Tag
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

func FindCategories() ([]category.Category, error) {
	stm := `SELECT t1.*, count(t1.id) AS post_qnt FROM post_categories AS t1
		INNER JOIN posts AS t2 ON t1.id=t2.category_id
		GROUP BY t1.id
		ORDER BY t1.name`
	rows, err := config.DB.Query(stm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ctgs []category.Category
	for rows.Next() {
		var ctg category.Category
		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.CreatedAt, &ctg.UpdatedAt, &ctg.PostQnt)
		if err != nil {
			return nil, err
		}

		ctgs = append(ctgs, ctg)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ctgs, nil
}
