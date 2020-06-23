def vowels(s)
  s.downcase.split("").count{ |e|  ["a","e","i","o","u"].include? e }
end


p vowels("Hi There!")
p vowels("Why do you ask?")
p vowels("Why?")