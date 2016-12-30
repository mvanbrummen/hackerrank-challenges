#!/usr/bin/ruby

n = gets.strip.to_i
a = Array.new(n)
for a_i in (0..n-1)
  a[a_i] = gets.strip.split(' ').map(&:to_i)
end
p = a.map.with_index {|row, i| row[i]} .inject :+
s = a.map.with_index {|row, i| row[-1-i]} .inject :+
puts "#{(p - s).abs}"
