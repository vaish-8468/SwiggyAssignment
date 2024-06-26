package main

import "fmt"



func Arena(players []Player) {
	/*********************************
	ANSI escape codes to format text
	*********************************/

	bold := "\033[1m"
	reset := "\033[0m"
	blue := "\033[34m"
	yellow := "\033[33m"

	fmt.Println(yellow+bold+"Choose any two players to play the game in the format: <player number> <player number>"+reset)

	/***************************************************************************
	nums1 and nums2 are the converted into the 0-based indexing of the players
	***************************************************************************/
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	if num1>len(players) || num2>len(players) {
		fmt.Println("Enter valid numbers")
		return 
	}
	fmt.Println(blue+bold+"Great! Game started...🎮"+reset)
	player1 :=players[num1-1]
	player2 :=players[num2-1]

	/*********************************************
	Determined who attacks first based on health
	**********************************************/
	var attacker, defender Player
	
	if player1.Health <= player2.Health {
		attacker = player1
		defender = player2
	} else {
		attacker = player2
		defender = player1
	}
	
	// Declare the winner
	// var winner Player

	/**********************************************
	Fight until one player's health reaches 0
	*********************************************/
	for attacker.Health > 0  {
		Fighting(&attacker, &defender)	
		// Swap attacker and defender after each turn
		attacker, defender = defender, attacker
	}

	fmt.Printf(blue+bold+"Player%d wins 🤴🏼 the game!\n"+reset, defender.PlayerNumber)
	
}







	

