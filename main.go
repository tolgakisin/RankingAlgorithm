package main

import (
	"fmt"
	"math"
)

func main() {
	// Chess Elo Rating System
	var rA float64 = 1200
	var rB float64 = 1100
	var k float64 = 20

	calculateRanking(rA, rB, k)
}

func calculateRanking(winnerRate, loserRate, k float64) {
	var winnerPoint, loserPoint float64
	winnerPoint = calculateRankingPoint(loserRate, winnerRate)
	loserPoint = calculateRankingPoint(winnerRate, loserRate)

	winnerRate = winnerRate + (1-winnerPoint)*k
	loserRate = loserRate + (0-loserPoint)*k

	fmt.Printf("WinnerRate: %v , LoserRate: %v", winnerRate, loserRate)
}

func calculateRankingPoint(rating1, rating2 float64) float64 {
	return 1 / (1 + math.Pow(10, (rating1-rating2)/400))
}

// 1/(1+10^(rB-rA/400))
// referenced from: https://en.wikipedia.org/wiki/Elo_rating_system
