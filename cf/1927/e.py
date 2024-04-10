from collections import deque

for _ in range(int(input())):
    n, k = map(int, input().split())
    a = [0 for i in range(n)]
    
    top = n
    low = 1
    
    def get_cur(i):
        global low
        global top
        if i % 2 == 0:
            low += 1
            return low - 1
        else:
            top -= 1
            return top + 1
    
    for i in range(k):
        j = i
        while j < n:
            a[j] = get_cur(i)
            j += k
    
    print(*a)