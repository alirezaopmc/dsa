n=int(input())

for i in range(n):
  print(i+1, " + " if i < n-1 else " = ", end='')

print(n * (n+1) // 2)