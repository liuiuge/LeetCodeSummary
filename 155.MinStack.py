# -*- coding:utf8 -*-


class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.storage = []
        self.min = []
        

    def push(self, x: int) -> None:
        if len(self.min) < 1 or x < self.min[-1]:
            self.min.append(x)
        return self.storage.append(x)

    def pop(self) -> None:
        i = self.storage.pop()
        if i not in self.storage and self.min[-1] == i:
            self.min.pop()
        return None

    def top(self) -> int:
        return self.storage[-1]
        

    def getMin(self) -> int:
        return self.min[-1]
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()

if __name__ == "__main__":
    obj = MinStack()
    obj.push(3)
    obj.pop()
    obj.push(1)
    x = obj.top()
    obj.push(2)
    y = obj.getMin()
    print(x, y)