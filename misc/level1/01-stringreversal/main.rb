#!/usr/bin/env ruby

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

p reverse('apple');