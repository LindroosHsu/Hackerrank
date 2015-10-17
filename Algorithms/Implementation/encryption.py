#https://www.hackerrank.com/challenges/encryption

from math import sqrt, ceil

words = input()
col = ceil(sqrt(len(words)))

for c in range(col):
    print(words[c::col], end=' ')
