#!/usr/bin/env ruby

def spiralize(size)
  a = Array.new(size).map{ |e| [0] * size }
               # right down   left   up
  directions = [[1,0],[0,1],[-1,0],[0,-1]]
  i = 0
  x, y = 0, 0
  while i < size * size
    d =  directions[i % 4 ]
    case d
    when [1,0]
      steps = find_steps(a[y], x) 
    when [0,1]
      steps = find_steps(a.transpose[x], y)
    when [-1, 0]
      steps = find_steps(a[y].reverse, x) 
    when [0, -1]
      steps = find_steps(a[y].reverse, x) 
    end
    x_p, y_p =  d
    steps.downto(0) do |n|
      a[y][x] = 1
      y += y_p if y < a.size - 1
      x += x_p if x < a.size - 1
      p "#{x},#{y}"
    end

    case d
    when [1,0]
      y += 1
    when [0,1]
      x -= 1
    when [-1, 0]
      #
    when [0, -1]
      #
    end
    i += 1 
  end

  a
end

def find_steps(array, start)
  p array
  index = nil
  for i in start ... array.size
    if array[i] == 1
      index = i
    end
  end
  return array.size - 1 if index.nil?
  (index - 2 ) < 0 ? nil : (index - 2 )
end



p spiralize 5

1111
0001
1001
1111

111
001 
111