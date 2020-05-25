#!/usr/bin/env ruby


require "set"

class Board
  def initialize(matrix)
    @matrix = matrix
  end

  def [] y
    @matrix[y]
  end

  def inspect
    @matrix.map do |row|
      row.inspect 
    end.join("\n")
  end

  def transpose
    @matrix.transpose
  end

  def matrix
    @matrix
  end

  def possibles(y,x)
    # sets at y,x : all the numbers already appear in row, column and block
    sets_at_xy = (rows[y] + columns[x] + blocks[which_yx_block(y,x)]).reject! { |x| x == 0 }.uniq.to_set
    (set9 - sets_at_xy).sort # {1..9} - {n sets already existed in y,x}  = remaining possibles
  end

  def rows
    @matrix
  end
  
  def columns
    @matrix.transpose
  end

  def blocks
    @matrix.map{ |r| r.each_slice(3).to_a }.transpose.flatten.each_slice(9).to_a
  end
  
  def which_yx_block(y, x)
    (x/3) * 3  + (y/3)
  end
  
  def set9 
    (1..9).to_set
  end
  
  def validate
    all_sections = rows + columns + blocks
    all_sections.all? { |section| contains_all_nine? section }
  end

  def contains_all_nine?(section)
    [1,2,3,4,5,6,7,8,9].to_set == section.to_set
  end

  def solve
    
    for y in 0...9
      for x in 0...9
        if @matrix[y][x] == 0
          possibles(y,x).each do |p|
            @matrix[y][x] = p
            solve
            @matrix[y][x] = 0
          end
          return
        end
      end
    end

    puts inspect
    exit
    #puts "------"
    # puts "Is valid solution? %s" % (validate ? "Yes" : "No")
    # print "More?(y/n) "
    # option = gets
    # if option =~ /[n|N][o|O]?/
    #   return
    # end
  end
end



if __FILE__ == $0
  load "./problem.rb"
  board = Board.new($problem[2])
  p board
  puts "-----"
  board.solve
end
