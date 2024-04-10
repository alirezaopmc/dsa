n = 10


def get_array():
    return [input() for _ in range(n)]


def convert(x, y):
    if x >= 5:
        x = n - x - 1
    if y >= 5:
        y = n - y - 1
    return x, y


def tier(x, y):
    x, y = convert(x, y)
    return min(x, y) + 1


for _ in range(int(input())):
    arr = get_array()
    s = 0
    for x in range(n):
        for y in range(n):
            if arr[x][y] == "X":
                # print(x, y, convert(x, y), tier(x, y))
                s += tier(x, y)
    print(s)
