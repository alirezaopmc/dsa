for _ in range(int(input())):
    n, k = map(int, input().split())

    is_even = lambda x: x % 2 == 0
    is_odd = lambda x: x % 2 == 1

    if is_even(n) and is_even(k):
        print((n+k-1)//k)
    elif is_even(n) and is_odd(k):
        k -= 1
        print((n+k-1)//k)
    elif is_odd(n) and is_even(k):
        n -= k-1
        print(1+(n+k-1)//k)
    else:
        n -= k
        k -= 1
        print(1+(n+k-1)//k)
