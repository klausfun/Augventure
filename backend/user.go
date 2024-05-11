package augventure

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	PfpUrl   string `json:"pfpURL"`
	Bio      string `json:"bio"`
}
