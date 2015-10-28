#https://www.hackerrank.com/challenges/utopian-tree

t = int(input())

for i in range(t):
    height = 1
    n = int(input())
    
    for j in range(n):
        if j >> 2 == 0: height *= 2
        else: height += 1
    
    print(height)
