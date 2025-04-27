for _ in range(int(input())):
    s = input()
    a = [0 for i in range(10)]
    result = ''
    for x in s:
        a[int(x)] += 1

    for j in range(10):
        low = 10-j-1
        high = 9
        for k in range(low, high+1):
            if a[k] > 0:
                a[k] -= 1
                result += str(k)
                break

    print(result)

