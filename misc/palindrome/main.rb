#!/usr/bin/env ruby

def palindrome(str)
  return  ( "#{str}" == "#{reverse(str)}" ) 
end



def reverse(str)
  
  len = str.size
  
  if len > 0
    
    (0 ... len/2).to_a.reverse_each do |i|
      str[i], str[len-1-i] =  str[len-1-i] , str[i]     
    end

    return str
  end
  
  ""
end

p palindrome("abba")
p palindrome("asdf")