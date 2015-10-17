#https://www.hackerrank.com/challenges/caesar-cipher-1

def rotate(s_num, k):
    if s_num >= ord('A') and s_num <= ord('Z'):
        s_num += k
        if s_num > ord('Z'): s_num -= 26
            
    elif s_num >= ord('a') and s_num <= ord('z'):
        s_num += k
        if s_num > ord('z'): s_num -= 26

    return s_num

n = int(input())
string = list(map(ord, input()))
k = int(input())
k %= (ord('z') - ord('a') + 1)

for i in range(len(string)):
    print(chr(rotate(string[i], k)), end='')
