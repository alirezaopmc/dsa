for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))

    s = 0
    for i in range(len(a)):
        s += a[i] if i % 2 == 0 else -a[i]

    print(s)
