#!/usr/bin/env ruby

def permutations(string)
  string.chars.permutation.map(&:join).uniq
end

p permutations("aabb")
