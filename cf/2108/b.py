for _ in range(int(input())):
    n, x = map(int, input().split())
    if n == 1 and x == 0:
        print(-1)
        continue

    b = bin(x)[2:].count('1') # The number of '1's in binary

    if n <= b:
        print(x)
    elif (n-b) % 2 == 0:
        print(x + (n-b))
    elif x == 0:
        if n == 1:
            print(-1)
        else:
            print(n+3)
    elif x == 1:
        print(n+3)
    else:
        print(x+n-b+1)

