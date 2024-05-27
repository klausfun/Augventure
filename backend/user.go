package augventure

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	PfpUrl   string `json:"pfp_url"`
	Bio      string `json:"bio"`
}

type Author struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	PfpUrl   string `json:"pfp_url" db:"pfp_url"`
	Bio      string `json:"bio" db:"bio"`
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (i UpdatePasswordInput) Validate() error {
	if len(i.NewPassword) < 8 {
		return errors.New("the password is too short")
	}

	return nil
}
