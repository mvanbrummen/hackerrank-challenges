#!/bin/ruby

n = gets.strip.to_i
(1..n).each { |i| puts "#{'#' * i}".rjust(n)}
