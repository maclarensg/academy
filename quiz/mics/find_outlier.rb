#!/usr/bin/env ruby

def find_outlier(integers)
  even_n, even_c = nil, 0
  odd_n,  odd_c  = nil, 0
  integers.each do |e|
    if e % 2 == 0 
      even_n = e
      even_c += 1
    else
      odd_n = e
      odd_c += 1
    end
  end
  even_c > odd_c ? odd_n : even_n
end

def find_outlier(integers)
  integers.partition(&:odd?).find(&:one?).first
end

p find_outlier([0, 1, 2])
p find_outlier([1, 2, 3]) 
p find_outlier([2,6,8,10,3])