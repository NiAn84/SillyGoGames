// A first try in Go by Niklas Andersson
// Teknisk Testare @ KYH Göteborg 2018
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var humanwin int
var computerwin int

func coinflip() {
	var choice int

	rand.Seed(time.Now().UnixNano())
	coinflip := 1 + rand.Intn(3)

	for {
		// fmt.Println(coinflip) // Remove comment for testing purposes only
		fmt.Printf("(1) Heads or (2) Tails?\n-->  ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("\nYou chose: Heads")
			break
		} else if choice == 2 {
			fmt.Println("\nYou chose: Tails")
			break
		} else {
			fmt.Println("Sorry, I didn't understand what you said to me.")
			fmt.Println("You have to choose 1 for Heads or 2 for Tails")
		}
	}

	if coinflip == 1 {
		fmt.Println("I flipped: Heads\n")
	} else if coinflip == 2 {
		fmt.Println("I flipped: Tails\n")
	} else {
		fmt.Println("You may have been flipped to infinity\n")
	}

	if coinflip == choice {
		humanwin++
		fmt.Println("You won!")
		score()
	} else if coinflip == 3 {
		fmt.Println("Oooops!\nSorry, but I lost the coin")
	} else {
		computerwin++
		fmt.Println("You lost.")
		score()
	}
}

func guessnumber() {
	rand.Seed(time.Now().UnixNano())
	thenumber := 1 + rand.Intn(10)
	var guess int
	var attempts int
	attempts = 3

	fmt.Println("\nI'm thinking of a number between 1 and 10.")
	fmt.Println("Can you guess which one? You get 3 attempts.")

	for {
		if attempts > 0 && guess != thenumber {
			// fmt.Println(thenumber) // Remove comment for testing purposes only
			fmt.Print("\nWhat is your guess? You have ", attempts, " attempt(s) left.\n-->")
			fmt.Scanln(&guess)
		}
		if guess != thenumber {
			attempts--
			fmt.Println("\nNope, the number is something else.")
		}
		if guess == thenumber {
			humanwin++
			fmt.Println("\nCongratulations you guessed the right number!")
			break
		}
		if attempts < 1 {
			computerwin++
			fmt.Println("\nYou guessed wrong 3 times, the number I was thinking on was", thenumber)
			break
		}
	}
	score()
}

func dicerules() {
	fmt.Println("\nYou get 1 point if the dice ends up at 6.")
	fmt.Println("But if you end up at 1, I get a point.")
}

func dice() {
	rand.Seed(time.Now().UnixNano())
	dice := 1 + rand.Intn(6)
	fmt.Println("\n---------")
	switch dice {
	case 1:
		computerwin++
		fmt.Println("|       |")
		fmt.Println("|   o   |")
		fmt.Println("|       |")
	case 2:
		fmt.Println("| o     |")
		fmt.Println("|       |")
		fmt.Println("|     o |")
	case 3:
		fmt.Println("| o     |")
		fmt.Println("|   o   |")
		fmt.Println("|     o |")
	case 4:
		fmt.Println("| o   o |")
		fmt.Println("|       |")
		fmt.Println("| o   o |")
	case 5:
		fmt.Println("| o   o |")
		fmt.Println("|   o   |")
		fmt.Println("| o   o |")
	case 6:
		humanwin++
		for i := 1; i <= 3; i++ {
			fmt.Println("| o   o |")
		}
	}
	fmt.Println("---------")
}

func score() {
	fmt.Println("\n### Score ###")
	fmt.Println("Human:   ", humanwin)
	fmt.Println("Computer:", computerwin)
}

func main() {
	for {
		var choice int8
		fmt.Println("\n##### Main Menu #####\n")
		fmt.Println("1: Flip coin")
		fmt.Println("2: Guess the number")
		fmt.Println("3: Toss 1 dice")
		fmt.Println("4: Toss 2 dices")
		fmt.Println("5: Show current score")
		fmt.Println("6: Exit")
		fmt.Print("\nChoice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			coinflip()
		case 2:
			guessnumber()
		case 3:
			dicerules()
			dice()
			score()
		case 4:
			dicerules()
			dice()
			dice()
			score()
		case 5:
			score()
		case 6:
			os.Exit(0)
		}
	}
}
