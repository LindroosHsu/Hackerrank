counter_table = [[] for _ in range(100)]

N = int(input())
update_limit = N // 2

for i in range(N):
    index, s = input().split()
    index = int(index)
    counter_table[index].append('-' if i < update_limit else s)
    
for words in counter_table:
    for w in words:
        print(w, end=' ')
