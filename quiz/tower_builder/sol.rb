#!/usr/bin/env ruby

def towerBuilder(n_floors)
  (1..n_floors).map{ |e| 
    spacing =(((n_floors-1)+1)-e)
    " " * spacing + "*" * (((e-1)*2)+1) +  " " * spacing
  }
end

towerBuilder(25).each do |l|
  puts l
end
