#!/usr/bin/env ruby

def high(x)
  high = 0 
  high_w = nil
  x.split(' ').each do |w|
    r = w.gsub(/[^a-z]/i, '').chars.map{ |c| c.downcase.ord - 96 }.sum
    if r > high
      high = r
      high_w = w
    end
  end
  high_w
end

p high('man i need a taxi up to ubud')