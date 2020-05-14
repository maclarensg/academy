def spinWords(string)
  string.gsub(/\w{5,}/, &:reverse)
end


p spinWords("Hey fellow warriors") #, "Hey wollef sroirraw"