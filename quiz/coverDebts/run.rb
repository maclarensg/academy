#!/usr/bin/env ruby

def coverDebts(s, debts, interests)

  spend = 0
  stack = interests.each_with_index.sort.reverse.map { |t| t[1] }
    
  while debts.sum > 0
    budget = s / 10.0
    
    stack.each do |i|
      if budget >= debts[i]
        budget = budget - debts[i]
        spend += debts[i]
        debts[i] = 0
      end
      if debts[i] > budget
        debts[i] = debts[i] - budget
        spend += budget
        budget = 0
      end
    end
    debts = debts.zip(interests).map{|x, y| x + x * y / 100.0}
  end
  spend = spend % 1 > 0 ? spend : spend.to_i
end


s = 50
debts = [2, 2, 5] 
interests = [200, 100, 150]
s = 40
debts = [2, 2, 5]
interests = [75, 25, 25]
p coverDebts(s, debts, interests)