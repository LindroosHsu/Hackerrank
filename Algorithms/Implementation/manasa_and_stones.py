#https://www.hackerrank.com/challenges/manasa-and-stones

def last_stone(a, b, n):
    if a == b:
        return [a * (n - 1)]
    if a < b:
        return last_stone(b, a, n)
    
    return [i * a + (n - i - 1) * b for i in range(n)]

T = int(input())

for _ in range(T):
    n = int(input())
    a = int(input())
    b = int(input())
    
    answer = last_stone(a, b, n)
    print(*answer)
