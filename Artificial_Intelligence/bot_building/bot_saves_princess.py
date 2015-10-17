#https://www.hackerrank.com/challenges/saveprincess

#!/usr/bin/python

class Node:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def set_x(self, x):
        self.x = x
    
    def set_y(self, y):
        self.y = y
    
    def get_x(self):
        return self.x
        
    def get_y(self):
        return self.y
    
    def distance(self, other):
        return abs(self.x - other.x) + abs(self.y - other.y)
     
def finding_pri_bot_pos(grid, grid_size):
    bot_pos = Node(0, 0)
    pri_pos = Node(0, 0)
    found_pri = False
    found_bot = False
    
    #finding princess and bot position
    for i in range(grid_size):
        for j in range(grid_size):
            if grid[i][j] == 'p':
                pri_pos.set_x(j)
                pri_pos.set_y(i)
                found_pri = True
                
            elif grid[i][j] == 'm':
                bot_pos.set_x(j)
                bot_pos.set_y(i)
                found_bot = True
                
            if found_bot and found_pri:
                break
        if found_bot and found_pri:
            break
            
    return pri_pos, bot_pos
    
def next_step(bot_x, bot_y, pri_x, pri_y):
    if bot_x > pri_x:
        for _ in range(bot_x - pri_x):
            print("LEFT")

    else:
        for _ in range(pri_x - bot_x):
            print("RIGHT")

    if bot_y > pri_y:
        for _ in range(bot_y - pri_y):
            print("UP")

    else:
        for _ in range(pri_y - bot_y):
            print("DOWN")
            

if __name__ == "__main__":
    m = int(input())
    grid = list(list(input().strip()) for _ in range(m))
    pri_loc_file_name = "princess_location"

    pri_pos, bot_pos = finding_pri_bot_pos(grid, m)
    next_step(bot_pos.get_x(), bot_pos.get_y(), pri_pos.get_x(), pri_pos.get_y())

