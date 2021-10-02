
def find_solution(n):
  s = [5, 10, 20, 50, 100, 200, 500]
  solution = 0
  def find(n, i):
    if i == len(s):
      return 0
    count = 0
    for k in range(0, n//s[i] + 1):
      if k*s[i] == n:
        count += 1
      elif k * s[i] < n:
        count += find(n - k*s[i], i+1)
    return count
  return find(n, 0)

print(find_solution(1000))