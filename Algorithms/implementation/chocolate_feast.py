#https://www.hackerrank.com/challenges/chocolate-feast

T = int(input())

for i in range(T):
    # N $ in his pocket
    # C The price of each chocolate
    # M wrappers he gives to the store
    N,C,M = map(int, input().split())
    answer = 0
    
    # Compute Answer
    answer = N // C
    wrapper_in_hand = answer
    
    while wrapper_in_hand >= M:
        cal = wrapper_in_hand // M
        answer += cal
        wrapper_in_hand -= cal * M
        wrapper_in_hand += cal
    
    print(answer)

