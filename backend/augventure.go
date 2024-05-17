package augventure

import "errors"

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

const (
	Voting       = "voting"
	Implementing = "implementing"
	Ended        = "ended"
)

type Sprint struct {
	Id      int `json:"id"`
	EventId int `json:"event_id" binding:"required"`
}

type Suggestion struct {
	Id          int    `json:"id"`
	SprintId    int    `json:"sprint_id"`
	TextContent string `json:"text_content"`
}

type UpdateSprintInput struct {
	SprintId           int    `json:"sprint_id"`
	SuggestionWinnerId *int   `json:"suggestion_winner_id"`
	Status             string `json:"status"`
}

type FinishImplementing struct {
	TextContent  string `json:"text_content"`
	IsLastSprint bool   `json:"is_last_sprint"`
}

type UpdateEventInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateEventInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
