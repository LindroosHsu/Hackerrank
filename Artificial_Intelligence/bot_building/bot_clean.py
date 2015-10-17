#https://www.hackerrank.com/challenges/botclean

#!/usr/bin/python

# Head ends here
def next_move(bot_y, bot_x, board):
    
    # the bot is under the dirty position
    if board[bot_y][bot_x] == 'd':
        print("CLEAN")
        
    else:
        dirty_pos = list()
        
        for i in range(5):
            for j in range(5):
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
                    
                    
# Tail starts here
if __name__ == "__main__":
    pos = list(map(int, input().split()))
    board = list(list(input()) for _ in range(5))
    next_move(pos[0], pos[1], board)

