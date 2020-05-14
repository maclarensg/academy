def capitalize(str)
  str.split(" ").map { |e| e.capitalize }.join(" ")
end

p capitalize("a short sentence")
p capitalize("a lazy fox")
p capitalize("look, it is working!")