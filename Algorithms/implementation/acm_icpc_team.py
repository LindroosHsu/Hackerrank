#https://www.hackerrank.com/challenges/acm-icpc-team

N, M = map(int, input().split())
arr = list(input() for _ in range(N))

max_topic = 0
max_teams = 0

for i in range(N-1):
    person1 = int(arr[i], 2)
    
    for j in range(i+1, N):
        tmp_topic = 0
        person2 = int(arr[j], 2)
        
        topic = bin(person1 | person2).count('1')
        
        if topic > max_topic:
            max_topic = topic
            max_teams = 1
            
        elif topic == max_topic:
            max_teams += 1

print(max_topic)
print(max_teams)
