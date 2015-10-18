#https://www.hackerrank.com/challenges/botcleanr

#!/usr/bin/python

#
# borad size in this problem is static 5*5
#

class Node:
    def __init__(self, y=0, x=0):
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
    dirty = Node()
    
    # searching dirty, because it is randomly appear one dirty node,
    # so it can not be using scanning all of the dirties method.
    for i in range(5):
        try:
            j = board[i].index('d')
            dirty.set_x(j)
            dirty.set_y(i)
            break
            
        except ValueError:
            pass
            
    return dirty

def next_move(bot_x, bot_y, target_x, target_y):
    
    # the bot is under the dirty position
    if bot_x == target_x and bot_y == target_y:
        print("CLEAN")
                
    else:       
        if bot_x > target_x:
            print("LEFT")

        elif bot_x < target_x:
            print("RIGHT")

        elif bot_y > target_y:
            print("UP")

        else:
            print("DOWN")
                    
                    
if __name__ == "__main__":
    bot = Node(*map(int, input().split()))
    board = list(list(input()) for _ in range(5))
    
    target = finding_dirty_pos(board)
    next_move(bot.get_x(), bot.get_y(), target.get_x(), target.get_y())
            
