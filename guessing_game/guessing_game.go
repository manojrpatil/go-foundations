package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit = false
	var attempts = 0
	rand.Seed(time.Now().Unix())

	secretNumber = rand.Intn(99)
	//	fmt.Println("Secret number: ", secretNumber)

	for quit != true {
		attempts++
		fmt.Print("attempt# ", attempts)
		fmt.Printf(": Please enter any number (between 1 to 99): ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You Won in", attempts, "attempts!!")
			quit = true
		} else if userInput < secretNumber {
			fmt.Println("Too low... \nPlease try again!")
		} else if userInput > secretNumber {
			fmt.Println("Too high... \nPlease try again!")
		}
	}
}
