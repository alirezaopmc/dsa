for i in range(int(input())):
    n, m, k = map(int, input().split())
    a = set(map(int, input().split()))
    b = set(map(int, input().split()))
    
    def is_valid(x):
        return 1 <= x <= k
    
    def filter_k(s):
        return set(filter(is_valid, s))
    
    both = filter_k(a.intersection(b))
    just_a = filter_k(a.difference(b))
    just_b = filter_k(b.difference(a))
    
    if len(just_a) + len(both) >= k // 2 and len(just_b) + len(both) >= k // 2:
        tmp = both.copy()
        tmp.update(just_a)
        tmp.update(just_b)
        if len(tmp) == k:
            print("YES")
        else:
            print("NO")
    else:
        print("NO")
    