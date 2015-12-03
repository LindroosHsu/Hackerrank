package main

import (
    "fmt"
    "math"
)

type Position struct {
	x, y int
}

type DirtyNode struct {
    Position
    dis int
}

func NewPosition(setX, setY int) *Position {
	return &Position{setX, setY}
}

func NewDirtyNode(setX, setY, distance int) *DirtyNode {
    return &DirtyNode{Position{setX, setY}, distance}
}

// 1. scaning grid to find all dirties and check is the bot standing on a dirty
// 2. calculate the shortest disctance then return the nearby dirty
func findDirties(bot *Position) *DirtyNode {
	var h, w int
	fmt.Scanf("%d %d\n", &h, &w)

    // the dirty should be cleaned means the shortest distance dirty
    theDirty := *NewDirtyNode(0, 0, h*w)
	for dirtyY := 0; dirtyY < h; dirtyY++ {
		for dirtyX := 0; dirtyX < w+1; dirtyX++ {
			var input int
			fmt.Scanf("%c", &input)

			if input == 'd' {
				if bot.x == dirtyX && bot.y == dirtyY {
                    // bot is standing on a dirty
                    return nil
                }

                // calculate the distance from bot to this dirty
                distance := int(math.Abs(float64(dirtyX-bot.x)) + math.Abs(float64(dirtyY-bot.y)))
                if distance < theDirty.dis {
                    theDirty.dis = distance
                    theDirty.x = dirtyX
                    theDirty.y = dirtyY
                }
			}
		}
	}

    return &theDirty
}

func printNextStep(bot *Position, dirty *DirtyNode) {
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

func main() {
    var botX, botY int

    // in hackerrank X and Y are reversed...
    fmt.Scanf("%d %d\n", &botY, &botX)

	bot := NewPosition(botX, botY)
	dirty := findDirties(bot)
    if dirty == nil {
        fmt.Println("CLEAN")
        return
    }

    printNextStep(bot, dirty)
}
