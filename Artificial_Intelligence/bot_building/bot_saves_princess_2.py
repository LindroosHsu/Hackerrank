#https://www.hackerrank.com/challenges/saveprincess2

#!/usr/bin/python

m = int(input())
pri_pos = {'x':0, 'y':0}
bot_pos = {'x':0, 'y':0}

bot_pos['y'], bot_pos['x'] = map(int, input().split())
grid = list(list(input().strip()) for _ in range(m))

#finding princess position
for i in range(m):
    try:
        j = grid[i].index('p')
        pri_pos['x'] = j
        pri_pos['y'] = i
        break
        
    except ValueError:
        pass

if bot_pos['x'] > pri_pos['x']:
    print("LEFT")
         
elif bot_pos['x'] < pri_pos['x']:
    print("RIGHT")
        
elif bot_pos['y'] > pri_pos['y']:
    print("UP")
        
elif bot_pos['y'] < pri_pos['y']:
    print("DOWN")
