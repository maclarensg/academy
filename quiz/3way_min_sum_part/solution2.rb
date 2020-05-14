#!/usr/bin/env ruby


require "prime"
 
class Integer
  def proper_divisors
    return [] if self == 1
    primes = prime_division.flat_map{|prime, freq| [prime] * freq}
    (1...primes.size).each_with_object([1]) do |n, res|
      primes.combination(n).map{|combi| res << combi.inject(:*)}
    end.flatten.uniq
  end
end
 
def min_sum_factorize(n, num_factors)
  if num_factors == 1
      return [n]
  end
  
  best = [n] + [1] * (num_factors - 1)
  best_sum = best.sum

  
  n.proper_divisors.each do |x|
      if x ** num_factors > n
          next
      end
      if n % x == 0
          rest = min_sum_factorize( n/x, num_factors - 1)
          total = rest.sum + x
          if total < best_sum
              best = [x] + rest
              best_sum = total
          end
      end
  end
  return best
end

p min_sum_factorize(12, 3)