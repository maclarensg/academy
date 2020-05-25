#!/usr/bin/env ruby



def reverse(x)
  flag = false
  if ( x < 0 )
    x = 0 - x
    flag = true
  end
  
  res = 0
  p = x
  
  while (p > 0)
    mod = p % 10
    p = p / 10
    res = res * 10 + mod
  end

  if flag 
    res = 0 - res
  end
  res
end



p reverse(15)