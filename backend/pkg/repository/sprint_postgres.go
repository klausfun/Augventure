package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SprintPostgres struct {
	db *sqlx.DB
}

func NewSprintPostgres(db *sqlx.DB) *SprintPostgres {
	return &SprintPostgres{db: db}
}

func (r *SprintPostgres) Create(eventId int) (int, error) {
	var id int
	createSprintQuery := fmt.Sprintf("INSERT INTO %s (event_id)"+
		"VALUES ($1) RETURNING id", sprintsTable)
	row := r.db.QueryRow(createSprintQuery, eventId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
