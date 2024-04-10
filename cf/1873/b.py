for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    a.sort()
    s = 1
    for i in range(1, n):
        s *= a[i]
    print(s * (a[0] + 1))
