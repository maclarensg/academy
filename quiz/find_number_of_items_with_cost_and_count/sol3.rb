#!/usr/bin/env ruby

require 'slop'


opts = Slop::Options.new

opts.float '-C', '--cost', 'Total cost of items' , required: true
opts.integer '-c', '--count', 'Total count of items' , required: true
opts.string '-p', '--prices', 'list of prices e.g 6,2,0.1' , required: true

# Create a parser with above slop options
parser = Slop::Parser.new(opts)

# Begin parsing
begin
  options = (parser.parse(ARGV)).to_hash
rescue Slop::MissingRequiredOption => msg
  puts msg
  puts opts
  exit 1
end

# Covert string to list
options[:prices] = options[:prices].split(',').map{ |e| e.to_f }.sort.reverse
if options[:prices].size != 3
  puts "Need exactly 3 prices"
end

def find_max_of_each(prices, count, cost)
  max = prices.clone
  remainder = cost - prices.sum 
  max.map do |e| 
    r = (remainder/e).to_i + 1
    possible = ( count - prices.size )
    r < possible ? r : possible
  end
end

def solve(cost, count, prices)
  max_of_each = find_max_of_each(prices, count, cost)
  for x in 1..max_of_each[0]
    (max_of_each[1]).downto(1).each do |y|
      z = count - x -y
      sum = prices[0] * x + prices[1] * y + prices[2] * z
      p sum
      
      if sum == cost
        return [x,y, z]
      end

      if sum < cost
        x += 1
        # break loop if x exceeds max_x as there is no solution 
        if x > max_of_each[0]
          return []
        end
      end
    end
  end
  []
end

result = solve(options[:cost], options[:count], options[:prices])

if result.empty? 
  puts "no solution"
else
  p result
end