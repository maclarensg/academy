# def maxchar(s)
#   hash = {}
  
#   s.split('').each do |c|
#     unless hash.has_key? c
#       hash[c] = 1
#     else
#       hash[c] += 1
#     end
#   end
#   p hash.max_by{ |k,v| v }.first
# end


def maxchar(s)
  hash = {}
  max = 0
  maxc = nil
  s.split('').each do |c|

    hash[c] = (hash[c] == nil) ?  1 : (hash[c]+=1) # ternary expression

    if hash[c] > max
      max = hash[c]
      maxc = c 
    end
  end

  maxc
end

p maxchar("abccccccd") #== "c"
p maxchar("apple 123111") #== "1"
