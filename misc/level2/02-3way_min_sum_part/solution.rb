#!/usr/bin/env ruby

def solve(n)
   
end

def get_divisors(n)
    factors = [1]
    for i in 2..n 
        factors.append i if n % i == 0
    end
    return factors
end

def min_sum_factorize(n, num_factors)
    if num_factors == 1
        return [n]
    end
    
    best = [n] + [1] * (num_factors - 1)
    best_sum = best.sum

    
    get_divisors(n).each do |x|
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