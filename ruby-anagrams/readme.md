# Ruby anagrams solution

Before running it, please make sure that:
  - You have configured ruby environment
  - Magic 🧙

It has two solutions - short and standart.

## Running

To run it:
- go to repo ```cd ./ruby-anagrams```
- run ```bundle install``` (for installing RSpec)
- short solution running: ```ruby -r "./anagrams.rb" -e "anagrams_resolver(shortly: true)"```
  <br/>Output: ```You are using short solution 🥳```
```[["abel", "able", "bale", "bela", "elba"], ["alger", "glare", "lager", "large", "regal"], ["angel", "angle", "galen", "glean", "lange"], ["caret", "carte", "cater", "crate", "trace"], ["elan", "lane", "lean", "lena", "neal"], ["evil", "levi", "live", "veil", "vile"]]```
- standart solution: ```ruby -r "./anagrams.rb" -e "anagrams_resolver```
  <br/>Output: ```Your are running standart solution 😏```
```[["abel", "able", "bale", "bela", "elba"], ["alger", "glare", "lager", "large", "regal"], ["angel", "angle", "galen", "glean", "lange"], ["caret", "carte", "cater", "crate", "trace"], ["elan", "lane", "lean", "lena", "neal"], ["evil", "levi", "live", "veil", "vile"]]```

#### Running tests
```rspec ./spec/anagrams_spec.rb```

If you want to improove in some way, please create PR - [Link](https://github.com/LvivScriptCommunity/rosetta-code-challenge-2/pulls)
