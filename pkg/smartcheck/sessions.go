package main

import "time"

type ResponseCreateSession struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	User struct {
		ID                     string    `json:"id"`
		Href                   string    `json:"href"`
		UserID                 string    `json:"userID"`
		Name                   string    `json:"name"`
		Description            string    `json:"description"`
		Role                   string    `json:"role"`
		PasswordChangeRequired bool      `json:"passwordChangeRequired"`
		Created                time.Time `json:"created"`
		Updated                time.Time `json:"updated"`
	} `json:"user"`
	Token           string    `json:"token"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
	Expires         time.Time `json:"expires"`
	RoleSessionName string    `json:"roleSessionName"`
}
