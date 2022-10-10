package main

import (
	"fmt"
	"math/rand"
	"time"
)

type void struct{}

const (
	numberOfPeopleInTheRoom = 23
	numberOfIterations      = 100000
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	successRate := runBdayParadoxScenarioIterations(numberOfIterations)
	fmt.Printf("%f%%", successRate*100)
}

func runBdayParadoxScenarioIterations(iterations int) float64 {
	succesfulIterations := 0
	for i := 0; i < iterations; i++ {
		foundBday := runBdayParadoxScenario()
		if foundBday {
			succesfulIterations += 1
		}
	}
	return float64(succesfulIterations) / float64(iterations)
}

func runBdayParadoxScenario() bool {
	var placeholder void
	bdays := make(map[int]void, numberOfPeopleInTheRoom)
	for i := 0; i < numberOfPeopleInTheRoom; i++ {
		generatedBday := rand.Intn(364) + 1
		bdays[generatedBday] = placeholder
	}
	return len(bdays) < numberOfPeopleInTheRoom // duplicate in the set
}
