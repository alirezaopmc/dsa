def solve(n, a):
    if n % 2 == 1:
        return True
    a.sort()
    i = 0
    while i < n:
        cnt = 0
        j = i
        while j < n and a[i] == a[j]:
            j += 1
            cnt += 1
        if cnt % 2 == 1:
            return True
        i = j
    return False

for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))

    print("YES" if solve(n, a) else "NO")
