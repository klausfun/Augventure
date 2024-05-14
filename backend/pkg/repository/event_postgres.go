package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) Create(userId int, event augventure.Event) (int, error) {
	var id int
	createEventQuery := fmt.Sprintf("INSERT INTO %s (title, description, start_date, author_id)"+
		"VALUES ($1, $2, $3, $4) RETURNING id", eventsTable)
	row := r.db.QueryRow(createEventQuery, event.Title, event.Description, event.Start, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
