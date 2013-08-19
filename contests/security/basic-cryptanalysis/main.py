import sys

dict_words = set(open("dictionary.lst").read().strip().lower().split("\n"))
input_words = sys.stdin.read().strip().lower().split("\n")

dict_by_count = {}
for word in dict_words:
    s = dict_by_count.get(len(word), set())
    s.add(word)

unresolved_words = sorted(list(set(input_words)), reversed=True, key=lenFreq)
translation = {' ': ' '}

while True:
    word = unresolved_words.pop(0)
    if tryWord(translation, unresolved_words, word):
        print(word + " -> " + 
        for dict_word in dict_with_len:
            # try pretending this word is this dict_word and see if it works.
            if tryWord(translation, word, dict_word):


def tryWord(translation, unresolved_words, word):
    dict_with_len = dict_by_count[len(word)]
    if len(dict_with_len) == 1:
        # there's only 1 word with this length. we win!
        return list(dict_with_len)[0]
    for dict_word in dict_with_len:
        if tryApplyWord(translation, unresolved_words, word, dict_word):
            return dict_word
    return None

def tryApplyWord(translation, unresolved_words, word, dict_word):
    # return if assuming word is dict_word is possible.
    translation = dict(translation)
    ok = applyWord(translation, word, dict_word)
    if not ok:
        return False
    unresolved_words = list(unresolved_words)
    next_word = unresolved_words.pop(0)
    return tryWord(translation, unresolved_words, next_word) is not None

def applyWord(translation, word, dict_word):
    for c, i in word:
        if c in translation:
            if translation[c] != dict_word[i]:
                return False
        else:
            translation[c] = dict_word[i]
    return True

def lenFreq(word):
    return len(dict_by_count.get(len(word), set()))

def possibleWords(translation, word):
    # discount words that are impossible due to translation
    words_with_len = dict_by_count[len(word)]
    return [dword for dword in words_with_len if possible(dword)]

    def possible(dword):
        for c, i in word:
            if c in translation and translation[c] != dword[i]:
                return False
        return True
