T = int(input())

for _ in range(T):
    N = int(input())
    chocols = list(map(int, input().split()))
    minChocol = min(chocols)

    answer = 1000000001
    for i in range(6):
        counter = 0
        for chocol in chocols:
            v = chocol - (minChocol - i)
            counter += int(v/5) + int((v%5)/2) + int((v%5)%2)

        if counter < answer:
            answer = counter

    print(answer)
