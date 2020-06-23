def matrix(n)

  padding=(n*n).to_s.size
  directions = [[0,1],[1,0],[0,-1],[-1,0]] #right, down, left, up
  array = Array.new(n).map {|e| Array.new(n)}
  
  i = 0
  x, y = 0, 0
  d = 0
  while i < n * n
    lx, ly = x, y 
  
    array[y][x] = "%0#{padding}d" % (i+1)

    y, x = shiftPos([ly,lx],directions[d])

    # If shifting of position in matrix is
		# - exceeds right boundary
		# - exceeds bottom boundary
		# - exceeds left boundary
		# change direction from the last position and shift
    if ( x > n - 1 && d == 0 ) or ( y > n - 1 && d == 1 )  or ( x < 0 && d == 2 )                 
      d = (d + 1) % 4
      y, x = shiftPos([ly,lx],directions[d])
    end
   
    # on the next point in matrix is already take up, back of and turn 
    if array[y][x] != nil
      d = (d + 1) % 4
      y, x = shiftPos([ly,lx],directions[d])
    end

    i+=1
  end
  printMatrix array
end

def shiftPos(a, ap)
  [ a[0] + ap[0], a[1] + ap[1] ]
end 


def printMatrix(a)
  a.each do |r|
    p r
  end
end

matrix(3)

