#https://www.hackerrank.com/challenges/insertionsort1

s = int(input())
ar = list(map(int, input().split()))
v = 0
pos = 0

for i in range(1, s):
    #finding V
    if ar[i] < ar[i-1]:
        v = ar[i]
        pos = i
        break
    
while ar[pos] < ar[pos-1] and pos > 0:
    ar[pos], ar[pos-1] = ar[pos-1], ar[pos]
    
    print_ar = ar.copy()
    print_ar[pos-1] = print_ar[pos]
    print(*print_ar)
    pos -= 1
print(*ar)
    
