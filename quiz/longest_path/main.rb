# def longestPath(fileSystem)
#   root = nil
#   stack = 0
#   max_count = 0
#   stack_count = 0
#   fileSystem.split("\f").each do |l|
#     if l =~ /^(\t)+(.*)/
#       word = $2
#       s = l.scan(/\t/m).size
#       if s > stack
#         stack_count += 1 + word.size
#         stack = s
#       else # s.size <= stack
#         stack_count = 1 + word.size
#         stack = s
#       end
#     else
#       root = l
#     end 

#     if  l =~ /(\w+\.\w+)/
#       if (root.size + stack_count) > max_count
#         max_count = (root.size + stack_count)
#       end
#     end
#   end

#   max_count
# end


def longestPath(fileSystem)
  track=[0]
  fileSystem.split("\f").map{|i|
    depth=i.count "\t"
    count = track[depth] + i.size - depth
    track[depth+1] = 1 + count
    i["."] ? count : 0
  }.max
end

fileSystem = "user\f\tpictures\f\tdocuments\f\t\tnotes.txt"
p longestPath(fileSystem)