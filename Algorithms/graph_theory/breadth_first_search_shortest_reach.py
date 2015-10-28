UNITS_LEN = 6
INF_NUM = 1000*UNITS_LEN

for _ in range(int(input())):
    N, M = map(int, input().split())
    edges = list()
    
    for _ in range(M):
        # node index start from 1, but array start from 0
        s, d = map(int, input().split())
        edges.append((s-1, d-1))
        edges.append((d-1, s-1))
    
    S = int(input())
    # node index start from 1, but array start from 0
    S -= 1
    dist = [INF_NUM] * N
    dist[S] = 0
    
    for _ in range(N):
        refresh = False
        
        for edge in edges:
            s, d = edge
            
            if dist[s]+UNITS_LEN < dist[d]:
                dist[d] = dist[s] + UNITS_LEN
                refresh = True
                
        if not refresh:
            break
    
    del dist[S]
    print(*[-1 if e == INF_NUM else e for e in dist])
    
