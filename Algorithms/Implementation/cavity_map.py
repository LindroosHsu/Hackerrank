#https://www.hackerrank.com/challenges/cavity-map

import copy

n = int(input())
mm = list(list(map(int, input())) for i in range(n))
result = copy.deepcopy(mm)

for i in range(1,n-1):
    for j in range(1,n-1):
        if mm[i-1][j] < mm[i][j] and \
        mm[i+1][j] < mm[i][j] and \
        mm[i][j-1] < mm[i][j] and \
        mm[i][j+1] < mm[i][j]:
            result[i][j] = 'X'

for i in result:
    print(*i, sep="")
