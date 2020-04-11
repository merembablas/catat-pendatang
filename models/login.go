package models

// Login represent login model
type Login struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
