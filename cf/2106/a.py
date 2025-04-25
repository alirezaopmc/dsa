for _ in range(int(input())):
    n = int(input())
    s = input()
    a = s.count('0')
    b = n - a

    print(a + (n-1) * b)
