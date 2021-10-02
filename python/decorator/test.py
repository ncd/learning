from tools import tracing

@tracing
def test(n: int):
  while n > 0:
    n -= 1

if __name__ == "__main__":
  test(10000)