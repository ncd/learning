from collections import Counter

data = (5, 4, 3, 2, 1, 5, 4, 3, 2, 5, 4, 3, 5, 4, 5)
n = 3
cnt = Counter(data)
res = cnt.most_common()[-n]
print(res[0])