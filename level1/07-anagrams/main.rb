


# def sorted_string(s)
#   arr = []
  
#   s.downcase.each_byte.to_a.map { |e| 
#     if e >= 0 and e <= 9  or e >= 97 and e <= 122
#       arr << e
#     end
#   }
  
#   mergeSort(arr).map{ |e| e.chr }.join('')
# end

def sort_string_with_point(s)
  arr = []
  s.each_byte do |e|
    # downcase
    if e >= 65 and e <= 90
      e += 32
    end

    # accept alphanumeric
    if e >= 0 and e <= 9  or e >= 97 and e <= 122
      arr << e
    end
  end
  arr.sum
end

def anagrams(s1, s2)
  sort_string_with_point(s1) == sort_string_with_point(s2)
end

p anagrams('rail safety', 'fairy tales')
p anagrams('RAIL! SAFETY!', 'fairy tales')
p anagrams('Hi there', 'Bye there')