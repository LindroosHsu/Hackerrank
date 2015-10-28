#https://www.hackerrank.com/challenges/service-lane

n, t = map(int, input().split())
width = list(map(int, input().split()))

for _ in range(t):
    i, j = map(int, input().split())
    if 1 in width[i:j+1]: print("1")
    elif 2 in width[i:j+1]: print("2")
    elif 3 in width[i:j+1]: print("3")
