#!/usr/bin/env ruby
def next_bigger(n)
  max = maxed(n)
  (n+1..max).each { |i| return i if max == maxed(i) }
  -1
end

def maxed(n)
  n.to_s.chars.sort.reverse.join.to_i
end 

p next_bigger(52189971297442)

