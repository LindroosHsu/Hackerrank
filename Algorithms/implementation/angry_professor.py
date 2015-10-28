#https://www.hackerrank.com/challenges/angry-professor

t = int(input())
for i in range(t):
    n, k = map(int, input().split())
    arrived_time = list(map(int, input().split()))
	
    if n - len(list(filter(lambda n: n > 0, arrived_time))) < k:
        print("YES")
    else:
        print("NO")
