#!/usr/bin/env ruby

def disemvowel(str)
  str.delete('aeiouAEIOU')
end

# For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

disemvowel("This website is for losers LOL!")