#https://www.hackerrank.com/challenges/funny-string

t = int(input())

for _ in range(t):
    s = input()
    r = s[::-1]
    fu = True
    
    for i in range(1, len(s)):
        if abs(ord(s[i])-ord(s[i-1])) != abs(ord(r[i])-ord(r[i-1])):
            fu = False
            break

    if fu: print("Funny")
    else: print("Not Funny")
    
