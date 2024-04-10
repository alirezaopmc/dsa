for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    k = 1
    for i in range(len(a)):
        k += 1
        if a[i] == k:
            k += 1
    print(k - 1)
