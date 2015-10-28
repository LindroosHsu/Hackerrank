from collections import Counter

s = input()
words_cnter = Counter(s).values()
odd_cnter = [cnt for cnt in words_cnter if cnt % 2 != 0]

print("YES" if len(odd_cnter) <= 1 else "NO")
