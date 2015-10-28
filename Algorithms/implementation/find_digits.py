#https://www.hackerrank.com/challenges/find-digits

t = int(input())

for i in range(t):
    n = input()
    l = list(filter(lambda n: n is not 0, [int(j) for j in n]))
    n = int(n)
    k = list()
    
    for j in l:
        if n % j == 0: k.append(j)
        
    print(len(k))
    
    
