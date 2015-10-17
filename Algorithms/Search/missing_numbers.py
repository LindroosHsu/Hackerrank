#https://www.hackerrank.com/challenges/missing-numbers

from collections import Counter

input()
a = Counter(map(int, input().split()))

input()
b = Counter(map(int, input().split()))

for ele in b:
    if a[ele] != b[ele]:
        print(ele, end=' ')
