package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // hide password in JSON response
}

var UserStore = map[string]*User{} // In-memory store
