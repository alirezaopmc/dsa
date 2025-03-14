import math

n, m, k = map(int, input().split())

def fill_row(arr, i, cnt=None):
  n, m = len(arr), len(arr[0])
  if cnt is None: cnt = m
  for j in range(m):
    arr[i][j] = 'X'

def fill_col(arr, j, cnt=None):
  n, m = len(arr), len(arr[0])
  if cnt is None: cnt = math.ceil(n/2)
  for i in range(n):
    arr[i][j] = 'X'

if math.ceil(n/2) * math.ceil(m/2) < k:
  print(-1)
else:
  arr = [['O' for j in range(m)] for i in range(n)]
  # print(arr)

  # if n % 2 == 0:
  #   fill_row(arr, n-1)
  # if m % 2 == 0:
  #   fill_col(arr, m-1)
  
  if k <= math.ceil(n/2):
    for row in range(k-1):
      fill_row(arr, 2 * row + 1)
    k = 0
  else:
    for row in range(math.ceil(n/2)-1):
      fill_row(arr, 2*row+1)
    k -= math.ceil(n/2)
  print(k)
  col = 0
  while k > math.ceil(n/2):
    fill_col(arr, 2*col+1)
    col += 1
    k -= math.ceil(n/2)
  print(k, col)
  for i in range(2*k):
    arr[i][2*col+1] = 'X'
  for row in arr:
    print(''.join(row))
