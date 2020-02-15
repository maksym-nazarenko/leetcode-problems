class Solution:
    """
    Given two strings A and B of lowercase letters,
    return true if and only if we can swap two letters in A so that the result equals B.

    """
    def buddyStrings(self, A: str, B: str) -> bool:
        if len(A) != len(B):
            return False

        if A == B and len(set(A)) != len(A):
            return True

        index_1 = index_2 = None

        for index, chars in enumerate(zip(A, B)):
            char_a, char_b = chars
            if char_a != char_b:
                if index_1 is None:
                    index_1 = index
                elif index_2 is None:
                    index_2 = index
                else:
                    return False

        if index_1 is None or index_2 is None:
            return False

        return (A[index_1], A[index_2]) == (B[index_2], B[index_1])


sol = Solution()

tests = {
    ("ab", "ba"): True,
    ("ab", "ab"): False,
    ("aa", "aa"): True,
    ("aaaaaaabc", "aaaaaaacb"): True,
    ("", "aa"): False,
}

for k, v in tests.items():
    print(f"<{k[0]}, {k[1]}> [{v}]: {sol.buddyStrings(*k)}")
