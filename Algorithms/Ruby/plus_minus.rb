#!/bin/ruby

n = gets.strip.to_f
arr = gets.strip.split(' ').map(&:to_i)
puts arr.count { |i| i > 0 } / n
puts arr.count { |i| i < 0 } / n
puts arr.count { |i| i == 0 } / n
