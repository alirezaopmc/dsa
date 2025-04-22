for _ in range(int(input())):
    n = int(input())
    s = input()
    s = '0' + s

    def compress(s):
        r = ''
        cur = s[0]
        i = 0
        while i < len(s):
            if s[i] != cur:
                r += cur
                cur = '1' if cur == '0' else '0'
            i += 1
        r += cur
        return r

    c = compress(s)
    
    def cost(c):
        if len(c) == 1:
            return 0
        elif len(c) <= 3:
            return 1
        else:
            return len(c) - 3

    all_cost = cost(c) + len(s) - 1
    print(all_cost)
