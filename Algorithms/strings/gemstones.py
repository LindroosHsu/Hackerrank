from collections import Counter

N = int(input())
words = list(set(input()) for _ in range(N))

# by using Counter needs to transform subarries into a main arry...
cnt_list = list()
for e in words:
    for w in e:
        cnt_list.append(w)

# create a set of word which is appeared in this problem
words = set(cnt_list)
c = Counter(cnt_list)
answer = 0

for e in words:
    if c[e] == N:
        answer += 1

print(answer)
