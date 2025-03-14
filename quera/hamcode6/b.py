s = input()
t = input()
n = int(input())

st = set('abcdefghijklmnopqrstuvwxyz')

def find_missing_char(s: str):
  d = st.difference(set(s))
  if len(d) == 0:
    return None
  else:
    return list(d)[0]


if t in s:
  print("-1")
else:
  tplus = chr(ord('a')+(ord(t[0])-ord('a')+1)%26)
  print(tplus * (n-len(s)), s, sep='')