#https://www.hackerrank.com/challenges/cut-the-sticks

input()
sticks = list(map(int, input().split()))

while sticks:
    cut = min(sticks)
    
    print(len(sticks))
    
    for i in range(len(sticks)-1, -1, -1):
        sticks[i] -= cut

        if sticks[i] <= 0:
            sticks.pop(i)

