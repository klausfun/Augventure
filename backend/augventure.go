package augventure

import (
	"errors"
)

type Event struct {
	Id           int    `json:"id" db:"id"`
	Title        string `json:"title" db:"title" binding:"required"`
	Description  string `json:"description" db:"description" binding:"required"`
	Start        string `json:"start_date" db:"start_date" binding:"required"`
	PictureUrl   string `json:"picture_url" db:"picture_url"`
	AuthorId     int    `json:"author_id" db:"author_id"`
	StateId      int    `json:"state_id" db:"state_id"`
	CreationDate string `json:"creation_date" db:"creation_date"`
}

type Sprint struct {
	Id      int `json:"id"`
	EventId int `json:"event_id" binding:"required"`
}

type Suggestion struct {
	Id          int    `json:"id" db:"id"`
	SprintId    int    `json:"sprint_id" binding:"required"`
	TextContent string `json:"text_content" binding:"required"`
}

type SprintId struct {
	Id int `json:"sprint_id" binding:"required"`
}

type AuthorId struct {
	Id int `json:"author_id" binding:"required"`
}

type FilterSuggestions struct {
	Id       int              `json:"id" db:"id"`
	AuthorId int              `json:"author_id" db:"author_id"`
	SprintId int              `json:"sprint_id" db:"sprint_id"`
	Author   AuthorSuggestion `json:"author"`
	Content  string           `json:"content" db:"link_to_the_text"`
	PostDate string           `json:"post_date" db:"post_date"` // time.Time
	Votes    int              `json:"votes" db:"votes"`
}

type Vote struct {
	VoteType bool `json:"this_is_a_like"`
}

type UpdateSprintInput struct {
	SprintId           int    `json:"sprint_id"`
	SuggestionWinnerId *int   `json:"suggestion_winner_id"`
	Status             string `json:"status"`
}

type FinishVoting struct {
	SuggestionWinnerId *int `json:"suggestion_winner_id"`
}

type FinishImplementing struct {
	TextContent  string `json:"text_content"`
	IsLastSprint bool   `json:"is_last_sprint"`
}

type UpdateEventInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}

func (i UpdateEventInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
