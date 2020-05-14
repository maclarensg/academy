#!/usr/bin/env ruby

def find_even_index(arr)
  left_sum = 0
  right_sum = arr.reduce(:+)
  

  arr.each_with_index do |e, ind|

    right_sum -= e
    
    return ind if left_sum == right_sum

    left_sum += e
  end
  
  -1  
end


p find_even_index([1,2,3,4,3,2,1])
#p find_even_index([1,100,50,-51,1,1])
#p find_even_index([20,10,30,10,10,15,35])
#p find_even_index([1,100,50,-51,1,1])