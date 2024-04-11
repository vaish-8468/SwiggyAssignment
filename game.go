package main

import "fmt"



func Arena(players []Player) {
	/*********************************
	ANSI escape codes to format text
	*********************************/
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"
	green := "\033[32m"
	// yellow := "\033[33m"

	fmt.Println(red+bold+"Choose any two players to play the game"+reset)

	/***************************************************************************
	nums1 and nums2 are the converted into the 0-based indexing of the players
	***************************************************************************/
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	if num1>len(players) || num2>len(players) {
		fmt.Println("Enter valid numbers")
		return 
	}
	fmt.Println(green+bold+"Great! You can start the game now"+reset)
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

	 /**********************************************
	 Fight until one player's health reaches 0
	 *********************************************/
	for player1.Health > 0 && player2.Health > 0 {
		Fighting(attacker, defender, num1, num2)
		
		// Swap attacker and defender after each turn
		attacker, defender = defender, attacker
	}

	// Declare the winner
	// var winner Player
	// if player1.Health <= 0 {
	// 	winner = player2
	// } else {
	// 	winner = player1
	// }
	// fmt.Printf("%s wins the game!\n", num1)
}






	

