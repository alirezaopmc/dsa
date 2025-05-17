def solve():
    n = int(input())
    a = list(map(int, input().split()))

    if a.count(1) == n:
        print("YES")
        return

    for i in range(n-1):
        if a[i] == a[i+1] == 0:
            print("YES")
            return

    print("NO")
    return

for _ in range(int(input())):
    solve()
