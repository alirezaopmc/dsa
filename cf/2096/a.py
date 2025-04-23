for _ in range(int(input())):
    n = int(input())
    s = input()

    low = 1
    high = n
    cur = n-1

    ans = [0] * n
    for x in s[::-1]:
        if x == '>':
            ans[cur] = high
            high -= 1
        else:
            ans[cur] = low
            low += 1
        cur -= 1
    ans[cur] = low

    print(*ans)
