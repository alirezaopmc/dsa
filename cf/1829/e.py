for _ in range(int(input())):
  # Inputs
  n, m = map(int, input().split())
  a = []
  for i in range(n):
    a.append(list(map(int, input().split())))

  # helper functions
  def index2num(i, j):
    return m*i+j
  def num2index(k):
    return (k//m,k%m)

  # Data structures
  s = [ [0 for j in range(m)] for i in range(n) ]
  p = [ [index2num(i, j) for j in range(m) ] for i in range(n) ]
  
  # DFS
  for i in range(n):
    for j in range(m):
      if i > 0:
        if a[i-1][j] > 0:
          p[i][j] = p[i-1][j]
      if j > 0:
        if a[i][j-1] > 0:
          p[i][j] = p[i][j-1]
      ip, jp = num2index(p[i][j])
      s[ip][jp] += a[i][j]
  
  print(max(max(s, key=lambda row: max(row))))