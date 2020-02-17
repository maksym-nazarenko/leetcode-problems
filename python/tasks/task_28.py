class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if needle == "":
            return 0

        if len(haystack) < len(needle):
            return -1

        for start_index in range(0, len(haystack) - len(needle) + 1):
            for index in range(0, len(needle)):
                if haystack[start_index + index] != needle[index]:
                    break
                if index == len(needle) - 1:
                    return start_index

        return -1
