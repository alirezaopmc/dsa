"""
gcd(x+d, y+d)=gcd(x-y, y+d)
"""

for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    a.sort()
    print(a[-1]-a[0])
