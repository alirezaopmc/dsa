for _ in range(int(input())):
  m, a, b, c = map(int, input().split())
  a_ok = min(m, a)
  b_ok = min(m, b)
  rem = 2 * m - a_ok - b_ok
  c_ok = min(rem, c)
  print(a_ok + b_ok + c_ok)