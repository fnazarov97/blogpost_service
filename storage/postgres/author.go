package postgres

import (
	"blockpost/genprotos/author"
	"database/sql"
	"errors"
	"fmt"
)

// AddAuthor ...
func (p Postgres) AddAuthor(req *author.CreateAuthorReq) (res *author.CreateAuthorRes, err error) {
	fmt.Printf("%v %v", req.Fullname, req.ID)
	_, err = p.DB.Exec(`Insert into author(id, fullname, created_at) 
							VALUES($1,$2,now())`, req.ID, req.Fullname)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetAuthorByID ...
func (p Postgres) GetAuthorByID(req *author.Id) (*author.GetAuthorByIdRes, error) {
	result := &author.GetAuthorByIdRes{
		Articles: make([]*author.Article, 0),
	}
	var (
		updated_at sql.NullString
		deleted_at sql.NullString
	)
	row := p.DB.QueryRow("SELECT id, created_at, updated_at, deleted_at, fullname FROM author WHERE id=$1", req.Id)
	err := row.Scan(&result.Id, &result.CreatedAt, &updated_at, &deleted_at, &result.Fullname)
	if updated_at.Valid {
		result.UpdatedAt = updated_at.String
	}
	if deleted_at.Valid {
		result.DeletedAt = deleted_at.String
	}
	if err != nil {
		return result, err
	}
	ars, err := p.GetArticlesByAuthorID(req)
	result.Articles = ars.Articles
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetArticlesByAuthorID ...
func (p Postgres) GetArticlesByAuthorID(req *author.Id) (*author.GetArticles, error) {
	resp := &author.GetArticles{
		Articles: make([]*author.Article, 0),
	}
	rows, err := p.DB.Queryx(`SELECT 
									 id, 
									 title, 
									 body, 
									 author_id, 
									 created_at,
									 updated_at,
									 deleted_at
							FROM article
							WHERE author_id = $1 `, req.Id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			update_at  sql.NullString
			deleted_at sql.NullString
		)
		row := author.Article{
			Content: &author.Post{},
		}
		err := rows.Scan(&row.Id, &row.Content.Title, &row.Content.Body,
			&row.AuthorId, &row.CreatedAt, &update_at, &deleted_at)
		if err != nil {
			return resp, err
		}
		if update_at.Valid {
			row.UpdatedAt = update_at.String
		}
		if deleted_at.Valid {
			row.DeletedAt = deleted_at.String
		}
		resp.Articles = append(resp.Articles, &row)
	}
	return resp, nil
}

// GetAuthorList ...
func (p Postgres) GetAuthorList(req *author.GetAuthorListReq) (*author.GetAuthors, error) {
	resp := &author.GetAuthors{
		Authors: make([]*author.Author, 0),
	}
	rows, err := p.DB.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at
	FROM author WHERE deleted_at IS NULL AND (fullname ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, req.Search, int(req.Limit), int(req.Offset))
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			a         author.Author
			update_at sql.NullString
		)

		err := rows.Scan(
			&a.Id,
			&a.Fullname,
			&a.CreatedAt,
			&update_at,
		)
		if err != nil {
			return resp, err
		}
		if update_at.Valid {
			a.UpdatedAt = update_at.String
		}
		resp.Authors = append(resp.Authors, &a)
	}

	return resp, nil
}

// UpdateAuthor ...
func (p Postgres) UpdateAuthor(req *author.UpdateAuthorReq) error {
	res, err := p.DB.NamedExec("UPDATE author  SET fullname=:f, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": req.Id,
		"f":  req.Fullname,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}
	return errors.New("author not found")
}

// DeleteAuthor ...
func (p Postgres) DeleteAuthor(req *author.Id) error {
	res, err := p.DB.Exec("UPDATE author  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", req.Id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("author had been deleted already")
}
