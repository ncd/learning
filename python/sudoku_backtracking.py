import numpy as np

metrics = [[3, 1, 6, 5, 7, 8, 4, 9, 2],
           [5, 2, 9, 1, 3, 4, 7, 6, 8],
           [4, 8, 7, 6, 2, 9, 5, 3, 1],
           [2, 6, 3, 0, 1, 5, 9, 8, 7],
           [9, 7, 4, 8, 6, 0, 1, 2, 5],
           [8, 5, 1, 7, 9, 2, 6, 4, 3],
           [1, 3, 8, 0, 4, 7, 2, 0, 6],
           [6, 9, 2, 3, 5, 1, 8, 7, 4],
           [7, 4, 5, 0, 8, 6, 3, 0, 0]]

print(np.matrix(metrics))

def possible(x, y, n) :
  for i in range(9) :
    if metrics[x][i] == n:
      return False
  for j in range(9) :
    if metrics[i][y] == n:
      return False
  x0 = (x // 3) * 3
  y0 = (y // 3) * 3
  for i in range(3) :
    for j in range(3):
      if metrics[x0 + i][y0 + j] == n:
        return False
  return True

def solve():
  global metrics
  for i in range(9):
    for j in range(9):
      if metrics[i][j] == 0:
        for n in range(1, 10):
          if possible(i, j, n):
            metrics[i][j] = n
            solve()
            metrics[i][j] = 0
        return
  print(np.matrix(metrics))

print("solving")
solve()


