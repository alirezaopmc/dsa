for _ in range(int(input())):
    n = int(input())
    s0, e0 = map(int, input().split())
    ok = True
    for i in range(1, n):
        s, e = map(int, input().split())
        if not ok:
            continue
        if s >= s0 and e >= e0:
            ok = False
            continue
    if ok:
        print(s0)
    else:
        print(-1)
