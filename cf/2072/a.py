import math

for _ in range(int(input())):
    n, k, p = map(int, input().split())
    ans = math.ceil(abs(k) / p)
    if ans <= n:
        print(ans)
    else:
        print("-1")
