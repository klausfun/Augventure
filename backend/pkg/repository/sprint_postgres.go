package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *SprintPostgres) Update(input augventure.UpdateSprintInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	var stateId int
	queryGetId := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", sprintStatesTable)
	err := r.db.Get(&stateId, queryGetId, input.Status)
	if err != nil {
		return err
	}
	setValues = append(setValues, fmt.Sprintf("state_id=$%d", argId))
	args = append(args, stateId)
	argId++

	if input.SuggestionWinnerId != nil {
		setValues = append(setValues, fmt.Sprintf("suggestion_winner_id=$%d", argId))
		args = append(args, *input.SuggestionWinnerId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND suggestion_winner_id=0",
		sprintsTable, setQuery, argId)
	args = append(args, input.SprintId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err = r.db.Exec(query, args...)

	return err
}
