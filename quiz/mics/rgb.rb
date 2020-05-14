#!/usr/bin/env ruby

def rgb(r, g, b)
  r = 0 if r < 0
  g = 0 if g < 0
  b = 0 if b < 0
  
  r = 255 if r > 255
  g = 255 if g > 255
  b = 255 if b > 255

  (to_hex(r) + to_hex(g) + to_hex(b)).upcase
end

def to_hex(n)
  r = n.to_s(16)
  return "0#{r}" if r.size == 1
  r
end

def rgb(r, g, b)
  "%.2X%.2X%.2X" % [r,g,b].map{|i| [[i,255].min,0].max}
end


p rgb(0,0,0)
p rgb(255, 255, 255)
p rgb(148, 0, 211) # returns 9400D3
