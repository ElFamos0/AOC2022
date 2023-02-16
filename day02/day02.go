/*

Author : Malo DAMIEN

Solution of the day02 of advent of code

*/

package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strings"
)

const (
	Rock = 1
	Paper = 2
	Scissors = 3
	Win = 6
	Loss = 0
	Draw = 3
)

func solve1() int {
	lignes, _ := readline("input.txt")
	
	// map des différentes possibilités 
	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	score := 0
	for _, ligne := range lignes {
		splitstring := strings.Split(ligne," ")
		// ajout au score en fonction du signe joué
		score += choices[splitstring[1]]
		// ajout au score en fonction de l'issue de la game - utilisation de switch case
		switch splitstring[1]{
		case "X":
			switch splitstring[0]{
			case "A":
				score += Draw
			case "B":
				score += Loss
			case "C":
				score += Win
			}
		case "Y":
			switch splitstring[0]{
			case "A":
				score += Win
			case "B":
				score += Draw
			case "C":
				score += Loss
		}
		case "Z":
			switch splitstring[0]{
			case "A":
				score += Loss
			case "B":
				score += Win
			case "C":
				score += Draw
	}
	}
}
	return score
}

func solve2() int {
	lignes, _ := readline("input.txt")
	roundscore := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}
	score := 0
	for _, ligne := range lignes {
		splitstring := strings.Split(ligne," ")
		// ajout au score en fonction du signe joué
		score += roundscore[splitstring[1]]
		// ajout au score en fonction de l'issue de la game - utilisation de switch case
		switch splitstring[0]{
		case "A":
			switch splitstring[1]{
			case "X":
				score += Scissors
			case "Y":
				score += Rock
			case "Z":
				score += Paper
			}
		case "B":
			switch splitstring[1]{
			case "X":
				score += Rock
			case "Y":
				score += Paper
			case "Z":
				score += Scissors
		}
		case "C":
			switch splitstring[1]{
			case "X":
				score += Paper
			case "Y":
				score += Scissors
			case "Z":
				score += Rock
	}
	}
}
	return score
}

func readline(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("cannot open the file")
	}
	defer file.Close() // meilleure utilisation pour fermer les fichiers
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func main() {
	fmt.Println("Réponse étape 1 : ", solve1())
	fmt.Println("Réponse étape 2 : ", solve2())
}
