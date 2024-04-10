for _ in range(int(input())):
    s = input()
    t = "abc"

    jump = False
    for i in range(len(s)):
        if s[i] == t[i]:
            print("YES")
            jump = True
            break
    if jump:
        continue
    print("NO")
