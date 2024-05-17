package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type SuggestionPostgres struct {
	db *sqlx.DB
}

func NewSuggestionPostgres(db *sqlx.DB) *SuggestionPostgres {
	return &SuggestionPostgres{db: db}
}

func (r *SuggestionPostgres) Create(userId int, suggestion augventure.Suggestion) (int, error) {
	var stateId int
	query := fmt.Sprintf("SELECT state_id FROM %s WHERE id = $1", sprintsTable)
	err := r.db.Get(&stateId, query, suggestion.SprintId)
	if err != nil {
		return 0, err
	}

	if stateId != 1 {
		return 0, errors.New("In this sprint, the voting is over!")
	}

	var id int
	createEventQuery := fmt.Sprintf("INSERT INTO %s (link_to_the_text, sprint_id, author_id)"+
		"VALUES ($1, $2, $3) RETURNING id", suggestionsTable)
	row := r.db.QueryRow(createEventQuery, suggestion.TextContent, suggestion.SprintId, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
