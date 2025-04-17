def solve(k: int, n: int, ans: list, cur: set, lim: int):
    if n == 0 and len(cur) == k:
        ans.append(list(cur))

    for i in range(lim+1, 10):
        if i not in cur and i <= n:
            cur.add(i)
            solve(k, n-i, ans, cur, i)
            cur.remove(i)

class Solution:

    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        ans = []
        solve(k, n, ans, set(), 0)
        return ans



