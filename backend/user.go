package augventure

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	PfpUrl   string `json:"pfpURL"`
	Bio      string `json:"bio"`
}
