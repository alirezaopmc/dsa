def solve(n, k, a, b):
    s = -1
    for i in range(n):
        if b[i] != -1:
            if s == -1:
                s = a[i] + b[i]
            elif s != a[i] + b[i]:
                return 0
    if s == -1:
        return k - max(a) + min(a) + 1
    for i in range(n):
        if a[i] > s or s - a[i] > k:
            return 0
    return 1
for _ in range(int(input())):
    n, k = map(int, input().split())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    print(solve(n, k, a, b))
