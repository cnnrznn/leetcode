#!/usr/bin/python

import sys

list1 = []
list2 = []
data = []

for line in sys.stdin:
    data.append(line)

list1 = [int(x) for x in data[0].split()]
list2 = [int(x) for x in data[1].split()]

def merge2k(l1, l2, k):
    result = []

    while k > 0:
        k = k - 1

        if len(l1) == 0 and len(l2) == 0:
            return result
        elif len(l1) == 0:
            result.append(l2[0])
            l2 = l2[1:]
        elif len(l2) == 0:
            result.append(l1[0])
            l1 = l1[1:]
        else:
            if l1[0] < l2[0]:
                result.append(l1[0])
                l1 = l1[1:]
            else:
                result.append(l2[0])
                l2 = l2[1:]

    return result

result = merge2k(list1, list2, 5)

print(list1)
print(list2)
print(result)
