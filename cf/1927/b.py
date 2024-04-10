from collections import defaultdict

for _ in range(int(input())):
    d = defaultdict(lambda: set())
    for i in range(26):
        d[0].add(chr(ord('a')+i))
    n = int(input())
    a = list(map(int, input().split()))
    for x in a:
        c = d[x].pop()
        d[x+1].add(c)
        print(c, end='')
    print()