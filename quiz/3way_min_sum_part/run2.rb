#!/usr/bin/env ruby




require "prime"
 
class Integer
  def proper_divisors
    return [] if self == 1
    primes = prime_division.flat_map{|prime, freq| [prime] * freq}
    (1...primes.size).each_with_object([1]) do |n, res|
      p "#{n}, #{res}"
      primes.combination(n).map{|combi| res << combi.inject(:*)}
    end.flatten.uniq
  end
end

p 100.proper_divisors