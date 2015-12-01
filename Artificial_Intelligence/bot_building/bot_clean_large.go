package main

import (
    "fmt"
    "math"
)

type Position struct {
	x, y int
}

func NewPosition(setX, setY int) *Position {
	return &Position{x:setX, y:setY}
}

// scaning grid to find all dirties and check is the bot
// standing on a dirty
func findDirties(bot *Position, dirties []Position) bool {
	var h, w int
	fmt.Scanf("%d %d\n", &h, &w)

	for i := 0; i < h+1; i++ {
		for j := 0; j < w+1; j++ {
			var input int
			fmt.Scanf("%c", &input)

			if input == 'd' {
				if bot.x == j && bot.y == i {
                    // bot is standing on a dirty
                    return true
                }

                dirties = append(dirties, *NewPosition(j, i))
			}
		}
	}

    return false
}

func printNextStep(bot, dirty *Position) {
	if bot.x > dirty.x {
		fmt.Println("LEFT")

	} else if bot.x < dirty.x {
		fmt.Println("RIGHT")

	} else if bot.y > dirty.y {
		fmt.Println("UP")

	} else if bot.y < dirty.y {
		fmt.Println("DOWN")
	}
}

func findShortDirty()

func main() {
    var botX, botY int

    // in hackerrank X and Y are reversed...
    fmt.Scanf("%d %d\n", &botY, &botX)
	bot := *NewPosition(botX, botY)

    var dirties []Position
	onClean := findDirties(dirties)

    if onClean {
        fmt.Println("CLEAN")
        return
    }

    printNextStep(&bot, &pri)
}
