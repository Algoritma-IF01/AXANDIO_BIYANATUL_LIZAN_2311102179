package main

import (
	"fmt"
)

type Match struct {
	clubA, clubB string
	scoreA, scoreB int
	winner         string
}

func getInputClubs() (string, string) {
	var clubA, clubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)
	return clubA, clubB
}

func getMatchResult(matchNumber int) (int, int, bool) {
	var scoreA, scoreB int
	fmt.Printf("Pertandingan %d: ", matchNumber)
	fmt.Scan(&scoreA, &scoreB)
	if scoreA < 0 || scoreB < 0 {
		return scoreA, scoreB, false
	}
	return scoreA, scoreB, true
}

func determineWinner(clubA string, clubB string, scoreA int, scoreB int) string {
	if scoreA > scoreB {
		return clubA
	} else if scoreA < scoreB {
		return clubB
	} else {
		return "Draw"
	}
}

func printResults(matches []Match) {
	var i int
	for i = 0; i < len(matches); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, matches[i].winner)
	}
}

func printWinners(matches []Match) {
	fmt.Println("\nDaftar Klub yang Memenangkan Pertandingan:")
	var i int
	for i = 0; i < len(matches); i++ {
		if matches[i].winner != "Draw" {
			fmt.Println(matches[i].winner)
		}
	}
}

func main() {
	var clubA, clubB, winner string
	var matches []Match
	var matchNumber, scoreA, scoreB int
	var valid bool

	clubA, clubB = getInputClubs()
	matches = []Match{}
	matchNumber = 1

	for {
		scoreA, scoreB, valid = getMatchResult(matchNumber)
		if !valid {
			fmt.Println("Pertandingan Selesai")
			break
		}

		winner = determineWinner(clubA, clubB, scoreA, scoreB)
		matches = append(matches, Match{
			clubA:  clubA,
			clubB:  clubB,
			scoreA: scoreA,
			scoreB: scoreB,
			winner: winner,
		})

		matchNumber++
	}

	printResults(matches)
}
