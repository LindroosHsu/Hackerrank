#https://www.hackerrank.com/challenges/the-grid-search

T = int(input())

for _ in range(T):
    R, C = map(int, input().split())
    G = list(list(map(int, input())) for i in range(R))
    r, c = map(int, input().split())
    P = list(list(map(int, input())) for i in range(r))
    found = False
    
    for i in range(R):
        for j in range(C):
            if G[i][j] == P[0][0]:
                found = True
                for k in range(r):
                    if G[i+k][j:j+c] != P[k][:]:
                        found = False
                        break
            if found: break
        if found: break
    
    if found: print("YES")
    else: print("NO")
    
