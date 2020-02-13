import unittest
from main import get_max_anagrams
from itertools import combinations


class MyTestCase(unittest.TestCase):
    def test_anagrams_count(self):
        self.assertEqual(len(get_max_anagrams()), 6)

    def test_anagrams(self):
        max_anagrams = get_max_anagrams()
        for ana in max_anagrams:
            for a, b in combinations(ana, 2):
                self.assertListEqual(sorted(a), sorted(b))

    def test_lists(self):
        output_list = get_max_anagrams()
        expected_list = [['abel', 'able', 'bale', 'bela', 'elba'],
                         ['alger', 'glare', 'lager', 'large', 'regal'],
                         ['angel', 'angle', 'galen', 'glean', 'lange'],
                         ['caret', 'carte', 'cater', 'crate', 'trace'],
                         ['elan', 'lane', 'lean', 'lena', 'neal'],
                         ['evil', 'levi', 'live', 'veil', 'vile']]
        sorted_output_list = list(map(sorted, output_list))
        sorted_expected_list = list(map(sorted, expected_list))
        self.assertListEqual(sorted_output_list, sorted_expected_list)


if __name__ == '__main__':
    unittest.main()
