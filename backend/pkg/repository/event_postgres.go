package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *EventPostgres) GetAll() ([]augventure.Event, error) {
	var events []augventure.Event
	query := fmt.Sprintf("SELECT * FROM %s", eventsTable)
	err := r.db.Select(&events, query)

	return events, err
}

func (r *EventPostgres) GetById(eventId int) (augventure.Event, error) {
	var event augventure.Event
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", eventsTable)
	err := r.db.Get(&event, query, eventId)

	return event, err
}

func (r *EventPostgres) Delete(userId, eventId int) error {
	var id = -1
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND author_id = $2 RETURNING id", eventsTable)
	row := r.db.QueryRow(query, eventId, userId)
	if err := row.Scan(&id); err != nil {
		return err
	}
	if id == -1 {
		errors.New("there is no event with this id or you do not have the authority to delete it")
	}

	return nil
}

func (r *EventPostgres) Update(userId, eventId int, input augventure.UpdateEventInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND author_id = $%d RETURNING id",
		eventsTable, setQuery, argId, argId+1)
	args = append(args, eventId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	var id = -1
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		return err
	}
	if id == -1 {
		errors.New("there is no event with this id or you do not have the authority to delete it")
	}

	return nil
}

func (r *EventPostgres) FinishVoting(userId, eventId int) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT spr.id FROM %s spr"+
		" INNER JOIN %s ev on spr.event_id = ev.id"+
		" INNER JOIN %s us on us.id = ev.author_id "+
		" WHERE ev.id = $1 AND us.id = $2"+
		" ORDER BY spr.id DESC LIMIT 1", sprintsTable, eventsTable, userTable)
	err := r.db.Get(&id, query, eventId, userId)

	return id, err
}
