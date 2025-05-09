n = int(input())
last = ''
ans = 0

for i in range(n):
    cur = input()
    if i == 0:
        ans += 1
        last = cur
    elif last != cur:
        ans += 1
        last = cur

print(ans)
