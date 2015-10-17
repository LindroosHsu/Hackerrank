#https://www.hackerrank.com/challenges/sherlock-and-array

T = int(input())

for _ in range(T):
    N = int(input())
    A = list(map(int, input().split()))

    left_index = 0
    right_index = N-1
    left_sum = A[left_index]
    right_sum = A[right_index]

    while left_index != right_index:
        if left_sum < right_sum:
            left_index += 1
            left_sum += A[left_index]
        else:
            right_index -= 1
            right_sum += A[right_index]
            
    if left_sum == right_sum:  
        print("YES")
    else:  
        print("NO")
