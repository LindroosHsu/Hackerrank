A, B, N = map(int, input().split())
C = B**2 + A

for i in range(3, N):
    A = B
    B = C
    C = B**2 + A

print(C)
