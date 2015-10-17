#https://www.hackerrank.com/challenges/the-time-in-words

mins = {
    1:"one minute",
    2:"two minutes",
    3:"three minutes",
    4:"four minutes",
    5:"five minutes",
    6:"six minutes",
    7:"seven minutes",
    8:"eight minutes",
    9:"nine minutes",
    10:"ten minutes",
    11:"eleven minutes",
    12:"twelve minutes",
    13:"thirteen minutes",
    14:"fourteen minutes",
    15:"quarter",
    16:"sixteen minutes",
    17:"seventeen minutes",
    18:"eighteen minutes",
    19:"nineteen minutes",
    20:"twenty minutes",
    21:"twenty one minutes",
    22:"twenty two minutes",
    23:"twenty three minutes",
    24:"twenty four minutes",
    25:"twenty five minutes",
    26:"twenty six minutes",
    27:"twenty seven minutes",
    28:"twenty eight minutes",
    29:"twenty nine minutes",
    30:"half"
}

hours = {
    1:"one",
    2:"two",
    3:"three",
    4:"four",
    5:"five",
    6:"six",
    7:"seven",
    8:"eight",
    9:"nine",
    10:"ten",
    11:"eleven",
    12:"twelve",
    13:"one"
}

H = int(input())
M = int(input())

if M == 0:
    print(hours[H] + " o' clock")
elif M <= 30:
    print(mins[M] + " past " + hours[H])
else:
    print(mins[60-M] + ' to ' + hours[H+1])
