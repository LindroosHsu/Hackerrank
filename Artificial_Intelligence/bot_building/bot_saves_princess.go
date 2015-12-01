package main

import (
	"fmt"
)

const (
	Border = 101
)

type Position struct {
	x, y int
}

func (p *Position) SetXY(x, y int) {
	p.x = x
	p.y = y
}

func NewPosition() *Position {
	return &Position{x:Border, y:Border}
}

func findBotAndPri(bot, pri *Position) {
	var N int
	fmt.Scanf("%d", &N)

	for i := 0; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			var input int
			fmt.Scanf("%c", &input)

			if input == 'p' {
				pri.SetXY(j, i)

			} else if input == 'm' {
				bot.SetXY(j, i)
			}
		}
	}
}

func printNextStep(bot, pri *Position) {
	for bot.x != pri.x || bot.y != pri.y {
		if bot.x - pri.x > 0 {
			fmt.Println("LEFT")
			bot.x--

		} else if bot.x - pri.x < 0 {
			fmt.Println("RIGHT")
			bot.x++

		} else if bot.y - pri.y > 0 {
			fmt.Println("UP")
			bot.y--

		} else if bot.y - pri.y < 0 {
			fmt.Println("DOWN")
			bot.y++
		}
	}
}

func main() {
	var bot = *NewPosition()
	var pri = *NewPosition()

	findBotAndPri(&bot, &pri)

	printNextStep(&bot, &pri)
}
