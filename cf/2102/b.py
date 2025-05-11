for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))

    for i in range(n):
        a[i] = a[i] if a[i] > 0 else -a[i]
    #print(a)
    m = a[0]

    a.sort()
    #print(a[:n//2])

    if m in a[:n//2+1]:
        print("YES")
    else:
        print("NO")
