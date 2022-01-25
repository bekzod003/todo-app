package repository

import (
	"fmt"
	todoapp "todo-app"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (p *AuthPostgres) CreateUser(user todoapp.User) (int, error) {
	query := `insert into %s (name, username, password_hash)
	values ($1, $2, $3) returning id`
	query = fmt.Sprintf(query, userTable)
	row := p.db.QueryRow(query, user.Name, user.Username, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
