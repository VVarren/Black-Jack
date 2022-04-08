package main1

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var arr [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

//TO-DO LISt
// UI


func main() {
	var input string
	notWin := true
	PlayerStand := false
	DealerStand := false
	playerHand := initialHand()
	dealerHand := initialHand()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Type 'Draw' to draw a card and 'STAND' to quit \nIf you wnat to change your ace(11) to a 1 or vice versa type 'Change' \nIf you want to know the commands, please  type 'Help'.")

	for notWin {
		fmt.Println(playerHand, " Player Hand", "\n", dealerHand, " Computer Hand")
		fmt.Print("What will you do next: ")
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		if input == "STAND" {
			PlayerStand = true
			dealerHand = smartBot(dealerHand)
			for !PlayerStand || !DealerStand {
				if sum(dealerHand)> 19 {
					DealerStand = true
			} else {
				num := rand.Intn(1)
				if num == 1 {
					DealerStand = true
				} else{
					dealerHand = append(dealerHand, arr[rand.Intn(len(arr))])
				}
			}

			if PlayerStand || DealerStand {
				if winTeller(playerHand, dealerHand, 2) == "Player has won!" {
					fmt.Println("Player has won!")
					notWin = false
				}
				if winTeller(playerHand, dealerHand, 2) == "Computer has won!" {
					fmt.Println("Computer has won!")
					notWin = false
				}
			}

		} else if input == "CHANGE" {
			playerHand = oneToEleven(playerHand)
			fmt.Println(playerHand, " Your hand has been changed")

		} else if input == "DRAW" {
			playerHand = append(playerHand, arr[rand.Intn(len(arr))])
			dealerHand = append(dealerHand, arr[rand.Intn(len(arr))])
			if winTeller(playerHand, dealerHand, 3) == "Computer Lost" {
				fmt.Println("Player has won")
				notWin = false
			} else if winTeller(playerHand, dealerHand, 3) == "Player Lost" {
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
	list := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	r1, r2 := arr[rand.Intn(len(arr))], arr[rand.Intn(len(arr))]
	list = append(list, r1)
	list = append(list, r2)
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
	case 3:
		if sum(dealerHand) > 21 {
			return "Computer Lost"
		} else if sum(playerHand) > 21 {
			return "Player Lost"
		} else if sum(playerHand) > sum(dealerHand) {
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

func smartBot(dealerHand []int) []int {
	//choices are draw when at 11, 16 or stand
	if sum(dealerHand) < 11 {
		dealerHand = append(dealerHand, arr[rand.Intn(len(arr))])
		return dealerHand
	} else if sum(dealerHand) < rand.Intn(3)+16 {
		num := rand.Intn(2)
		if num == 0 {
			return dealerHand
		} else if num == 1 {
			dealerHand = append(dealerHand, arr[rand.Intn(len(arr))])
			return dealerHand

		}
	}
	return dealerHand
}
