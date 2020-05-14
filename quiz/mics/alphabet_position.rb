#!/usr/bin/env ruby

def alphabet_position(text)
  text.chars.map do |e|
    r = e.downcase.unpack('c').first - 96 
    if r > 0 and r < 27 
      r
    else
      nil
    end
  end.compact.join(' ')
end

def alphabet_position(text)
  text.gsub(/[^a-z]/i, '').chars.map{ |c| c.downcase.ord - 96 }.join(' ')
end


p alphabet_position("The sunset sets at twelve o' clock.")


