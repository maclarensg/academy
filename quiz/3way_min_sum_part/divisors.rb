#!/usr/bin/env ruby

def get_divisors(n)
  div_list = []
  for i in 1..(Math.sqrt(n).to_i+1) 
    if (n % i == 0) 
      p "#{n} #{i}"
      div_list << (i) unless div_list.include? i
      div_list << (n / i).to_i unless div_list.include? (n / i).to_i
    end
  end
  return div_list.sort
end

def get_divisors2(n)
  div_list = []
   (1..(Math.sqrt(n).to_i+1)).map do |i|
    if n % i == 0
      [i,  (n / i).to_i ].uniq.each { |j| div_list << j unless div_list.include? j }
    end
  end
  return div_list.uniq.sort
end


p get_divisors2(12)
