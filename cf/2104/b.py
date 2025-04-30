for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))


    p_max = a.copy()
    p_sum = a.copy()
    for i in range(1, n):
        p_max[i] = max(p_max[i-1], p_max[i])
        p_sum[i] += p_sum[i-1]

    # print(p_max)
    # print(p_sum)

    for k in range(1, n+1):
        m = p_max[n-k-1]
        s = p_sum[-1]
        # print(k, m, s)
        if k < n:
            s -= p_sum[n-1-k]
            if a[-k] < m:
                s += m - a[-k]
        print(s, end=' ')

    print()
