package main

import "time"

type (
	RequestCreateSessionUserCredentials struct {
		UserID   string
		Password string
	}
	RequestCreateSessionUser struct {
		User RequestCreateSessionUserCredentials
	}

	RequestCreateSessionSamlCredentials struct {
		Response     string
		SelectedRole string
	}

	RequestCreateSessionSaml struct {
		Saml RequestCreateSessionSamlCredentials
	}
	// Not used:
	RequestCreateSession struct {
		User struct {
			UserID   string `json:"userID"`
			Password string `json:"password"`
		} `json:"user"`
		Saml struct {
			Response     string `json:"response"`
			SelectedRole string `json:"selectedRole"`
		} `json:"saml"`
	}

	ResponseCreateSession struct {
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
)
