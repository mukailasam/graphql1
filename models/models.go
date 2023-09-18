package models

import (
	"database/sql"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

type ModelInterface interface {
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type Model struct {
	DB ModelInterface
}

func (dbExec Model) CreateMessage(title, body string) error {
	sql := `INSERT INTO posts(title, body) VALUES($1, $2)`
	stmt, err := dbExec.DB.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title, body)
	if err != nil {
		return err
	}

	stmt.Close()
	return nil
}

func (dbExec Model) ReadMessage(id int) (interface{}, error) {
	p := Post{}
	sqlCmd := `SELECT title, body FROM posts WHERE id=$1`

	stmt, err := dbExec.DB.Prepare(sqlCmd)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&p.Title, &p.Body)
	if err != nil {
		return nil, err
	}

	stmt.Close()
	return p, nil

}

func (dbExec Model) ReadMessages() (interface{}, error) {
	mp := []Post{}

	sqlCmd := `SELECT id, title, body FROM posts`
	stmt, err := dbExec.DB.Prepare(sqlCmd)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp := Post{}
		err = rows.Scan(&sp.ID, &sp.Title, &sp.Body)
		if err != nil {
			return nil, err
		}

		mp = append(mp, sp)
	}

	stmt.Close()
	return mp, nil

}

func (dbExec Model) UpdateMessage(id int, title, body string) error {
	sqlCmd := `UPDATE posts SET title=$1, body=$2 WHERE id=$3`
	stmt, err := dbExec.DB.Prepare(sqlCmd)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title, body, id)
	if err != nil {
		return err
	}

	stmt.Close()
	return nil
}

func (dbExec Model) DeleteMessage(id int) error {
	sqlCmd := `DELETE FROM posts WHERE id=$1`
	stmt, err := dbExec.DB.Prepare(sqlCmd)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt.Close()
	return nil
}

func (dbExec Model) PostExist(id int) bool {
	var title string
	sqlCmd := `SELECT title FROM posts WHERE id=$1`
	stmt, err := dbExec.DB.Prepare(sqlCmd)
	err = stmt.QueryRow(id).Scan(&title)
	if err == sql.ErrNoRows {
		return false
	}

	return true
}
