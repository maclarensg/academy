#!/usr/bin/env ruby
def unique_in_order(iterable)
  i = nil
  
  if iterable.is_a? String
    array = iterable.chars
  else 
    array = iterable
  end

  
  array.reject{ |e|  r = (e == i ? e : nil); i = e ; r }
end

def unique_in_order(iterable)
  (iterable.is_a?(String) ? iterable.chars : iterable).chunk { |x| x }.map(&:first)
end
# unique_in_order('AAAABBBCCDAABBB'), ['A','B','C','D','A','B']
p unique_in_order('AAAABBBCCDAABBB')