package main

import (
	"fmt"
)

func ProcessData(players []Player) {
	avgPointsPerTeam := calculateAvgPointsPerTeam(players)
	for team, avgPoints := range avgPointsPerTeam {
		fmt.Printf("Team: %s, Average Points: %.2f\n", team, avgPoints)
	}
}

func calculateAvgPointsPerTeam(players []Player) map[string]float64 {
	teamPoints := make(map[string]float64)
	teamCounts := make(map[string]int)

	for _, player := range players {
		teamPoints[player.TeamAbbrev] += player.AvgPointsPerGame
		teamCounts[player.TeamAbbrev]++
	}

	for team := range teamPoints {
		teamPoints[team] /= float64(teamCounts[team])
	}

	return teamPoints
}
