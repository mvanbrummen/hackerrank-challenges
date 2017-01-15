#!/bin/ruby

arr = gets.strip.split(' ').map(&:to_i).sort
min = arr[0, 4].reduce(&:+)
max = arr[1, 4].reduce(&:+)
puts "#{min} #{max}"
