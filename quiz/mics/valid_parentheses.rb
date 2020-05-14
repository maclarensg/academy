#!/usr/bin/env ruby

def valid_parentheses(string)
  array = []
  string.chars.each do |e|
    if e == '('
      array.push(1)
    elsif e == ')' and array.size > 0
      array.pop
    elsif e == ')' and array.size < 1
      return false
    end
  end
  return true if array.size == 0 
  false
end


def valid_parentheses(string)
  open = 0
  string.chars.each do |c|
    open += 1 if c == "("
    open -= 1 if c == ")"
    return false if open < 0
  end
  open == 0
end

p valid_parentheses("hi(hi)()")