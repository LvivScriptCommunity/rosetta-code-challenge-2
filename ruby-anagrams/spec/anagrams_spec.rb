require_relative '../anagrams'

describe "Anagrams" do
  it "works" do
    expected_result = [
      ["evil", "levi", "live", "veil", "vile"],
      ["abel", "able", "bale", "bela", "elba"],
      ["elan", "lane", "lean", "lena", "neal"],
      ["alger", "glare", "lager", "large", "regal"],
      ["angel", "angle", "galen", "glean", "lange"],
      ["caret", "carte", "cater", "crate", "trace"],
    ]
    expect(expected_result).to match_array(anagrams_resolver)
  end
end
