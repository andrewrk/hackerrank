import sys
import string
from collections import OrderedDict

lines = sys.stdin.read().upper().strip().split("\n")
lines.pop(0)
while len(lines):
    key = lines.pop(0)
    cipherText = lines.pop(0)
    
    first = list(OrderedDict.fromkeys(key))
    alphabet = list(OrderedDict.fromkeys(key + string.ascii_uppercase))[len(first):]
    arrays = [[c] for c in first]
    next_index = 0
    while len(alphabet):
        arrays[next_index].append(alphabet.pop(0))
        next_index = (next_index + 1) % len(arrays)
        
    arrays = sorted(arrays, key=lambda a: a[0])
    mapping = {}
    alphabet = list(string.ascii_uppercase)
    for array in arrays:
        for c in array:
            mapping[c] = alphabet.pop(0)

    for c in cipherText:
        sys.stdout.write(mapping.get(c, c))
    sys.stdout.write("\n")
