#!/usr/bin/env ruby

# The parameter of the function findNb (find_nb, find-nb, findNb) will be an integer m 
# and you have to return the integer n such as n^3 + (n-1)^3 + ... + 1^3 = m if such a n exists or -1 if there is no such n.

def find_nb(m)
  n = 1
  while m > 0 
    m = m - (n ** 3)
    n += 1
  end
  (m == 0)? n : -1
end

def find_sol(m, n)
  mp = m - (n ** 3)
  if mp == 0
    return n
  elsif mp > 0
    find_sol(mp, n+1)
  else
    -1
  end
end

