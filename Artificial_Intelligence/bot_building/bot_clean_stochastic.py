#https://www.hackerrank.com/challenges/botcleanr

bot_y, bot_x = map(int, input().split())
board = list(list(input()) for _ in range(5))

if board[bot_y][bot_x] == 'd':
    print("CLEAN")
else:
    dirty_pos = {'x':0, 'y':0}

    #searching dirty
    for i in range(5):
        try:
            j = board[i].index('d')
            dirty_pos['x'] = j
            dirty_pos['y'] = i
            break
        except ValueError:
            pass
    
    if bot_x > dirty_pos['x']:
        print("LEFT")

    elif bot_x < dirty_pos['x']:
        print("RIGHT")

    elif bot_y > dirty_pos['y']:
        print("UP")

    elif bot_y < dirty_pos['y']:
        print("DOWN")
