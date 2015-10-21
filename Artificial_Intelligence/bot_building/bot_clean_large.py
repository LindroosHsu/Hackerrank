#https://www.hackerrank.com/challenges/botcleanlarge

from itertools import permutations

class Node:
    def __init__(self, y, x):
        self.x = x
        self.y = y
    
    def distance(self, other):
        return abs(self.x - other.x) + abs(self.y - other.y)

def finding_dirty_pos(board, h, w):
    for i in range(h):
        for j in range(w):
            if board[i][j] == 'd':
                yield(Node(i, j))

def calculate_total_distance(bot, dirties, shorted_path):
    emu_bot = Node(bot.y, bot.x)
    total_distance = 0
        
    for dirty in dirties:
        if total_distance >= shorted_path:
            break
            
        total_distance += emu_bot.distance(dirty)
        emu_bot.x = dirty.x
        emu_bot.y = dirty.y
            
    return total_distance
                
def next_move(bot, target, file_name):
    
    # the bot is under the dirty position
    if bot.x == target.x and bot.y == target.y:
        print("CLEAN")
        
        # since we cleaned the dirty, remove it from file
        with open(file_name) as f:
            new_file = list()
            
            for line in f:
                if not "{0} {1}".format(target.y, target.x) in line:
                    new_file.append(line)
        
        # because I want to overwrite the original file,
        # so read file first then write file after.
        with open(file_name, 'w') as f:
            f.writelines(new_file)
                
    else:       
        if bot.x > target.x:
            print("LEFT")

        elif bot.x < target.x:
            print("RIGHT")

        elif bot.y > target.y:
            print("UP")

        else:
            print("DOWN")
    

if __name__ == "__main__":
    bot = Node(*map(int, input().split()))
    h, w = map(int, input().split())
    board = list(list(input()) for _ in range(h))
    dirty_pos_file = "dirty_positions"
    
    try:
        with open(dirty_pos_file) as f:
            target = Node(*map(int, f.readline().split()))
        
    except FileNotFoundError:
        dirty_poses = finding_dirty_pos(board, h, w)
        
        # calculate best step logic
        shorted_path = h * w
        target_dirties = None
        
        for dirties in permutations(dirty_poses):
            tmp = calculate_total_distance(bot, dirties, shorted_path)
            
            if tmp < shorted_path:
                shorted_path = tmp
                target_dirties = dirties
        
        with open(dirty_pos_file, 'w') as f:
            for target in target_dirties:
                f.write("{0} {1}\n".format(target.y, target.x))
                
        target = target_dirties[0]
    
    next_move(bot, target, dirty_pos_file)
