from collections import Counter

for _ in range(int(input())):
    N = int(input())
    A = Counter(map(int, input().split()))
    answer = 0
    
    for value in A.values():
        if value != 1:
            answer += value * (value-1)
            
    print(answer)

