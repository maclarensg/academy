def matrix(n)

  padding=(n*n).to_s.size
  count = 1
  directions = [[0,1],[1,0],[0,-1],[-1,0]] #right, down, left, up
  array = Array.new(n).map {|e| Array.new(n)}
  
  i = 0
  x, y = 0, 0
  d = 0
  while i < n * n
    lx, ly = x, y 
  
    array[y][x] = "%0#{padding}d" % (i+1)

    y, x = addMatrix([ly,lx],directions[d])

    if x > n - 1 && d == 0                    # on the top right most point of the matrix turn
      x = n - 1
      d = (d + 1) % 4
      y, x = addMatrix([ly,lx],directions[d])
    elsif y > n - 1 && d == 1                 # on the bottom right most point of the matrix turn
      y = n - 1
      d = (d + 1) % 4
      y, x = addMatrix([ly,lx],directions[d]) 
    elsif x < 0 && d == 2                     # on the bottom left most point of the matrix turn
      x += 1
      d = (d + 1) % 4
      y, x = addMatrix([ly,lx],directions[d])
    # because we should never hit the start or top left most of the matrix
    # this condition should never hit
    # y < 0 && d == 3
    end
   
    # on the next point in matrix is already take up, back of and turn 
    if array[y][x] != nil
      d = (d + 1) % 4
      y, x = addMatrix([ly,lx],directions[d])
    end

    i+=1
  end
  printMatrix array
end

def addMatrix(a, ap)
  [ a[0] + ap[0], a[1] + ap[1] ]
end 


def printMatrix(a)
  a.each do |r|
    p r
  end
end

matrix(3)

