#https://www.hackerrank.com/challenges/kaprekar-numbers

p = int(input())
q = int(input())
answer = list()

for number in range(p, q+1):
    power_str = str(number ** 2)
    power_len = len(power_str)
    
    if power_len & 1 != 0:
        power_str = power_str.zfill(power_len+1)
        power_len = len(power_str)
    
    limit = power_len // 2
    
    if int(power_str[:limit]) + int(power_str[limit:]) == number:
        answer.append(number)
    
if answer:    
    print(*answer)
else:
    print("INVALID RANGE")
