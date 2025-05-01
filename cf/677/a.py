n, h = map(int, input().split())
a = list(map(int, input().split()))

c = 0
for x in a:
    if x > h:
        c += 1

print(n+c)
