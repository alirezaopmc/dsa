for _ in range(int(input())):
    l, r, d, u = map(int, input().split())
    
    if l == r == d == u:
        print("YES")
    else:
        print("NO")
