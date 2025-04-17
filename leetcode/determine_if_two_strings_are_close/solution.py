def occs(s: str):
    """
    "aaabbbbbbccccc" ➜ [3, 6, 5]
    """
    if not s:
        return []

    a = list(s)
    a.sort()

    result = []
    count = 1

    for i in range(1, len(s)):
        if a[i] == a[i - 1]:
            count += 1
        else:
            result.append(count)
            count = 1

    result.append(count)  # Add the last group count
    result.sort()
    return result

class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        # Condition 1 - Length
        if len(word1) != len(word2):
            return False
        
        # Condition 2 - Char Set
        if set(word1) != set(word2):
            return False

        # Condition 3 - Occurrence Count
        if occs(word1) != occs(word2):
            return False

        return True
