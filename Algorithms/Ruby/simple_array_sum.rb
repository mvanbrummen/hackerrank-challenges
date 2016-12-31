#!/bin/ruby

n = gets.strip.to_i
puts gets.strip.split(' ').map(&:to_i).reduce(:+)
