j = int(input())

for i in range(j):
    n, m, k = map(int, input().split(' '))
    a = sorted(map(int, input().split(' '))) + [-1]
    b = sorted(map(int, input().split(' '))) + [-1]
    t = 1
    ai = 0
    an = 0
    bi = 0
    bn = 0
    cond = True
    while t<k+1:
        while ai < n and a[ai] < t: ai+=1
        while bi < m and b[bi] < t: bi+=1
        if a[ai] != t and b[bi] != t: 
            cond=False 
            break
        if a[ai] == b[bi] == t:
            t+=1
        else:
            if a[ai] == t:
                t+=1
                an+=1
            if b[bi] == t:
                t+=1
                bn+=1
    print(an, bn)
    if cond==False or an > k/2 or bn > k/2: print("NO")
    else: print("YES")
