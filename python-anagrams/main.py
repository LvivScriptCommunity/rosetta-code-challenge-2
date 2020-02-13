import urllib.request


def get_max_anagrams():
    words = urllib.request.urlopen('http://wiki.puzzlers.org/pub/wordlists/unixdict.txt').read().decode().split()

    anagrams = {}
    for word in words:
        sorted_word = ''.join(sorted(word))
        anagrams.setdefault(sorted_word, []).append(word)

    max_count = max(len(ana) for ana in anagrams.values())
    max_anagrams = [ana for ana in anagrams.values() if len(ana) == max_count]

    return max_anagrams


if __name__ == '__main__':
    anagrams = get_max_anagrams()
    print(anagrams)
