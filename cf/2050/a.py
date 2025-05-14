for _ in range(int(input())):
    n, m = map(int, input().split())
    s = [input() for i in range(n)]

    i = 0
    while i < n and m - len(s[i]) >= 0:
        m -= len(s[i])
        i += 1

    print(i)
