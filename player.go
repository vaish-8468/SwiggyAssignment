package main

import (
	"math/rand"
	"time"
)

type Player struct {
	 Health int
	 Strength int
	 Attack int
}


// RollDice simulates rolling a dice with the given sides
func RollDice(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

// Attack calculates the damage done by an attacking player
func (attacker *Player) AttackTurn() int {
	return RollDice(6) * attacker.Attack
}

// Defend calculates the damage defended by a defending player
func (defender *Player) DefendTurn() int {
	return RollDice(6) * defender.Strength
}