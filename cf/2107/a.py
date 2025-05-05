def solve():
    n = int(input())
    a = list(map(int, input().split()))

    p = min(a)
    q = max(a)
    if p == q:
        print("NO")
        return
    else:
        print("YES")
        for x in a:
            print(2 if x == q else 1, end=' ')
        print()

for _ in range(int(input())):
    solve()
