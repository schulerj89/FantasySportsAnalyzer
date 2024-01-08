package main

type Player struct {
	Position         string
	FullName         string
	ID               string
	RosterPosition   string
	Salary           int
	GameInfo         string
	TeamAbbrev       string
	AvgPointsPerGame float64

	// Extended player info
	Nationality  string `json:"strNationality"` 
	BirthDate    string `json:"dateBorn"`
	BirthPlace   string `json:"strBirthPlace"`
	Description  string `json:"strDescriptionEN"`
	Gender       string `json:"strGender"`
	Height       string `json:"strHeight"`
	Weight       string `json:"strWeight"`
	PlayerThumb  string `json:"strThumb"`
	PlayerCutout string `json:"strCutout"`
}

type APIResponse struct {
	Player []Player `json:"player"`
}
