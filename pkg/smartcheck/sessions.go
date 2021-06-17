package main

type Sessions struct {
	Id   string
	Href string
	User struct {
		Id                     string
		Href                   string
		UserID                 string
		Name                   string
		Description            string
		Role                   string
		PasswordChangeRequired string
		Created                string
		Updated                string
	}
	Token           string
	Created         string
	Updated         string
	Expires         string
	RoleSessionName string
}
