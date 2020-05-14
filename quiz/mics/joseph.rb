#!/usr/bin/env ruby



def josephus(items,k)
  Array.new(items.size){ items.rotate!(k).pop }
end

def josephus(items, k)
  res = []
  while not items.empty?
    for i in 1..k
      p items
      items << items.shift
    end
    res << items.pop
  end
  res
end

p josephus([1,2,3,4,5,6,7],3)

p josephus(["C","o","d","e","W","a","r","s"],4)






