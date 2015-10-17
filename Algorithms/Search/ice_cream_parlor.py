#https://www.hackerrank.com/challenges/icecream-parlor

T = int(input())

for _ in range(T):
    M = int(input())
    N = int(input())
    C = list(map(int, input().split()))
    
    for i in range(N-1):
        for j in range(i+1, N):
            if C[i] + C[j] == M:
                print(i+1, j+1)
                break
