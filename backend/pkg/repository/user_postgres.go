package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(userId int) (augventure.Author, error) {
	var user augventure.Author
	query := fmt.Sprintf("SELECT bio, email, pfp_url, username FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}
