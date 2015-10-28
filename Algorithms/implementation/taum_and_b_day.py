#https://www.hackerrank.com/challenges/taum-and-bday

T = int(input())

for _ in range(T):
    B, W = map(int, input().split())
    X, Y, Z = map(int, input().split())
    black_cost = X
    white_cost = Y
    
    if X > Y + Z:
        black_cost = Y + Z
        
    elif Y > X + Z:
        white_cost = X + Z
    
    print(black_cost*B + white_cost*W)
        
