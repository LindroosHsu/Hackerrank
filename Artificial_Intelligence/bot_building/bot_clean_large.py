#https://www.hackerrank.com/challenges/botcleanlarge

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

def finding_dirty_pos(board, h, w):
    dirty_pos = list()
    
    for i in range(h):
        for j in range(w):
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
                    new_file.append(line)
        
        with open(file_name, 'w') as f:
            f.writelines(new_file)
                
                
    else:       
        if bot.get_x() > target.get_x():
            print("LEFT")

        elif bot.get_x() < target.get_x():
            print("RIGHT")

        elif bot.get_y() > target.get_y():
            print("UP")

        else:
            print("DOWN")

def backtrack(n, limit):
    global small_steps
    global target_pos
    global solution
    
    if n == limit:
        emu_bot_x = bot_x
        emu_bot_y = bot_y
        steps = 0
        
        for dirty in solution:
            steps += abs(emu_bot_x-dirty['x']) + abs(emu_bot_y-dirty['y'])
            emu_bot_x = dirty['x']
            emu_bot_y = dirty['y']
        
        if steps < small_steps:
            small_steps = steps
            target_pos['x'] = dirty_pos[0]['x']
            target_pos['y'] = dirty_pos[0]['y']
        
        return;
    
    for i in range(limit):
        if not used[i]:
            used[i] = True
            
            solution[n]['x'] = dirty_pos[n]['x']
            solution[n]['y'] = dirty_pos[n]['y']
            backtrack(n+1, limit)
            
            used[i] = False


if __name__ == "__main__":
    bot = Node(*map(int, input().split()))
    h, w = map(int, input().split())
    board = list(list(input()) for _ in range(h))
    dirty_pos_file = "dirty_positions"
    
    try:
        with open(dirty_pos_file) as f:
            target = Node(*map(int, f.readline().split()))
        
    except FileNotFoundError:
        dirty_pos = finding_dirty_pos(board, h, w)
        
        # calculate best step logic
        len_dirty = len(dirty_pos)
        used = [False] * len_dirty
        small_steps = 10000
        solution = [{'x':0, 'y':0}] * len_dirty
        
        backtrack(0, len_dirty)
        
        with open(dirty_pos_file, 'w') as f:
            for target in dirty_pos:
                f.write("{0} {1}\n".format(target.get_y(), target.get_x()))
                
        target = dirty_pos[0]
    
    next_move(bot, target, dirty_pos_file)
