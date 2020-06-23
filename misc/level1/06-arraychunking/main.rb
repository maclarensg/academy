# def chunk(a, n)
#   res = []
#   if a.size < n
#     return a 
#   else
#     i=0
    
#     while i < a.size
#       step = (i+n > a.size)?  a.size  : i+n 
#       tmp = a[i...step]
#       res << tmp
#       i = i + n
#     end
#   end
#   res
# end

def chunk(items, n)
  chunk = []
  while n < items.size
    items , chunk = items[n...items.size], chunk.append(items[0...n])
  end
  chunk = chunk.append(items)
  (chunk.size == 1) ? chunk.first :  chunk
end

p chunk([1,2,3,4], 2)
p chunk([1,2,3,4,5], 2)
p chunk([1,2,3,4,5,6,7,8], 3)
p chunk([1,2,3,4,5], 4)
p chunk([1,2,3,4,5], 10)