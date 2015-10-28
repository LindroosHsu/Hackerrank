class Pair:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.diff = abs(x-y)

N = int(input())
arr = list(map(int, input().split()))
arr.sort()
small_diff = 10**7 + 10**7
pair_arr = list()

for i in range(N-1):
    pair = Pair(arr[i], arr[i+1])
    pair_arr.append(pair)
    if pair.diff < small_diff:
        small_diff = pair.diff
        
for pair in pair_arr:
    if pair.diff == small_diff:
        print(pair.x, pair.y, end=' ')
