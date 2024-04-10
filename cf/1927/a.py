for _ in range(int(input())):
    n = int(input())
    s = input()
    a = s.find('B')
    b = n - s[::-1].find('B') - 1
    if a == -1:
        print(0)
    else:
        print(b-a+1)