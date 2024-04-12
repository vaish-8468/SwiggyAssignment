package main

import(
	"testing"
)


func TestRandomNumber(t *testing.T){
	var i int
	for i<100{
		got := RollDiceHelperFunc(6)
		if got>6 || got<1{
			t.Errorf("got %d, wanted between 1 to 6", got)
		}
		i++
	}
}

func TestDefenderHealth(t *testing.T){
	var got int
	var want int

	/************************
	Test Case 1
	(when attackDamage > defendDamage)
	**************************/
	got = FightResult(50,20,100)
	want =70

	if got!=want{
		t.Errorf("got %d, wanted %d", got, want)
	}

	/************************
	Test Case 2
	(when attackDamage<defendDamage)
	**************************/
	got = FightResult(50,60,40)
	want =40

    if got!=want{
		t.Errorf("got %d, wanted %d", got, want)
	}
}

