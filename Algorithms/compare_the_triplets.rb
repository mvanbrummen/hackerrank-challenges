#!/bin/ruby

a = gets.strip.split(" ").map(&:to_i)
b = gets.strip.split(" ").map(&:to_i)
a_score, b_score = 0, 0
a.zip(b) do |a, b|
  if a < b
    b_score += 1
  elsif a > b
    a_score += 1
  end
end
puts "#{a_score} #{b_score}"
