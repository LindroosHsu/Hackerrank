# http://www.geeksforgeeks.org/anagram-substring-search-search-permutations/

from collections import defaultdict

for _ in range(int(input())):
    s = input()
    len_s = len(s)
    d = defaultdict(int)
    
    for i in range(len_s):
        for j in range(1, len_s-i+1):
            sub_s = s[i:i+j]
            sub_s = sorted(sub_s)
            d["".join(sub_s)] += 1
            
    answer = 0
    for v in d.values():
        answer += v*(v-1) // 2
    
    print(answer)
