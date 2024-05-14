package augventure

type Event struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Start       string `json:"start" binding:"required"`
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
