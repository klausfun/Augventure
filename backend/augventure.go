package augventure

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
	Id      int    `json:"id"`
	EventId int    `json:"event_id"`
	State   string `json:"state"`
}

type Suggestion struct {
	Id          int    `json:"id"`
	SprintId    int    `json:"sprint_id"`
	TextContent string `json:"text_content"`
}
