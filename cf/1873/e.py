for _ in range(int(input())):
    n, x = map(int, input().split())
    a = list(map(int, input().split()))

    def solve(h):
        s = 0
        for i in range(n):
            s += max(0, h - a[i])
        return s <= x

    l, r = 0, 10**10
    best = 0

    while l <= r:
        m = l + (r - l) // 2

        if solve(m):
            best = max(m, best)
            l = m + 1
        else:
            r = m - 1
    print(best)
