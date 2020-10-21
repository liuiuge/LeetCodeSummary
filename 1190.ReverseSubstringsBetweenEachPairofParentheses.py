# -*- coding:utf-8 -*-

class Solution:
    def reverseParentheses(self, s: str) -> str:
        open_stack = []
        i=0
        while (i <len(s)):
            if s[i] == "(":
                open_stack.append(i)
            elif s[i] == ")":
                start = open_stack.pop(-1)
                sub_string = s[start+1:i]
                if i == len(s)-1:
                    s = s[:start] + sub_string[::-1]
                else:
                    s = s[:start] + sub_string[::-1] + s[i+1:]
                i -= 2
            i += 1
        return s

    def reverseParenthesesfast(self, s):
        stack = [""]
        for i in s:
            if i == "(":
                stack.append("")
            elif i == ")":
                back = stack.pop()[::-1]
                stack[-1] += back
            else:
                stack[-1] += i
        return "".join(stack)

if __name__ == "__main__":
    sol = Solution()
    # print(sol.reverseParentheses("(abcd)"))
    # print(sol.reverseParentheses("(u(love)i)"))
    # print(sol.reverseParentheses("(ed(et(oc))el)"))
    print(sol.reverseParenthesesfast("ta()usw((((a))))"))