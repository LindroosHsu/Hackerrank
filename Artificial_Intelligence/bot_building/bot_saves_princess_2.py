#https://www.hackerrank.com/challenges/saveprincess2

#!/usr/bin/python

class Node:
    def __init__(self, y, x):
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

def finding_pri_pos(grid, grid_size):
    pri_pos = Node(0, 0)
    
    for i in range(grid_size):
        try:
            j = grid[i].index('p')
            pri_pos.set_x(j)
            pri_pos.set_y(i)
            break
            
        except ValueError:
            pass
            
    return pri_pos

def next_step(bot_pos, pri_pos):
    if bot_pos.get_x() > pri_pos.get_x():
        print("LEFT")

    elif bot_pos.get_x() < pri_pos.get_x():
        print("RIGHT")

    elif bot_pos.get_y() > pri_pos.get_y():
        print("UP")

    else:
        print("DOWN")


if __name__ == "__main__":
    m = int(input())
    bot_pos = Node(*map(int, input().split()))
    grid = list(list(input().strip()) for _ in range(m))
    pri_pos_file = "princess_position"
    
    try:
        with open(pri_pos_file) as f:
            pri_pos = Node(*map(int, f.readline().split()))
        
    except FileNotFoundError:
        pri_pos = finding_pri_pos(grid, m)
        
        with open(pri_pos_file, 'w') as f:
            f.write(str(pri_pos.get_y()) + " " + str(pri_pos.get_x()))
        
    next_step(bot_pos, pri_pos)
