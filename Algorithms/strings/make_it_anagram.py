from collections import Counter

standard = "abcdefghijklmnopqrstuvwxyz"
A = Counter(input())
B = Counter(input())
answer = 0

for w in standard:
    answer += abs(A[w] - B[w])

print(answer)
