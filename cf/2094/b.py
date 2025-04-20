for _ in range(int(input())):
    n, m, l, r = map(int, input().split())
    la = r-m
    ra = r
    if la > 0:
        ra -= la
        la = 0
    print(la, ra)
