package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("codewars-highest-scoring-word  please run unit tests")
}

func HighestScoringWord(words string) string {
	highScore := 0
	highScoreWord := ""

	wordArr := strings.Split(words, " ")
	for _, word := range wordArr {
		thisWordScore := 0
		
		for _, c := range strings.ToLower(word) {
			thisWordScore += int(c-'a') + 1
		}
		
		fmt.Printf("word: %v  score: %v\n", word, thisWordScore)
		if thisWordScore >= highScore {
			highScore = thisWordScore
			highScoreWord = word
		}

	}

	return highScoreWord
}