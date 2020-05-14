#!/usr/bin/env ruby

#require "mathn"
#def sumOfDivided(lst)
#  factors = Hash.new(0)
#  lst.each do |n|
#    n.abs.prime_division.each do |p,_|
#      factors[p] += n
#    end
#  end
#  factors.sort
#end

require 'prime'
def sumOfDivided(lst)
  a = {}
  lst.each{|x|
    p x
    Prime.prime_division(x.abs).each{|y|
      p y
      a[y[0]] = (a[y[0]]||0) + x
    }
  }
  p a
  a.to_a.sort_by{|x|x[0]}
end

p sumOfDivided([12, 15])