#https://www.hackerrank.com/challenges/sherlock-and-the-beast

limit = int(input())

for i in range(limit):
    n = int(input())
    m = n
    
    while m > 0:
        if m % 3 is 0: break
        else: m -= 5
    
    if ((n-m) % 5 is not 0) or (m < 0):
        print("-1", end='')
    else:
        for i in range(m):
            print("5", end='')
        for i in range(n-m):
            print("3", end='')
    print()
