#!/bin/ruby

n, k, q = gets.strip.split(' ')
n = n.to_i
k = k.to_i
q = q.to_i
a = gets.strip.split(' ').map(&:to_i)
k = k % n if k > n
#puts "k = #{k}"
a = a.drop(n - k) | a.take(n - k) if k != 0 || k != n
(0..q-1).each do
    m = gets.strip.to_i
    puts a[m]
end
