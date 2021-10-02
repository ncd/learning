import time
from functools import wraps

# class myfunc:
#   def __new__(cls, func, /, *args, **keywords):
#     print(type(func))
#     self = super(myfunc, cls).__new__(cls)
#     self.__name__ = func.__name__
#     self.func = func
  
#   def __call__(self, *args, **keywords):
#     print('xxxx')
#     self.func.__call__(*args, **keywords)

# def mywraps(func, assigned=WRAPPER_ASSIGNMENTS):
#   return myfunc(assigned)

def tracing(func):
  @wraps(func)
  def wrapper(*args, **kwargs):
    start = time.time()
    result = func(*args, **kwargs)
    end = time.time()
    print(func.__name__, end - start)
    return result
  return wrapper