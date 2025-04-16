def cnt_fiboness(a):
    return len(list(filter(lambda i: a[i] + a[i+1] == a[i+2], range(3))))

for _ in range(int(input())):
    a1, a2, a4, a5 = list(map(int, input().split()))
    ans = 1
    a3 = a1 + a2
    ans = max(ans, cnt_fiboness([a1, a2, a3, a4, a5]))
    a3 = a4 - a2
    ans = max(ans, cnt_fiboness([a1, a2, a3, a4, a5]))
    a3 = a5 - a4
    ans = max(ans, cnt_fiboness([a1, a2, a3, a4, a5]))
    print(ans)

