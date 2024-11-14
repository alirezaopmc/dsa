for _ in range(int(input())):
  n, k = map(int, input().split())
  a = list(map(int, input().split()))
  a.sort()
  
  def check(p):
    s = 0
    colors = [ set() for i in range(k) ]
    for x in a:
      si = s
      while x not in colors[si]: si += 1
      
  
  left, right = 0, n // k
  while left < right:
    mid = left + (right - left) // 2
    if check()