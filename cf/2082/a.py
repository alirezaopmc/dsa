def solve():
    n, m = map(int, input().split())
    a = [input() for _ in range(n)]

    p, q = 0, 0
    for i in range(n):
        xor = 0
        for j in range(m):
            xor ^= int(a[i][j])
        p += xor # if xor is 1 then add to p

    for j in range(m):
        xor = 0
        for i in range(n):
            xor ^= int(a[i][j])
        q += xor # if xor is 1 then add ot q

    if p == 0 or q == 0:
        print(max(p, q))
    else:
        print(min(p, q) + abs(p-q))


for _ in range(int(input())):
    solve()
