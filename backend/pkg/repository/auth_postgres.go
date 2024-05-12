package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user augventure.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email)"+
		" values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
