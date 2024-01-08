package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getPlayerData(playerName string) Player {
	// playername needs _ instead of spaces
	playerName = strings.Replace(playerName, " ", "_", -1)
	sportsDbUrl := "https://thesportsdb.com/api/v1/json/3/searchplayers.php?p=%s"

	url := fmt.Sprintf(sportsDbUrl, playerName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Player{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Player{}
	}

	var apiResponse APIResponse
	json.Unmarshal(body, &apiResponse)

	if len(apiResponse.Player) == 0 {
		fmt.Println("No player found")
		return Player{}
	}

	player := apiResponse.Player[0]
	return player
}

func ParseCSV(reader *csv.Reader) ([]Player, error) {
	var players []Player

	// Read the header line
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount || err == csv.ErrBareQuote || err == csv.ErrQuote {
				continue
			}
			break
		}

		player, err := parseRecord(record)
		if err != nil {
			continue
		}

		players = append(players, player)
	}

	for index, player := range players {
		apiData := getPlayerData(player.FullName)
		player = mergePlayerData(player, apiData)
		players[index] = player
	}

	return players, nil
}

func mergePlayerData(player Player, apiData Player) Player {
	player.Nationality = apiData.Nationality
	player.BirthDate = apiData.BirthDate
	player.BirthPlace = apiData.BirthPlace
	player.Description = apiData.Description
	player.Gender = apiData.Gender
	player.Height = apiData.Height
	player.Weight = apiData.Weight
	player.PlayerThumb = apiData.PlayerThumb
	player.PlayerCutout = apiData.PlayerCutout

	return player
}

func parseRecord(record []string) (Player, error) {
	if len(record) < 9 {
		return Player{}, csv.ErrFieldCount
	}

	salary, err := strconv.Atoi(strings.TrimSpace(record[5]))
	if err != nil {
		return Player{}, err
	}

	avgPoints, err := strconv.ParseFloat(strings.TrimSpace(record[8]), 64)
	if err != nil {
		return Player{}, err
	}

	return Player{
		Position:         strings.TrimSpace(record[0]),
		FullName:         strings.TrimSpace(record[2]),
		ID:               strings.TrimSpace(record[3]),
		RosterPosition:   strings.TrimSpace(record[4]),
		Salary:           salary,
		GameInfo:         strings.TrimSpace(record[6]),
		TeamAbbrev:       strings.TrimSpace(record[7]),
		AvgPointsPerGame: avgPoints,
	}, nil
}
