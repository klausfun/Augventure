package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
	"github.com/redis/go-redis/v9"
)

type AuthPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewAuthPostgres(db *sqlx.DB, redis *redis.Client) *AuthPostgres {
	return &AuthPostgres{db: db, redis: redis}
}

func (r *AuthPostgres) CreateUser(user augventure.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email)"+
		" values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	jsonString, err := json.Marshal(augventure.Author{
		Id:       id,
		Username: user.Username,
		Email:    user.Email,
	})
	if err != nil {
		fmt.Printf("Failed to marshal: %s", err.Error())
		return id, err
	}

	personKey := fmt.Sprintf("person:%d", id)
	err = r.redis.Set(context.Background(), personKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
	}

	return id, err
}

func (r *AuthPostgres) GetUser(password, email string) (augventure.Author, error) {
	var user augventure.Author
	query := fmt.Sprintf("SELECT id, name, username, email, pfp_url, bio FROM %s"+
		" WHERE password_hash=$1 AND email=$2", userTable)
	err := r.db.Get(&user, query, password, email)

	return user, err
}
