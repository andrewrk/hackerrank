import sys
import string
import random
from collections import OrderedDict

#DICT_FILE = 'dictionary.lst'
DICT_FILE = '/usr/share/dict/words'

dict_words = set(open(DICT_FILE).read().strip().lower().split("\n"))
input_text = sys.stdin.read().strip().lower()
input_words = set(input_text.split())

word_subsets = set()
for dict_word in dict_words:
    for i in range(len(dict_word)):
        word_subsets.add(dict_word[0:i])


def score_translation(translation):
    score = 0
    for word in input_words:
        translated = translate(translation, word)
        i = len(translated)
        while i > 0 and translated[0:i] not in word_subsets:
            i -= 1
        score += (i / len(translated)) ** 2 * len(translated)
    return score

def random_translation():
    translation = {}
    alphabet = list(string.ascii_lowercase)
    random.shuffle(alphabet)
    for c in string.ascii_lowercase:
        translation[c] = alphabet.pop()
    return translation

def translate(translation, text):
    buf = ""
    for c in text:
        buf += translation.get(c, c)
    return buf

surv_count = 10
random_surv_count = 1
offspring_count = 16
mut_chance = 0.01
generation_limit = 90000
gen_size = offspring_count * (surv_count + random_surv_count)
translation_set = [random_translation() for i in range(gen_size)]
score_set = [None for i in range(gen_size)]
generation_count = 0
next_generation_index = 0

def makeBabies(translation):
    global next_generation_index
    for i in range(offspring_count):
        specimen = dict(translation)
        translation_set[next_generation_index] = specimen
        for letter in string.ascii_lowercase:
            if random.random() < mut_chance:
                # swap letter with another
                c = string.ascii_lowercase[random.randint(0, len(string.ascii_lowercase)-1)]
                tmp = specimen[c]
                specimen[c] = specimen[letter]
                specimen[letter] = tmp

        next_generation_index += 1


best_score = 0
best_translation = None
best_output = ""
while generation_count != generation_limit:
    generation_count += 1
    generation_max_score = 0
    for i, translation in enumerate(translation_set):
        score = score_translation(translation)
        score_set[i] = (score, translation)
        if score > best_score:
            best_score = score
            best_translation = translation
            best_output = translate(translation, input_text)
            print("New best output (score", best_score, "):", best_output)
        if score > generation_max_score:
            generation_max_score = score
    #print("generation", generation_count, "max", generation_max_score)
    
    # take a subset of the top scoring programs and breed them to get a new
    # set of programs to evaluate
    score_set.sort(reverse=True, key=lambda x: x[0])
    next_generation_index = 0
    for survivor_index in range(surv_count):
        _, translation = score_set[survivor_index]
        makeBabies(translation)

    # take some random programs and breed them
    for i in range(random_surv_count):
        makeBabies(translation_set[random.randint(0, len(translation_set)-1)])

print(best_output)
