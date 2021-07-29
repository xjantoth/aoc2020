#!/usr/bin/env python
import fileinput

t = str.maketrans("FBLR", "0101")
print(t)
s = set(int(l.translate(t), 2) for l in fileinput.input())
print(s)
lo, hi = min(s), max(s)

print(hi)
print(next(i for i in range(lo + 1, hi) if i not in s))
