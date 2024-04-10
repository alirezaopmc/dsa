for _ in range(int(input())):
    n, k = map(int, input().split())
    s = input()

    i = 0
    cnt = 0
    while i < n:
        if s[i] == "B":
            cnt += 1
            i += k
        else:
            i += 1
    print(cnt)
