#!/usr/bin/env ruby

#def fib(m)
#  return m-1 if m == 1 or m == 2
#  sum = fib(m-1) + fib(m-2)
#end

#def fib(m)
#  return m-1 if m == 1 or m == 2
#  sum = fib(m-1) + fib(m-2)
#end

def productFib(prod)
  p fib((prod ** 0.5).to_i)
end

#p productFib(5895)

#def fib(m)
#  if m < 3
#    return m -1
#  else
#    array = [0, 1]
#    for i in 0 ... m - 2
#      array << array[-1] + array[-2]
#    end
#    array[-1]
#  end
#end
def fib(m)
  return m -1 if m < 3
  i,j = 0, 1
  (2...m).each do 
    i,j = j, j+i
  end
  j
end

p fib(1)
p fib(2)
p fib(3)
p fib(4)
p fib(5)
p fib(6)
p fib(7)
p fib(8)
p fib(9)
p fib(10)



def productFib(prod)
  a, b = [0, 1]
  while prod > a * b
    a, b = [b, a + b]
  end
  [a, b, prod == a * b]
end

