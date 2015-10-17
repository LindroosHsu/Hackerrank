#https://www.hackerrank.com/challenges/saveprincess

#!/usr/bin/python

m = int(input())
grid = list(list(input().strip()) for _ in range(m))
bot_pos = {'x':0, 'y':0}
pri_pos = {'x':0, 'y':0}

#finding princess and bot position
for i in range(m):
    for j in range(m):
        if grid[i][j] == 'p':
            pri_pos['x'] = j
            pri_pos['y'] = i
            
        elif grid[i][j] == 'm':
            bot_pos['x'] = j
            bot_pos['y'] = i
            
        if bot_pos['x'] != 0 and pri_pos['x'] != 0:
            break
    if bot_pos['x'] != 0 and pri_pos['x'] != 0:
        break

if bot_pos['x'] > pri_pos['x']:
    for step in range(bot_pos['x']-pri_pos['x']):
        print("LEFT")
        
elif bot_pos['x'] < pri_pos['x']:
    for step in range(pri_pos['x']-bot_pos['x']):
        print("RIGHT")
        
if bot_pos['y'] > pri_pos['y']:
    for step in range(bot_pos['y']-pri_pos['y']):
        print("UP")
        
elif bot_pos['y'] < pri_pos['y']:
    for step in range(pri_pos['y']-bot_pos['y']):
        print("DOWN")

