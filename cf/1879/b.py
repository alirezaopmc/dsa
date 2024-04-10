for _ in range(int(input())):
    n = int(input())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))
    A = sum(a)
    B = sum(b)
    print(
        min(
            n * min(b) + A,
            n * min(a) + B,
        )
    )
