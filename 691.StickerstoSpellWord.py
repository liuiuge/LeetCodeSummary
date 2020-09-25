import collections

class Solution():
    def minStickers(self, stickers, target):
        t_count = collections.Counter(target)
        print(t_count)
        A = [collections.Counter(sticker) & t_count
             for sticker in stickers]
        print(A)

        for i in range(len(A) - 1, -1, -1):
            if any(A[i] == A[i] & A[j] for j in range(len(A)) if i != j):
                A.pop(i)
        print(A)

        self.best = len(target) + 1
        def search(ans = 0):
            if ans >= self.best: return
            if not A:
                if all(t_count[letter] <= 0 for letter in t_count):
                    self.best = ans
                return

            sticker = A.pop()
            print(ans, A, [(t_count[letter] - 1) // sticker[letter] + 1
                        for letter in sticker])
            used = max((t_count[letter] - 1) // sticker[letter] + 1
                        for letter in sticker)
            used = max(used, 0)

            for c in sticker:
                t_count[c] -= used * sticker[c]

            search(ans + used)
            for i in range(used - 1, -1, -1):
                for letter in sticker:
                    t_count[letter] += sticker[letter]
                search(ans + i)

            A.append(sticker)

        search()
        return self.best if self.best <= len(target) else -1

if __name__ == "__main__":
    sol = Solution()
    print(sol.minStickers(["with", "example", "science"], "thehat"))