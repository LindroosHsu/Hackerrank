from collections import deque

def is_substring(A, B):
    found = False
    
    while A:
        if A.pop() in B:
            found = True
            
    return found

for _ in range(int(input())):
    A = deque(set(input()))
    B = deque(set(input()))
    
    print("YES" if is_substring(A, B) else "NO")
