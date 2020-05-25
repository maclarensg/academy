#!/usr/bin/env ruby

def fizzbuzz(n)
  (1..100).each do |i| 
    is_mult_3 = i % 3 == 0
    is_mult_5 = i % 5 == 0
    if is_mult_3 and is_mult_5
      print "FizzBuzz "
    elsif is_mult_3
      print "Fizz "
    elsif is_mult_5
      print "Buzz "
    else
      print "#{i} "
    end
  end
  puts
end

def main
  fizzbuzz(100)
end

if  __FILE__ == $0
  main
end