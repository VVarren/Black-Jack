package main

import (
	"fmt"
	"math/rand"
	"time"
)

//TO-DO LISt
// UI
// Maybe make the computer more intelligent and entirely up to the player

func main() {
	var input string
	arr := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	notWin := true
	playerHand := initialHand()
	dealerHand := initialHand()
	rand.Seed(time.Now().UnixNano())

	for notWin {
		fmt.Print(playerHand)
		fmt.Println("Player Hand")
		fmt.Println("Type 'Draw' to draw a card and 'Quit' to quit")
		fmt.Println("If you wnat to change your ace(11) to a 1 or vice versa type 'Change' ")

		fmt.Scan(&input)
		fmt.Println("Computer Won")
		if input == "Quit" {

			if winTeller(playerHand, dealerHand, 2) == "Player has won!" {
				fmt.Println("Player has won!")
				notWin = false
			}
			if winTeller(playerHand, dealerHand, 2) == "Computer has won!" {
				fmt.Println("Computer has won!")
				notWin = false
			}

		} else if input == "Change" {
			playerHand = oneToEleven(playerHand)
			fmt.Print("It has been changed")
			fmt.Println(playerHand)

		} else if input == "Draw" {
			playerHand = append(playerHand, arr[rand.Intn(len(arr)-1)])
			dealerHand = append(dealerHand, arr[rand.Intn(len(arr)-1)])
			fmt.Print(playerHand)
			fmt.Println(dealerHand)
			if winTeller(playerHand, dealerHand, 1) == "Computer Lost" {
				fmt.Println("Player has won")
				notWin = false
			} else if winTeller(playerHand, dealerHand, 1) == "Player Lost" {
				fmt.Println("Computer Won")
				notWin = false
			}
		}

	}
}

func oneToEleven(Hand []int) []int {
	for i := 0; i < len(Hand); i++ {
		if Hand[i] == 11 {
			Hand[i] = 1
		} else if Hand[i] == 1 {
			Hand[i] = 11
		}
	}
	return Hand
}
func initialHand() []int {
	arr := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	list := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	r1, r2 := arr[rand.Intn(len(arr)-1)], arr[rand.Intn(len(arr)-1)]
	list = append(list, r1)
	list = append(list, r2)
	fmt.Println(list)
	return list
}

func winTeller(playerHand []int, dealerHand []int, cases int) string {

	switch cases {
	case 1:
		if sum(dealerHand) > 21 {
			return "Computer Lost"
		} else if sum(playerHand) > 21 {
			return "Player Lost"
		}

	case 2:
		if sum(playerHand) > sum(dealerHand) {
			return "Player has won!"
		} else if sum(dealerHand) > sum(playerHand) {
			return "Computer has won!"
		}
	}
	return "This doesnt work"
}

func sum(arr []int) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total
}
