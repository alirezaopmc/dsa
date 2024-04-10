from collections import defaultdict

for _ in range(int(input())):
    n = int(input())
    fr = {}
    fr = defaultdict(lambda: 0)
    a = [0 for i in range(n)]
    for i in range(n):
        a[i] = list(map(int, input().split()))
        for j in range(1, len(a[i])):
            fr[a[i][j]] += 1
    ans = 0
    # print(fr)
    m = len(fr.keys())
    for i in range(n):
        p = 0
        for x in a[i]:
            if fr[x] == 1:
                p += 1
        # print(i, p, fr[i])
        if m - p >= ans and p > 0:
            ans = m - p
    print(ans)
