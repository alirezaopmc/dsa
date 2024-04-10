from collections import deque

for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    left = [-1 for i in range(n)]
    # right = [0 for i in range(n)]
    
    # Left
    last_two = deque(maxlen=2)
    for i in range(n):
        x = a[i]
        if i == 0:
            last_two.append((x, i))
            continue
        if last_two[-1][0] != x:
            last_two.append((x, i))
            left[i] = last_two[0][1]
        else:
            if last_two[0][0] != x:
                left[i] = last_two[0][1]
            last_two[-1] = (x, i)

    """
    # Right
    last_two = deque(maxlen=2)
    b = a.copy()[::-1]
    for i in range(n):
        x = b[i]
        if i == 0:
            last_two.append((x, i))
            continue
        if last_two[-1][0] != x:
            last_two.append((x, i))
            right[i] = last_two[0][1]
        else:
            right[i] = last_two[-1][1]
            last_two[-1] = (x, i)
    right = right[::-1]
    """
     
    # print(left)
    # print(right)
    
    
    q = int(input())
    for __ in range(q):
        l, r = map(int, input().split())
        l -= 1
        r -= 1
        ans = left[r]
        if ans >= l:
            print(ans+1, r+1)
        else:
            print(-1, -1)