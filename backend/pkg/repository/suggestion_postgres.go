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
	createSuggestionQuery := fmt.Sprintf("INSERT INTO %s (link_to_the_text, sprint_id, author_id)"+
		"VALUES ($1, $2, $3) RETURNING id", suggestionsTable)
	row := r.db.QueryRow(createSuggestionQuery, suggestion.TextContent, suggestion.SprintId, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	LastId = id
	return id, nil
}

func (r *SuggestionPostgres) GetBySprintId(sprintId int) ([]augventure.FilterSuggestions, error) {
	var suggestions []augventure.FilterSuggestions

	query := fmt.Sprintf("SELECT * FROM %s WHERE sprint_id = $1", suggestionsTable)
	err := r.db.Select(&suggestions, query, sprintId)
	if err != nil {
		return nil, err
	}

	for i, curSuggestion := range suggestions {
		var author augventure.AuthorSuggestion

		queryAuthor := fmt.Sprintf("SELECT id, name, username, email, bio, pfp_url FROM %s WHERE id = $1", userTable)
		err := r.db.Get(&author, queryAuthor, curSuggestion.AuthorId)
		if err != nil {
			return nil, err
		}

		suggestions[i].Author = author
	}

	return suggestions, err
}

func (r *SuggestionPostgres) Vote(voteType bool, suggestionId, userId int) (int, error) {
	count := 0
	if voteType {
		count = 1
	} else {
		count = -1
	}

	var voteId int
	queryGetVoteId := fmt.Sprintf("SELECT id FROM %s WHERE suggestion_id = $1 AND user_id = $2", votesTable)
	err := r.db.Get(&voteId, queryGetVoteId, suggestionId, userId)

	tx, err2 := r.db.Begin()
	if err2 != nil {
		return 0, err2
	}
	if err != nil {
		createVotesQuery := fmt.Sprintf("INSERT INTO %s (user_id, suggestion_id, vote_type)"+
			" VALUES ($1, $2, $3)", votesTable)
		_, err = tx.Exec(createVotesQuery, userId, suggestionId, voteType)
		if err != nil {
			fmt.Println(createVotesQuery)
			tx.Rollback()
			return 0, err
		}

		query := fmt.Sprintf("UPDATE %s sug "+
			"SET votes = sug.votes + $1 "+
			"FROM %s vot "+
			"INNER JOIN %s us ON us.id = vot.user_id "+
			"WHERE vot.suggestion_id = sug.id AND sug.id = $2 AND us.id = $3 AND vot.vote_type = $4",
			suggestionsTable, votesTable, userTable)
		_, err = tx.Exec(query, count, suggestionId, userId, voteType)
		if err != nil {
			fmt.Println(query)
			tx.Rollback()
			return 0, err
		}
	} else {
		queryUpdateSuggestions := fmt.Sprintf("UPDATE %s sug SET votes = sug.votes + 2*$1"+
			" FROM %s vot"+
			" INNER JOIN %s us on us.id = vot.user_id"+
			" WHERE vot.suggestion_id = sug.id AND sug.id = $2 AND us.id = $3 AND vot.vote_type = $4",
			suggestionsTable, votesTable, userTable)
		_, err = tx.Exec(queryUpdateSuggestions, count, suggestionId, userId, !voteType)
		if err != nil {
			fmt.Println(queryUpdateSuggestions)
			tx.Rollback()
			return 0, err
		}

		queryUpdateVotes := fmt.Sprintf("UPDATE %s vot SET vote_type = NOT vot.vote_type"+
			" FROM %s sug, %s us"+
			" WHERE vot.suggestion_id = sug.id AND vot.user_id = us.id AND sug.id = $1 AND us.id = $2 AND vot.vote_type = $3",
			votesTable, suggestionsTable, userTable)
		_, err = tx.Exec(queryUpdateVotes, suggestionId, userId, !voteType)
		if err != nil {
			fmt.Println(queryUpdateVotes)
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()

	var votes int
	queryVotes := fmt.Sprintf("SELECT votes FROM %s WHERE id = $1", suggestionsTable)
	err = r.db.Get(&votes, queryVotes, suggestionId)

	return votes, err
}
