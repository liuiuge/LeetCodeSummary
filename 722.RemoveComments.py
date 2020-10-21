# -*- coding:utf-8 -*-

class Solution:
    def removeComments(self, source):
        comment_mode = False
        multi_line_mode = False
        
        result = []
        line_to_add = []
        
        for line in source:
            i = 0
            while i < len(line):
                if not comment_mode:
                    if line[i:i+2] == "/*" or line[i:i+2] == "//":
                        comment_mode = True
                        if line[i+1] == "*":
                            multi_line_mode = True
                        i += 1
                    else:
                        line_to_add.append(line[i])
                else:
                    if not multi_line_mode:
                        break
                    if i+1 < len(line) and line[i] == '*' and line[i+1] == '/' :
                        comment_mode = False
                        multi_line_mode = False
                        i += 1
                i += 1
            if not multi_line_mode:
                comment_mode = False
                if line_to_add:
                    result.append(''.join(line_to_add))
                    line_to_add = []
        
        return result
            

if __name__ == "__main__":
    sol = Solution()
    print(sol.removeComments(
        ["/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", 
            "/* This is a test", "   multiline  ", "   comment for ", "   testing */", 
            "a = b + c;", "}"]))
