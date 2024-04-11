package main
import (
	"fmt"
)



func Fighting(player1 Player, player2 Player, num1 int, num2 int)  {
	fmt.Printf("Player%s attacks Player%s\n", num1, num2)
	attackDamage := player1.AttackTurn()
	defendDamage := player2.DefendTurn()

	fmt.Printf("Player%s deals %d damage\n", num1, attackDamage)
	fmt.Printf("Player%s defends %d damage\n", num2, defendDamage)

	if attackDamage > defendDamage {
		damage := attackDamage - defendDamage
		player2.Health -= damage
		fmt.Printf("Player%s takes %d damage\n", num1, damage)
	} else {
		fmt.Printf("Player%s defends all damage\n", num2)
	}

	fmt.Printf("Player%s's health: %d\n\n", num2, player2.Health)
}