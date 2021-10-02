#!/bin/python3

import math
import os
import random
import re
import sys
import numpy

# Complete the organizingContainers function below.
def organizingContainers(container):
    con = [0]*(len(container))
    typ = [0]*(len(container))
    for i in range(0, len(container)):
        for j in range(0, len(container)):
            con[i] += container[i][j]
            typ[i] += container[j][i]
    print(numpy.matrix(container))
    print(con)
    print(typ)
    for i in range(0, len(con)):
      for j in range(0, len(typ)):
        if con[i] != typ[i]:
            return 'Impossible'
    return 'Possible'

if __name__ == '__main__':
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        n = int(input())

        container = []

        for _ in range(n):
            container.append(list(map(int, input().rstrip().split())))

        result = organizingContainers(container)

        print(result)

    # fptr.close()
