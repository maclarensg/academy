def steps(n)
  for i in 0 ... n
    p ("#" * (i+1)  + " " * (n-1-i))
  end
end


steps(2)
steps(3)
steps(4)