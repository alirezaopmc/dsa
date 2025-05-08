def solve():
    x, n, m = map(int, input().split())

    floor = lambda x: x // 2
    ceil = lambda x: (x+1) // 2

    def min_x(x, n, m):
        while n > 0 and m > 0 and x > 1:
            if x & 2 == 1:
                x = floor(x)
                n -= 1
            else:
                x = ceil(x)
                m -= 1

        while n > 0 and x > 0:
            x = floor(x)
            n -= 1

        while m > 0 and x > 1:
            x = ceil(x)
            m -= 1

        return x

    def max_x(x, n, m):
        while n > 0 and m > 0 and x > 1:
            if x & 2 == 1:
                x = ceil(x)
                m -= 1
            else:
                x = floor(x)
                n -= 1

        while n > 0 and x > 0:
            x = floor(x)
            n -= 1

        while m > 0 and x > 1:
            x = ceil(x)
            m -= 1

        return x

    minx = min_x(x, n, m)
    maxx = max_x(x, n, m)

    print(minx, maxx)

for _ in range(int(input())):
    solve()

