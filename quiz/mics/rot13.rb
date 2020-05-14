#!/usr/bin/env ruby

def rot13(string)
  string.chars.map { |e| norm_index(e) ? rot13_table(norm_index(e)) : e }.join
end

def norm_index(c)
  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz".index(c)
end

def rot13_table(n)
  "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"[n]
end


def rot13(string)
  string.tr("A-Za-z", "N-ZA-Mn-za-m")
end

p rot13("HelloWorld!")