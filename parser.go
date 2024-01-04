package main

import (
	"encoding/csv"
	"strconv"
	"strings"
)

// ParseCSV takes a csv.Reader and returns a slice of Player structs
func ParseCSV(reader *csv.Reader) ([]Player, error) {
	var players []Player

	// Read the header line
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Iterate over the rest of the lines
	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount || err == csv.ErrBareQuote || err == csv.ErrQuote {
				continue // Skip malformed lines
			}
			break
		}

		player, err := parseRecord(record)
		if err != nil {
			continue // Skip records that can't be parsed
		}

		players = append(players, player)
	}

	return players, nil
}

// parseRecord converts a record from the CSV into a Player struct
func parseRecord(record []string) (Player, error) {
	if len(record) < 9 { // Ensure there are enough fields
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
