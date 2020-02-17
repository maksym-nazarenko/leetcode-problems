class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        trimmed_s = s.strip()

        return len(trimmed_s.split()[-1]) if len(trimmed_s) else 0
