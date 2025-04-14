for _ in range(int(input())):
    x, y, a = map(int, input().split())
    if a % (x+y) < x:
        print("NO")
    else:
        print("YES")
