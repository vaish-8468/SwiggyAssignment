package main
import (
	"fmt"
	"strings"
	"strconv"
)



func Fighting(player1 *Player, player2 *Player)  {
	// ANSI escape codes to format text
	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"
	green := "\033[32m"
	yellow := "\033[33m"

	boxWidth := 76


	attackDamage := player1.AttackTurn()
	defendDamage := player2.DefendTurn()
	fmt.Print(""+"\n"+"\n"+"\n"+"")
	
	fmt.Printf("Player%d attacks 🤜 Player%d 🔥\n",player1.PlayerNumber,player2.PlayerNumber)
	fmt.Printf("%s\n",strings.Repeat("-", boxWidth))
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))
	fmt.Printf("|%15s%1s%1d%20s|%15s%1s%1d%20s|\n", "Player", "", player1.PlayerNumber,"","Player", "", player2.PlayerNumber,"")
	fmt.Printf("|%15s%1s%1d%1s%18s|%15s%1s%1d%1s%18s|\n", "(Health:", "", player1.Health,")","","(Health:", "", player2.Health,")","")
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))
	fmt.Printf("%s\n",strings.Repeat("-", boxWidth))

	// str1:=fmt.Sprintf("Player%d "+red+bold+"causes"+reset+" 💥 %d damage | Player%d "+green+bold+"defends"+reset+" 🛡️ with %d strength\n",player1.PlayerNumber,attackDamage,player2.PlayerNumber,defendDamage)
	str1:=fmt.Sprintf("Player%d "+red+bold+"causes"+reset+" %d damage",player1.PlayerNumber,attackDamage)

	str2:=fmt.Sprintf("Player%d "+green+bold+"defends"+reset+" 🛡️ with %d strength\n",player2.PlayerNumber,defendDamage)
	fmt.Printf("%s|\n",strings.Repeat(" ", 38))
	fmt.Printf("%s|\n",strings.Repeat(" ", 38))


	extraLen:= 40-len(str1)+10
	fmt.Printf("%s %s| %s",str1,strings.Repeat(" ",extraLen),str2)
	
	fmt.Printf("%s|\n",strings.Repeat(" ", 38))
	fmt.Printf("%s|\n",strings.Repeat(" ", 38))
	


	var damage int
	
	if attackDamage > defendDamage {
		damage = attackDamage - defendDamage
		if(len(strconv.Itoa(player2.Health-damage))>len(strconv.Itoa(player2.Health))){
			boxWidth += len(strconv.Itoa(player2.Health-damage))

		}
		//modify the health of player 2
		player2.Health -= damage

	/****************************************************
	BoxWidth is adjusted as per the length of the health
	*****************************************************/
	
	fmt.Printf("%s\n",strings.Repeat("-", boxWidth))
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))
	fmt.Printf("|%15s%1s%1d%20s|%15s%1s%1d%20s%s|\n", "Player", "", player1.PlayerNumber,"","Player", "", player2.PlayerNumber,"",strings.Repeat(" ",boxWidth-76))
	fmt.Printf("|%15s%1s%1d%1s%18s|"+red+bold+"%15s%1s%1d%1s%16s%s"+reset+"|\n", "(Health:", "", player1.Health,")","","(Health:", "", player2.Health,")","",strings.Repeat(" ",len(strconv.Itoa(player2.Health))))
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))

	} else {
	damage = 0
	fmt.Printf("%s\n",strings.Repeat("-", boxWidth))
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))
	fmt.Printf("|%15s%1s%1d%20s|%15s%1s%1d%20s|\n", "Player", "", player1.PlayerNumber,"","Player", "", player2.PlayerNumber,"")
	fmt.Printf("|%15s%1s%1d%1s%18s|"+green+bold+"%15s%1s%1d%1s%18s"+reset+"|\n", "(Health:", "", player1.Health,")","","(Health:", "", player2.Health,")","")
	fmt.Printf("|%s|\n", strings.Repeat(" ", boxWidth-1))
	}
	fmt.Printf("%s\n",strings.Repeat("-", boxWidth))
	fmt.Printf(yellow+bold+"Turn changed...🔁\n"+reset)

}