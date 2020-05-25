# Fibonacci with Iteration
def fibI(n)
  i,j = 0, 1

  if n < 2 
    return [i,j][n]
  end
  (2..n).each do 
    i,j = j, j+i
  end
  j
end

# Fibonacci with Recursion
@cache = [0,1]
def fib(n)
  return @cache[n] if @cache[n]
  
  @cache[n] = fib(n-1) + fib(n-2)
end

# Main function
def main
  n =  ARGV[0].to_i
  #puts "fibonacci with iteration O(n): #{fibI(n)}"
  p fib(n)
  # puts "fibonacci with recursion O(2^n): #{fibR(n)}" #comment out due to slowness
end

# Program start
if  __FILE__ == $0
  main
end