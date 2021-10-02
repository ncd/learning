//https://leetcode.com/problems/n-queens-ii/submissions/
class Solution(object):
    def totalNQueens(self, n):
        """
        :type n: int
        :rtype: int
        """
        board = [x[:] for x in [[0] * n] * n]
        def feasible(i, j, n):
            for k in range(0, n):
                if board[i][k] == 1:
                    return False
                if board[k][j] == 1:
                    return False
            for k in range(-n, n):
                if i+k >= 0 and j+k >= 0 and i+k < n and j+k < n:
                    if board[i+k][j+k] == 1:
                        return False
                if i+k >= 0 and j-k >= 0 and i+k < n and j-k < n:
                    if board[i+k][j-k] == 1:
                        return False
            return True
        def solve(i, n):
            count = 0
            for j in range(0, n):
                if feasible(i, j, n):
                    board[i][j] = 1
                    if i == n-1:
                        count += 1
                    else:
                        count += solve(i+1, n)
                    board[i][j] = 0
            return count
        return solve(0, n)