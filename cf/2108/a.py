for _ in range(int(input())):
    n = int(input())

    m = n // 2
    k = m ** 2

    if n % 2 == 1:
        print(n * (n-1) // 2 - k + 1)
    else:
        print(2 * k // 2 + 1)
