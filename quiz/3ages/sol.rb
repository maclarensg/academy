#!/usr/bin/env ruby
#c = a + 2b
#2a + 3b = c+3

#1. 2a + 3b = a + 2b +3 => Substitue c into 2nd equate
#2. a + b = 3  
#3. a = 3 - b

#4. c = (3 - b) + 2b  => Substitue a into 1st equate
#5. c = 3 - b + 2b
#6. c - 3  = b
#7. c = b + 3
#8. c = b + (a + b) # either a is 2 and b is 1 


possibles = [[1,2],[2,1]]

possibles.each do |tuple|
  a = tuple[0]
  b = tuple[1]
  c = a + 2 * b
  if  2*a + 3*b == c+3
    p [a,b,c]
    break
  end
end




