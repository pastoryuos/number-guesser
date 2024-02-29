package main

import (
	"fmt"
	"math/rand"
)

type IPlayer interface {
	guess() int
}

type Human struct {
}

func (*Human) guess() int {

	var max, min int = 10, 0
	x := rand.Intn(max-min) + min
	//fmt.Println(x) for debugging
	Nguess := 0
	for Nguess < 3 {
		fmt.Println("Enter your next guess:")
		var i int
		fmt.Scanln(&i)

		if i == x {
			fmt.Println("You Win")
			return 0
		} else if i > x {
			fmt.Println("Too High")
			Nguess++
		} else if i < x {
			fmt.Println("Too Low")
			Nguess++
		}
	}
	fmt.Println("You ran out of guesses. Game over")
	return 1
}

type Autoguess struct {
}

func (*Autoguess) guess() int {
	var max, min int = 10, 0
	x := rand.Intn(max-min) + min
	Nguess := 0
	i := (max-min)/2 + min
	for Nguess < 3 {
		
		fmt.Println("The computer has chosen ", i)

		if i == x {
			fmt.Println("You Win")
			
			return 0
		} else if i > x {
			fmt.Println("Too High")
			i = (i-min)/2 + min
			Nguess++
		} else if i < x {
			fmt.Println("Too Low")
			i += (max-i)/2
			Nguess++
		}
	}
	fmt.Println("You ran out of guesses. Game over")
	return 1
}

type Game struct {
	player IPlayer
}

func (g Game) play() {
	//fmt.Println(g.player.guess()) i dont see the need of the return value being print, might as well makes the 2 funct void funct but oh well
	g.player.guess()
}

func main() {
	var gm Game

	fmt.Println("You have 3 guesses to guess a number from 1 to 10/n")

	//the scan funct dont work as intended
	fmt.Println("Do you want to make the guesses? (y/n -- if n guesses will be generated for you):")
	var choice string
	fmt.Scanln(&choice)

	if choice == "n" {
		gm.player = new(Autoguess)
	} else {
		gm.player = new(Human)
	}
	gm.play()
}