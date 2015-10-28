#https://www.hackerrank.com/challenges/insertionsort2

s = int(input())
ar = list(map(int, input().split()))

for i in range(1, s):

    cur_value = ar[i]
    pos = i

    while pos > 0 and ar[pos-1] > cur_value:
        ar[pos] = ar[pos-1]
        pos = pos-1

    ar[pos] = cur_value
    
    print(*ar)
