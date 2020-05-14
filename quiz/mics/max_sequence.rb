#!/usr/bin/env ruby

def max_sequence(arr)
  max_sum =  0
  for i in 0 ... arr.size 
    for j in 0 ... arr.size - i
      p  "#{arr[j..i+j]}"
      if arr[j..i+j].sum > max_sum
        max_sum = arr[j..i+j].sum
        max_seq = arr[j..i+j]
      end 
    end
  end
  max_sum
end

def max_sequence(array)
  (1..array.size).map { |i|  array.each_cons(i).map(&:sum) }.flatten.push(0).max
end

p max_sequence([-2, 1, -3, 4, -1, 2, 1, -5, 4])
#p max_sequence([11])