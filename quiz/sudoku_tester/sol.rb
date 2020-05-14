#!/usr/bin/env ruby


puzzle = [
  [5, 3, 4, 6, 7, 8, 9, 1, 2],
  [6, 7, 2, 1, 9, 5, 3, 4, 8],
  [1, 9, 8, 3, 4, 2, 5, 6, 7],
  [8, 5, 9, 7, 6, 1, 4, 2, 3],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 6, 1, 5, 3, 7, 2, 8, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 4, 5, 2, 8, 6, 1, 7, 9]
]

puzzle2 = [
  [5, 3, 4, 6, 7, 8, 9, 1, 2], 
  [6, 7, 2, 1, 9, 0, 3, 4, 8],
  [1, 0, 0, 3, 4, 2, 5, 6, 0],
  [8, 5, 9, 7, 6, 1, 0, 2, 0],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 0, 1, 5, 3, 7, 2, 1, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 0, 0, 4, 8, 1, 1, 7, 9]
]

puzzle3=[
  [1, 2, 3, 4, 5, 6, 7, 8, 9], 
  [2, 3, 4, 5, 6, 7, 8, 9, 1], 
  [3, 4, 5, 6, 7, 8, 9, 1, 2], 
  [4, 5, 6, 7, 8, 9, 1, 2, 3], 
  [5, 6, 7, 8, 9, 1, 2, 3, 4], 
  [6, 7, 8, 9, 1, 2, 3, 4, 5], 
  [7, 8, 9, 1, 2, 3, 4, 5, 6], 
  [8, 9, 1, 2, 3, 4, 5, 6, 7], 
  [9, 1, 2, 3, 4, 5, 6, 7, 8]
]

def validSolution(board)
  board.each do |row|
    return false if row.uniq.size != 9 
  end
  board.transpose.each do |row|
    return false if row.uniq.size != 9 
  end
  for i in 0...3
    for j in 0...3
      array = []
      for y in 0...3
        for x in 0...3
          unless array.include? board[y+(3*i)][x+(3*j)] 
            array << board[y+(3*i)][x+(3*j)] 
          else
            return false
          end
        end
      end
    end
  end
  true
end

p puzzle.map{|row| row.each_slice(3).to_a }.transpose.flatten.each_slice(9).to_a

p validSolution(puzzle3)