def solve(n, x, a):
    a.sort()

    if a[0] >= x:
        return n

    strongs = 0
    i = n-1
    cnt = 1

    while i >= 0:
        #print(f"i={i}, cnt={cnt}, a[i]={a[i]}")
        if cnt * a[i] >= x:
            strongs += 1
            cnt = 0
        i -= 1
        cnt += 1

    return strongs


for _ in range(int(input())):
    n, x = map(int, input().split())
    a = list(map(int, input().split()))

    print(solve(n, x, a))
