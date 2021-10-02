import numpy as np

metrics = [[0, 0, 0, 1, 0, 0, 0, 4, 0],
           [0, 0, 0, 0, 8, 6, 0, 2, 5],
           [3, 0, 8, 0, 0, 4, 0, 0, 0],
           [6, 3, 0, 0, 0, 0, 2, 5, 0],
           [0, 4, 1, 0, 0, 0, 0, 8, 0],
           [0, 2, 0, 0, 0, 0, 0, 0, 6],
           [0, 0, 0, 0, 6, 0, 4, 0, 0],
           [0, 8, 0, 3, 4, 2, 0, 0, 0],
           [0, 0, 0, 0, 0, 7, 9, 0, 0]]

print(np.matrix(metrics))

def possible(x, y, n) :
  for i in range(9) :
    if metrics[x][i] == n:
      return False
  for j in range(9) :
    if metrics[j][y] == n:
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


