package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"

)

/*******************************************************
This program can take multiple players 
and it allows to choose any two players to play the game
********************************************************/


func main() {
	// ANSI escape codes to format text
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"


	// Created a slice to store player objects
	var players []Player



	/**************************************************************
	not all terminal emulators support all ANSI escape codes, 
	so the appearance of text may vary depending on the terminal being used.
	***************************************************************/

	fmt.Println(blue+bold+"*************************************************"+reset+"\n" +
	         red + bold + "💥💥💥💥💥Welcome to this Magical Arena💥💥💥💥💥" + reset + "\n" +
	         blue+bold+"*************************************************"+reset)




	/***************************************************************
	This code block has been coded such that it can become scalable 
	if we want to include more number of players than only 2
	****************************************************************/
	fmt.Println(yellow+"Enter the number of players 🦹🏼 in the game: "+reset)
    var numberOfPlayers int
	fmt.Scanln(&numberOfPlayers)


	if numberOfPlayers<2 {
		//loops till the number of players become greater than or equal to 2
		for numberOfPlayers<2 {
		fmt.Println(red+bold+"Too less players")
		fmt.Println("Minimum number of players is 2"+reset)
		fmt.Println(yellow+"Enter the number of players in the game: "+reset)
		fmt.Scanln(&numberOfPlayers)
		}

	}


	t := 1
	//player numeration starts from 1



	for numberOfPlayers>0{
		fmt.Printf(green+"Enter Player %d Details in the format: <Health> <Strength> <Attack>\n"+reset,t)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		data := strings.Split(input, " ")

		playerNum := t
		health, _ := strconv.Atoi(data[0])
		strength, _ := strconv.Atoi(data[1])
		attack, _ := strconv.Atoi(data[2])

		// Created a new Player object and appended it to the players slice
		p := new(Player)
		p.PlayerNumber = playerNum
		p.Health = health
		p.Attack = attack
		p.Strength = strength
		players = append(players, *p)
		numberOfPlayers--
		t++
	}

	// Calling the Arena function with the players slice
	Arena(players)

}