#!/usr/bin/env ruby


def buildtree(n)
  (1..n).map do |t|
    (1..5).map do |e|
      spacing =(((5-1)+1)-e) + (n-t)
      stars = (((e-1)*2)+1) + (t-1) *2
      " " * spacing + "*" * stars +  " " * spacing
    end
  end
end


buildtree(5).each do |t|
  t.each do |l|
    p l
  end
end