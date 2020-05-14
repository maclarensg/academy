#!/usr/bin/env ruby

# Get cost and count 
if ARGV.count < 2 
  puts "./find <COST> <COUNT>"
  exit
end

# add numeric method to String class
class String
  def numeric?
    return true if self =~ /^\d+$/
    true if Float(self) rescue false
  end
end 

#  check args if numeric
if ARGV[0].numeric? and ARGV[1].numeric?
  $cost=ARGV[0].to_i
  $count=ARGV[1].to_i
else
  puts "./find <COST> <COUNT>\nCost and Count needs to be numeric"
  exit
end

# Puts args
puts "cost:#{$cost} count:#{$count}"

# Find max of x and y
$max_x = ($cost - 3 - 1)/6
$max_y = ($cost - 6 - 1)/3

# If max of x or y is great than count, that means no solution
if $max_x > ($count-2) or $max_y > ($count-2)
  puts "no solution"
  exit
end

# Puts max of x and y
puts "max_x=#{$max_x} max_y=#{$max_y}" 

# Counter for recursion
$r_count = 0

def search(x , y)
  $r_count += 1
  z = $count - x - y
  puts "X: #{x} Y:#{y} Z:#{z} Cost = #{6*x + 3*y + 0.1*z}"
    
  if 6*x + 3*y + 0.1* ($count - x - y) == $cost
    puts "X: #{x} Y: #{y} Z: ß#{z} - Recurse: #{$r_count}"
    exit
  end
  
  # recursion exit condition
  return if y < 1           # this exits recursion while decreasing y
  return if x > $max_x      # this exits recursion while increasing x
  
  # If we combine the below, and start from x=1, y=max of y. We will eventually find a solution (if exist).
  if 6*x + 3*y + 0.1* ($count - x - y) > $cost
    search(x , y-1)         # we decrease y till we the cost is lower the total cost (decrease y will increase z)
  else 
    search(x+1 , y)         # if the cost is lower then total cost we increase x   
  end
  
end

search(1 , $max_y)

# if search not found... prints below
puts "no solution"