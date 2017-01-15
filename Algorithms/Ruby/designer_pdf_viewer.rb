#!/bin/ruby

h = gets.strip.split(' ').map(&:to_i)
word = gets.strip
h = word.each_byte.map { |i| h[i - 'a'.ord] }
puts h.max * word.length * 1
