package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"

)
type Player struct {
	Health int
	Strength int
	Attack int
}



func main() {
	// ANSI escape codes to format text
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"
	green := "\033[32m"
	yellow := "\033[33m"


	// // Created a slice to store player objects
	// var players []Player



	/**************************************************************
	not all terminal emulators support all ANSI escape codes, 
	so the appearance of text may vary depending on the terminal being used.
	***************************************************************/
	fmt.Println("**********************************************")
	fmt.Println(red + bold + "Welcome to this Magical Arena" + reset)
	fmt.Println("**********************************************")
	 fmt.Println(red+bold+"$       $       $ $$$$$$"+reset)
	 fmt.Println(red+bold+" $     $ $     $  $"+reset)
	 fmt.Println(red+bold+"  $   $   $   $   $$$$"+reset)
	 fmt.Println(red+bold+"   $ $     $ $    $"+reset)
	 fmt.Println(red+bold+"    $       $     $$$$$$"+reset)





	/***************************************************************
	This code block has been coded such that it can become scalable 
	if we want to include more number of players than only 2
	****************************************************************/
	fmt.Println(yellow+"Enter the number of players in the game: "+reset)
    var num int
	fmt.Scanln(&num)
	if num<2 {
		fmt.Println("Too less players")
		fmt.Println("Minimum number of players is 2")
	}
	t := 1
	// Created a slice to store player objects
	var players []Player
	for num>0{
		fmt.Printf(green+"Enter Player %d Details in the format: <Health> <Strength> <Attack>\n"+reset,t)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		data := strings.Split(input, " ")

		health, _ := strconv.Atoi(data[0])
		strength, _ := strconv.Atoi(data[1])
		attack, _ := strconv.Atoi(data[2])
		p := new(Player)
		p.Health = health
		p.Attack = attack
		p.Strength = strength
		// Created a new Player object and appended it to the players slice
		players = append(players, *p)
		num--
		t++
	}

	// Calling the Arena function with the players slice
	// err := Arena(players)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}