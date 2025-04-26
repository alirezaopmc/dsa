for _ in range(int(input())):
    n, x = map(int, input().split())
    i = 0
    x_skip = False
    while i < n:
        if i != x:
            print(i, end=' ')
        else:
            x_skip = True
        i += 1
    if x_skip:
        print(x)
    else:
        print()
