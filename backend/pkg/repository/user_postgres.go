package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
	"github.com/redis/go-redis/v9"
	"log"
)

type ProfilePostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewProfilePostgres(db *sqlx.DB, redis *redis.Client) *ProfilePostgres {
	return &ProfilePostgres{db: db, redis: redis}
}

func (r *ProfilePostgres) GetById(userId int) (augventure.Author, error) {
	var user augventure.Author

	personKey := fmt.Sprintf("person:%d", userId)
	val, err := r.redis.Get(context.Background(), personKey).Result()
	if err != nil {
		fmt.Printf("Failed to get value from redis: %s", err.Error())

		query := fmt.Sprintf("SELECT bio, email, pfp_url, username FROM %s WHERE id = $1", userTable)
		err = r.db.Get(&user, query, userId)

		return user, err
	}

	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
		return user, err
	}

	return user, err
}

func (r *ProfilePostgres) UpdatePassword(userId int, input augventure.UpdatePasswordInput) error {
	query := fmt.Sprintf("UPDATE %s SET password_hash = $1 WHERE id = $2 AND password_hash = $3", userTable)
	_, err := r.db.Exec(query, input.NewPassword, userId, input.OldPassword)

	return err
}
