#!/usr/bin/env ruby


def evaluate(string)
  array = string.split(" ")
  calc(array)
end

  def calc(array)
    # find * or /
    while array.include? "*" or array.include? "/"
      (1..array.length-1).step(2).each do |i|
        if array[i] == '*' 
          mult(array, i)
        elsif array[i] == '/' 
          div(array, i)
        end
      end
    end

    # find + or -
    while array.include? "+" or array.include? "-"
      (1..array.length-1).step(2).each do |i|
        if array[i] == '+' 
          add(array, i)
          break
        elsif array[i] == '-' 
          minus(array, i)
          break
        end
      end
    end

    if array[0].is_a? String
      if array[0] =~ /\d+\.\d+/
        array[0].to_f
      else
        array[0].to_i
      end
    else
      array[0]
    end
  end

  def mult(array, i)
    a = array[i-1].to_i
    b = array[i+1].to_i
    array.delete_at(i+1)
    array.delete_at(i)
    array[i-1] = a * b
  end

  def div(array, i)
    a = array[i-1].to_i
    b = array[i+1].to_i
    array.delete_at(i+1)
    array.delete_at(i)
    if a % b == 0
      array[i-1] = a / b 
    else
      array[i-1] = a / b.to_f
    end
  end

  def add(array, i)
    a = array[i-1].to_i
    b = array[i+1].to_i
    array.delete_at(i+1)
    array.delete_at(i)
    array[i-1] = a + b
  end

  def minus(array, i)
    a = array[i-1].to_i
    b = array[i+1].to_i
    array.delete_at(i+1)
    array.delete_at(i)
    array[i-1] = a - b
  end


  def evaluate(string)
    [' + ', ' - ', ' * ', ' / '].each do |op|
      if string.include?(op)
        return string.split(op).map { |s| evaluate(s) }.inject(op.strip)
      end
    end
    string.to_f
  end

p evaluate "2 + 3 * 4 / 3 - 6 / 3 * 3 + 8"
