import sys
from collections import OrderedDict

dict_words = set(open("dictionary.lst").read().strip().lower().split("\n"))
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
unresolved = {word: set(dict_by_count[len(word)]) for word in unresolved_words}

def resolveNextWord(translation, unresolved_words):
    word = unresolved_words.pop(0)
    dict_with_len = dict_by_count[len(word)]
    if len(dict_with_len) == 1:
        winner = list(dict_with_len)[0]
        return applyWord(translation, word, winner)
    best_dict_word = None
    best_count = 0
    for dict_word in dict_with_len:
        answer = tryApplyWord(translation, unresolved_words, word, dict_word)
        if answer is None:
            continue
        if len(answer) > best_count:
            best_dict_word = dict_word
    if best_dict_word is None:
        return False
    else:
        return applyWord(translation, word, best_dict_word)


def tryApplyWord(translation, unresolved_words, word, dict_word):
    # return if assuming word is dict_word is possible.
    translation = dict(translation)
    ok = applyWord(translation, word, dict_word)
    if not ok:
        return None
    ok = solve(translation, list(unresolved_words))
    if ok:
        return translation
    else:
        return None

def applyWord(translation, word, dict_word):
    for i, c in enumerate(word):
        if c in translation:
            if translation[c] != dict_word[i]:
                return False
        else:
            translation[c] = dict_word[i]
    return True

def solve(translation, unresolved_words):
    while len(unresolved_words) > 0:
        ok = resolveNextWord(translation, unresolved_words)
        if not ok:
            return False
    return True

translation = {' ': ' '}
ok = solve(translation, unresolved_words)
if ok:
    for c in input_text:
        sys.stdout.write(translation[c])
    sys.stdout.write("\n")
else:
    print("could not solve")
