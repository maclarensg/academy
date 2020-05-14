#!/usr/bin/env ruby 

def validBraces(braces)
  array = []

  braces.chars.each do |c|
    array.push(1) if c == "("
    array.push(2) if c == "{"
    array.push(3) if c == "["
    return false if c == ")" and array.pop != 1 
    return false if c == "}" and array.pop != 2
    return false if c == "]" and array.pop != 3
  end
  array.empty?
end

#def validBraces(braces)
#  while braces.gsub!(/(\(\)|\[\]|\{\})/,'') do; end
#  braces.empty?
#end

p validBraces( "[(a)]" )

