def fillingBlocks n
  a = c = d = 1
  b = 0
  p ">>>> a: #{a}, b: #{b}, c: #{c}, d: #{d}"
  n.times{
    p "< a: #{a}, b: #{b}, c: #{c}, d: #{d}"
    p d
    a, c, d = b, d, d-a+b+5*b=c
    p "> a: #{a}, b: #{b}, c: #{c}, d: #{d}"
    p "--------------"
  }
  c
end

p fillingBlocks 4