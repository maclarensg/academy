

# def pyramid(n)
#   array = Array.new(n)
#   array = array.each_with_index.map do  |e, i| 
#     " " * (n-i-1) + ("#" * (i*2+1)) + " " * (n-i-1)
#   end
#   array.each { |l| puts l}
# end 


def pyramidR(n, size)
  if n > 0 
    pyramidR(n-1, size)
  end
  # p "#{size}, #{n}"
  puts  "'" +(" " * (size-n)) + ("#" * (n*2+1)) + (" " * (size-n)) + "'"
end

def pyramid(n)
  pyramidR(n, n)
end

pyramid 5