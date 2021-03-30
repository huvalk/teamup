package models

type Team struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Members UserArr `json:"members"`
	EventID int     `json:"eventid"`
}

//easyjson:json
type TeamArr []Team

type AddToTeam struct {
	TID int `json:"tid"`
	UID int `json:"uid"`
}

type AddToUser struct {
	UID1 int `json:"uid1"`
	UID2 int `json:"uid2"`
}

type TeamWinner struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	EventID int    `json:"eventid"`
	Prize   Prize  `json:"prize"`
}

//easyjson:json
type TeamWinnerArr []TeamWinner
