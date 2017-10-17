require_relative './helpers'

large_file = ARGV[0] || "large_file.txt"
n_number = ARGV[1] || 10
@result = Array.new(n_number, 0)

def meta (number, highest_numbers)
  if number > highest_numbers.first
    highest_numbers.unshift(number)
    highest_numbers.pop
  elsif number > highest_numbers.last 
    unless highest_numbers.include?(number)
      highest_numbers.pop
      highest_numbers.push(number)
    end
  end
  highest_numbers.sort!.reverse!
end

print_memory_usage do
  print_time_spent do
    IO.foreach(large_file) do | number |
      @result = meta(number.to_i, @result)
    end
  end
end



puts "File Size: #{File.size(large_file)/1048576} MB"
puts "Result:"
puts @result
