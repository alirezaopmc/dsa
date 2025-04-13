def is_prime(n):
    if n == 1:
        return False
    i = 2
    while i*i <= n:
        if n % i == 0:
            return False
        i += 1
    return True

for _ in range(int(input())):
    x, k = map(int, input().split())
    if x > 1 and k > 1:
        print("NO")
    elif x == 1:
        print("YES" if is_prime(int('1'*k)) else "NO")
    else:
        print("YES" if is_prime(x) else "NO")

