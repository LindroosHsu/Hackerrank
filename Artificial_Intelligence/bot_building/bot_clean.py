#https://www.hackerrank.com/challenges/botclean

#!/usr/bin/python

#
# borad size in this problem is static 5*5
#

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

def finding_dirty_pos(board):
    dirty_pos = list()
    
    for i in range(5):
        for j in range(5):
            if board[i][j] == 'd':
                dirty_pos.append(Node(i, j))
            
    return dirty_pos

def next_move(bot, target, file_name):
    
    # the bot is under the dirty position
    if bot.get_x() == target.get_x() and bot.get_y() == target.get_y():
        print("CLEAN")
        
        # since we cleaned the dirty, remove it from file
        with open(file_name) as f:
            new_file = list()
            
            for line in f:
                if not "{0} {1}".format(target.get_y(), target.get_x()) in line:
                    y, x = map(int, line.split())
                    new_file.append(Node(y, x))
        
        new_file = sorted(new_file, key=bot.distance)
        
        with open(file_name, 'w') as f:
            for target in new_file:
                f.write("{0} {1}\n".format(target.get_y(), target.get_x()))
                
    else:       
        if bot.get_x() > target.get_x():
            print("LEFT")

        elif bot.get_x() < target.get_x():
            print("RIGHT")

        elif bot.get_y() > target.get_y():
            print("UP")

        else:
            print("DOWN")
                    
                    
if __name__ == "__main__":
    bot = Node(*map(int, input().split()))
    board = list(list(input()) for _ in range(5))
    dirty_pos_file = "dirty_positions"
    
    try:
        with open(dirty_pos_file) as f:
            target = Node(*map(int, f.readline().split()))
        
    except FileNotFoundError:
        dirty_pos = sorted(finding_dirty_pos(board), key=bot.distance)
        
        with open(dirty_pos_file, 'w') as f:
            for target in dirty_pos:
                f.write("{0} {1}\n".format(target.get_y(), target.get_x()))
                
        target = dirty_pos[0]
    
    next_move(bot, target, dirty_pos_file)

