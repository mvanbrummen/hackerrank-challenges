#!/bin/ruby

s, t = gets.strip.split(' ').map(&:to_i)
a, b = gets.strip.split(' ').map(&:to_i)
m, n = gets.strip.split(' ').map(&:to_i)
apples = gets.strip.split(' ').map(&:to_i)
oranges = gets.strip.split(' ').map(&:to_i)
puts apples.count { |i| (a + i).between?(s, t) }
puts oranges.count { |i| (b + i).between?(s, t) }
