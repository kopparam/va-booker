package main

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	FirstLogin  string `json:"first_login"`
	MemberID    string `json:"member_id"`
}

type bookableClass struct {
	BookingID     string
	ClassName     string
	Instructor    string
	StartDateTime string
}
