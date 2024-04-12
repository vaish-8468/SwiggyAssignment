package main

import (
	"math/rand"
	"time"
)

type Player struct {
	 PlayerNumber int
	 Health int
	 Strength int
	 Attack int
}


/*****************************************************
RollDice simulates rolling a dice with the given sides
******************************************************/
func RollDiceHelperFunc(sides int) int {

	
	/**********************************************************************************************
	By seeding the random number generator with a changing value (such as the current time), 
	we ensure that the sequence of random numbers generated is different each time the program runs
	***********************************************************************************************/
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(sides) + 1
}

// AttackTurn calculates the damage done by an attacking player
func (attacker *Player) AttackTurn() int {
	return RollDiceHelperFunc(6) * attacker.Attack
}

// DefendTurn calculates the damage defended by a defending player
func (defender *Player) DefendTurn() int {
	return RollDiceHelperFunc(6) * defender.Strength
}