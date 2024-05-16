package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type SprintPostgres struct {
	db *sqlx.DB
}

func NewSprintPostgres(db *sqlx.DB) *SprintPostgres {
	return &SprintPostgres{db: db}
}

func (r *SprintPostgres) Create(eventId int, sprint augventure.Sprint) (int, error) {
	var id int
	createSprintQuery := fmt.Sprintf("INSERT INTO %s (event_id)"+
		"VALUES ($1) RETURNING id", sprintsTable)
	if eventId != -1 {
		row := r.db.QueryRow(createSprintQuery, eventId)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	} else {
		row := r.db.QueryRow(createSprintQuery, sprint.EventId)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}
