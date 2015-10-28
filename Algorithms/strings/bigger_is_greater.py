# http://www.nayuki.io/page/next-lexicographical-permutation-algorithm

def next_permutation(arr):
    i = j = len(arr) - 1
    
    # Find non-increasing suffix
    while i > 0 and arr[i-1] >= arr[i]:
        i -= 1
    if i <= 0:
        return False
    
    # Find successor to pivot
    while arr[j] <= arr[i-1]:
        j -= 1
    
    # Swap with pivot
    arr[i-1], arr[j] = arr[j], arr[i-1]
    
    # Reverse suffix
    arr[i:] = arr[len(arr)-1:i-1:-1]
    
    return True

for _ in range(int(input())):
    s = list(input())
    
    if next_permutation(s):
        print("".join(s))
    else:
        print("no answer")

