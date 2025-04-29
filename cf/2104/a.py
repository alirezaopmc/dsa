def solve(a, b, c):
    s = a + b + c

    if s % 3 != 0:
        return False

    x = s // 3

    return a <= x and b <= x and c >= x

for _ in range(int(input())):
    a, b, c = map(int, input().split())

    print("YES" if solve(a, b, c) else "NO")
