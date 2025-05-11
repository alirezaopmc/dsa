for _ in range(int(input())):
    n, m, p, q = map(int, input().split())

    if n % p == 0:
        s = (n // p) * q
        print("YES" if s == m else "NO")
    else:
        print("YES")
