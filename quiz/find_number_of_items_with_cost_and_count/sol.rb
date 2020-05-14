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

# method to calculate cost
def cal_cost(x, y)
	6*x + 3*y + 0.1* ($count - x - y)
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
# I've purposely roundup value of z, so that z takes 1 unit
$max_x = ($cost - 3 - 1)/6
$max_y = ($cost - 6 - 1)/3
 
# If max of x is great than count, that means no solution
if $max_x > ($count-2) 
	puts "no solution"
	exit
end
 
# Puts max of x and y
puts "max_x=#{$max_x} max_y=#{$max_y}"
 
# Counter for recursion
$counter = 0
 
x=1

for y in ($max_y).downto(1)
	
	current_cost = cal_cost(x, y)
	$counter += 1
	z = $count - x - y
	puts "X: #{x} Y:#{y} Z:#{z} Cost = #{6*x + 3*y + 0.1*z}"
	
	# if current_cost is higher than the actual cost skip loop
	next if current_cost > $cost
    
	# if cost is lower than actual cost search x 
	if current_cost < $cost
		x += 1
		# break loop if x exceeds max_x as there is no solution 
		if x > $max_x
			break
		end
	end
  
	# Solution is found
	if current_cost == $cost
		puts "X: #{x} Y:#{y} Z:#{z} - Interations: #{$counter}" 
		exit
	end
end 
# if search not found... prints below
puts "no solution"