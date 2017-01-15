#!/usr/bin/ruby

require 'date'

time = gets.strip
parsed_time = DateTime.strptime(time, '%I:%M:%S%p')
puts parsed_time.strftime("%H:%M:%S")

