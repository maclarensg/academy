#!/usr/bin/env ruby
class String
  def camelcase
    self.split(" ").map{ |e| e.capitalize }.join
  end
end


p "ali baba".camelcase

p " camel case word".camelcase
