a = [0] * (2 * 10**5 + 100)
k = 0

MOD = 998244353

dp = [-1] * (2 * 10**5 + 100)
dp[0] = dp[1] = 1


def rev(n):
    return pow(n, MOD - 2, MOD)


def fact(n):
    if dp[n] == -1:
        dp[n] = fact(n - 1) * n % MOD
    return dp[n]


for _ in range(int(input())):
    s = input()
    ans = 1
    k = 0
    last = 0
    for i in range(1, len(s)):
        if s[i] != s[last]:
            a[k] = i - last
            k += 1
            last = i
    a[k] = len(s) - last
    k += 1
    print(a[:k])
    p = 0
    for i in range(k):
        ans = (ans * a[i]) % MOD
        # ans = (ans * fact(a[i])) % MOD
        # if a[i] > 1:
        #     ans = (ans * rev(a[i] - 1)) % MOD
    ans = (ans * fact(len(s) - k)) % MOD
    print(len(s) - k, ans)
