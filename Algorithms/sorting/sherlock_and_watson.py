N, K, Q = map(int, input().split())
A = list(map(int, input().split()))
K  %= N

for _ in range(Q):
    idx = int(input())
    idx -= K
    if idx < 0:
        idx += N
        
    print(A[idx])
