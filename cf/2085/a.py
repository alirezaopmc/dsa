def solve():
    n, k = map(int, input().split())
    s = input()

    st = set(s)

    if len(s) == 1 or (k == 0 and s >= s[::-1]) or len(st) == 1:
        print("NO")
    else:
        print("YES")

for _ in range(int(input())):
    solve()
