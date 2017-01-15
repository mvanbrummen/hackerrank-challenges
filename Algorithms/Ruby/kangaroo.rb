#!/bin/ruby

x1, v1, x2, v2 = gets.strip.split(' ').map(&:to_i)
if (x2 > x1 && v2 > v1) || v2 - v1 == 0
  puts "NO"
else
  puts (x1 - x2) % (v2 - v1) == 0 ? "YES" : "NO"
end
