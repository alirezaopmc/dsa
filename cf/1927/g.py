from collections import deque

for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    dp = [1 for i in range(n)]
    
    for i in range(1, n):
        x = a[i]
        
        # Use it
        
        
        # Don't use it