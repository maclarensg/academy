#!/usr/bin/env ruby


require "prime"
 
# class Integer
#   def proper_divisors
#     return [] if self == 1
#     primes = prime_division.flat_map{|prime, freq| [prime] * freq}
#     p primes
#     (1...primes.size).each_with_object([1]) do |n, res|
#       primes.combination(n).map{|combi| res << combi.inject(:*)}
#     end.flatten.uniq
#   end
# end

class Integer
  def divisors
    return [] if self == 1
    r = [1]
    for i in 2..(self/2) do
      r << i if self % i == 0
    end
    r
  end
end

 
def min_sum_factorize(n, num_factors)
  if num_factors == 1
      return [n]
  end
  
  best = [n] + [1] * (num_factors - 1)
  best_sum = best.sum

  # slowly reduce the factorial.  num_factors if 3 means [x,y,z] = n 
  # so it reduces till we get only [z]
  # and if the new total is less than best_sum
  # register best  and  best_sum as we iterate thru the combinations
  n.divisors.each do |x|
    rest = min_sum_factorize( n/x, num_factors - 1)
    total = rest.sum + x
    if total < best_sum
      best = [x] + rest
      best_sum = total
    end
  end

  return best
end

p min_sum_factorize(1890, 3)