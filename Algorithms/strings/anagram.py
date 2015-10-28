from collections import Counter

def anagram(s):
    s_len = len(s)
    
    if s_len % 2 == 1:
        return -1
    
    mid = s_len // 2
    c1 = Counter(s[:mid])
    c2 = Counter(s[mid:])
    answer = 0
    
    for e in c1:
        if e in c2:
            answer += max(0, c1[e]-c2[e])
        else:
            answer += c1[e]
            
    return answer

for _ in range(int(input())):
    s = input()
    print(anagram(s))
    
