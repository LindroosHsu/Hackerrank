#https://www.hackerrank.com/challenges/botcleanlarge

bot_y, bot_x = map(int, input().split())
h, w = map(int, input().split())
board = list(list(input()) for _ in range(h))

if board[bot_y][bot_x] == 'd':
    print("CLEAN")
        
else:
    dirty_pos = list()
        
    for i in range(h):
        for j in range(w):
            if board[i][j] == 'd':
                dirty_pos.append({'x':j, 'y':i})
                    
    small_pos = {'x':1000, 'y':1000}
    for pos in dirty_pos:
        if abs(pos['x']-bot_x)+abs(pos['y']-bot_y) < \
        abs(small_pos['x']-bot_x)+abs(small_pos['y']-bot_y):
            small_pos['x'] = pos['x']
            small_pos['y'] = pos['y']
                
    if bot_x > small_pos['x']:
        print("LEFT")

    elif bot_x < small_pos['x']:
        print("RIGHT")

    elif bot_y > small_pos['y']:
        print("UP")

    elif bot_y < small_pos['y']:
        print("DOWN")

