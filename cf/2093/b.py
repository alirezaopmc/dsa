for _ in range(int(input())):
    s = input()
    i = len(s)-1
    while s[i] == '0':
        i -= 1
    zeros_after = len(s)-i-1
    zeros_before = s[:i].count('0')
    print(len(s)-1-zeros_before)
