package main
import (
	"fmt"
)



func Fighting(player1 *Player, player2 *Player)  {
	// ANSI escape codes to format text
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"
	green := "\033[32m"
	// yellow := "\033[33m"


	fmt.Printf("Player%d attacks Player%d\n", player1.PlayerNumber, player2.PlayerNumber)
	attackDamage := player1.AttackTurn()
	defendDamage := player2.DefendTurn()
	fmt.Printf("Player%d health is: %d\n", player1.PlayerNumber, player1.Health)
	fmt.Printf("Player%d health is: %d\n", player2.PlayerNumber, player2.Health)
	fmt.Printf("Player%d causes %d damage\n", player1.PlayerNumber, attackDamage)
	fmt.Printf("Player%d defends %d damage\n", player2.PlayerNumber, defendDamage)

	var damage int
	if attackDamage > defendDamage {
		damage = attackDamage - defendDamage
		fmt.Printf(red+bold+"Player%d takes %d damage\n"+reset, player1.PlayerNumber, damage)
	} else {
		damage = 0
		fmt.Printf(green+bold+"Player%d defends all damage\n"+reset, player2.PlayerNumber)
	}
	player2.Health -= damage

	fmt.Printf("Player%d's health: %d\n\n", player2.PlayerNumber, player2.Health)
}