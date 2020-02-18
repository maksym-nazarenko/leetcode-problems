from collections import deque


class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) == 0:
            return True

        opening_brackets = ('(', '{', '[')
        closing_brackets = (')', '}', ']')
        stack = deque()

        for c in s:
            if c in opening_brackets:
                stack.append(c)
                continue

            if c in closing_brackets:
                if not len(stack):
                    return False

                if stack.pop() != opening_brackets[closing_brackets.index(c)]:
                    return False

        return len(stack) == 0
