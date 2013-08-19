import sys
import string
from collections import OrderedDict

#DICT_FILE = 'dictionary.lst'
DICT_FILE = '/usr/share/dict/words'

dict_words = set(open(DICT_FILE).read().strip().lower().split("\n"))
input_text = sys.stdin.read().strip().lower()
input_words = input_text.split()

dict_by_count = {}
for word in dict_words:
    s = dict_by_count.get(len(word), set())
    s.add(word)
    dict_by_count[len(word)] = s

def lenFreq(word):
    return len(dict_by_count.get(len(word), set()))
unresolved_words = sorted(list(set(input_words)), key=lenFreq)

# map input words to sets of possible dictionary words
unresolved = OrderedDict([(word, set(dict_by_count[len(word)])) for word in unresolved_words])


# {'asda': ['shut', 'pike', 'sits']}
# 'a': ['s', 'p']
# 's': ['h', 'i']
# 'd': ['u', 'k', 't']
# 'a': ['t', 'e', 's']
impossible = {}
alphabet = set(string.ascii_lowercase)
again = True

def learnImpossible():
    # TODO: take into account repeat letters in the same word
    # learn impossible letter transformations
    for word, possible_set in unresolved.items():
        for i in range(len(word)):
            source = word[i]
            dest = alphabet - set([w[i] for w in possible_set])
            current_set = impossible.get(source, set())
            current_set.update(dest)
            impossible[source] = current_set

def eliminateImpossible():
    def notPossible(source, dest):
        for i, sc in enumerate(source):
            dc = dest[i]
            if dc in impossible[sc]:
                again = True
                return True
        return False
    for word, possible_set in unresolved.items():
        possible_set.difference_update([w for w in possible_set if notPossible(word, w)])

while again:
    learnImpossible()
    again = False
    eliminateImpossible()
    
for c in string.ascii_lowercase:
    print(c, set(string.ascii_lowercase) - impossible.get(c, set()))
#print(impossible)
print(unresolved)
#for c in input_text:
#    sys.stdout.write(translation[c])
#sys.stdout.write("\n")
