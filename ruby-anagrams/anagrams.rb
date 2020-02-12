require 'open-uri'

def anagrams_resolver(shortly: false)
  anagrams = load_data
  result = []

  if shortly
    result = short_solution(anagrams)
  else
    result = standart_solution(anagrams)
  end
  puts "\n"
  p result
  result
end

def load_data
 anagrams = Hash.new { |hash, key| hash[key] = [] }

 open('https://raw.githubusercontent.com/R4meau/46-simple-python-exercises/master/exercises/unixdict.txt') do |f|
   words = f.read.split

   words.each do |word|
     anagrams[word.split("").sort.join] << word
   end
 end
 anagrams
end

def short_solution(anagrams)
  puts "You are using short solution 🥳"
  anagrams.values.group_by(&:size).max.last
end

def standart_solution(anagrams)
  puts "Your are running standart solution 😏"
  result = []
  counter = anagrams.values.map(&:length).max

  anagrams.each_value do |ana|
    result << ana if ana.length >= counter
  end
  result
end
