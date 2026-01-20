package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	mod := float64(score - 10) / 2
	finalValue := math.Floor(mod)
	return int(finalValue)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	count := 0
	diceScore := make([]int, 4)

	for i := range 4 {
		diceScore[i] = rand.Intn(6) + 1
	}

	sort.Ints(diceScore)
	bestThree := diceScore[1:]
	for _, value := range bestThree {
		count += value
	}
	return count
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var c Character

	c.Charisma = Ability()
	c.Constitution = Ability()
	c.Dexterity = Ability()
	c.Hitpoints = Ability()
	c.Intelligence = Ability()
	c.Strength = Ability()
	c.Wisdom = Ability()

	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
