for _ in range(int(input())):
  s = input()
  s = s[::-1]
  s = s.replace('p', 'a')
  s = s.replace('q', 'p')
  s = s.replace('a', 'q')
  print(s)