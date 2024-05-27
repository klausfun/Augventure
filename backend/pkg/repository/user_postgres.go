package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) GetById(userId int) (augventure.Author, error) {
	var user augventure.Author
	query := fmt.Sprintf("SELECT bio, email, pfp_url, username FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *ProfilePostgres) UpdatePassword(userId int, input augventure.UpdatePasswordInput) error {
	query := fmt.Sprintf("UPDATE %s SET password_hash = $1 WHERE id = $2 AND password_hash = $3", userTable)
	_, err := r.db.Exec(query, input.NewPassword, userId, input.OldPassword)

	return err
}
