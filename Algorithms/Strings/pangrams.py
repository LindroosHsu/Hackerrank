#https://www.hackerrank.com/challenges/pangrams

string = input()

if ('a' in string or 'A' in string) and\
('b' in string or 'B' in string) and\
('c' in string or 'C' in string) and\
('d' in string or 'D' in string) and\
('e' in string or 'E' in string) and\
('f' in string or 'F' in string) and\
('g' in string or 'G' in string) and\
('h' in string or 'H' in string) and\
('i' in string or 'I' in string) and\
('j' in string or 'J' in string) and\
('k' in string or 'K' in string) and\
('l' in string or 'L' in string) and\
('m' in string or 'M' in string) and\
('n' in string or 'N' in string) and\
('o' in string or 'O' in string) and\
('p' in string or 'P' in string) and\
('q' in string or 'Q' in string) and\
('r' in string or 'R' in string) and\
('s' in string or 'S' in string) and\
('t' in string or 'T' in string) and\
('u' in string or 'U' in string) and\
('v' in string or 'V' in string) and\
('w' in string or 'W' in string) and\
('x' in string or 'X' in string) and\
('y' in string or 'Y' in string) and\
('z' in string or 'Z' in string):
    print("pangram")
else:
    print("not pangram")

