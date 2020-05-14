#!/usr/bin/env ruby 


def count_smileys(arr)
  arr.map do |s| 
    if s.size == 2
      (has_valid_eyes?(s[0]) and has_valid_mouth?(s[1]))? 1 : 0
    else
      ( has_valid_eyes?(s[0]) and has_valid_nose?(s[1]) and has_valid_mouth?(s[2]) )? 1 : 0
    end
  end.sum
end

def has_valid_eyes?(e)
  e == ':' or e == ';' 
end

def has_valid_nose?(e)
  e == '-' or e == '~' 
end

def has_valid_mouth?(e)
  e == ')' or e == 'D' 
end


def count_smileys(arr)
  arr.count { |e| e =~ /(:|;){1}(-|~)?(\)|D)/ }
end

p count_smileys([" )", ":P", "-(", "8P", "~(", "8~P", "8P", ": D", "; (", ":D", "P", ": D", "8 D", ";P", "P", ": (", ";-)", "8)", "8 )", ": )", "-)", ":~P", " P"])