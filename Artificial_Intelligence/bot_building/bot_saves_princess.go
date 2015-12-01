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

func findBotAndPri(bot, pri *Position, botCh, priCh chan int) {
	var N int
	fmt.Scanf("%d", &N)

	for i := 0; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			var input int
			fmt.Scanf("%c", &input)

			if input == 'p' {
				priCh <- j
                priCh <- i
                close(priCh)

			} else if input == 'm' {
                botCh <- j
                botCh <- i
                close(botCh)
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
	bot := *NewPosition()
	pri := *NewPosition()
    botCh := make(chan int, 2)
    priCh := make(chan int, 2)

    // trying not to read all grid,
    // when found princess and bot location then jump into printNextStep
	go findBotAndPri(&bot, &pri, botCh, priCh)
    bot.SetXY(<-botCh, <-botCh)
    pri.SetXY(<-priCh, <-priCh)

	printNextStep(&bot, &pri)
}
