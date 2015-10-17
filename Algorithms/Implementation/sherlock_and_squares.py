#https://www.hackerrank.com/challenges/sherlock-and-squares

from math import sqrt, floor

t = int(input())

for i in range(t):
    a, b = map(int, input().split())
    print(floor(sqrt(b)) - floor(sqrt(a-1)))
